package http

import (
	"context"
	"github.com/cd-home/Hissssss/api/pb/api"
	"github.com/cd-home/Hissssss/internal/app/api/config"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func StartHttpServer(config config.Config, logger *zap.Logger) {
	gwMux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			md := metadata.Pairs("auth", header)
			return md
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, m runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			newError := runtime.HTTPStatusError{
				HTTPStatus: 400,
				Err:        err,
			}
			runtime.DefaultHTTPErrorHandler(ctx, mux, m, writer, request, &newError)
		}),
		// StatusMethodNotAllowed StatusNotFound StatusBadRequest
		runtime.WithRoutingErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, m runtime.Marshaler, writer http.ResponseWriter, request *http.Request, i int) {
		}),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := api.RegisterApiHandlerFromEndpoint(
		context.Background(),
		gwMux,
		config.Spec.Node.Addr,
		opts,
	)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	gwServer := gin.New()
	gwServer.Use(gin.Logger())
	gin.SetMode(gin.ReleaseMode)
	_ = gwServer.SetTrustedProxies([]string{"*"})
	gwServer.Group("v1/*{grpc_gateway}").Any("", gin.WrapH(gwMux))
	logger.Info("serving grpc-gateway on http://0.0.0.0:" + config.Spec.Node.HTTP)
	_ = gwServer.Run(":" + config.Spec.Node.HTTP)
}
