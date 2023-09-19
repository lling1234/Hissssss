package code

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// api
const (
	Success          = 200
	BadRequest       = 400
	Unauthorized     = 401
	PermissionDenied = 403
	NotFound         = 404
	InternalError    = 500
)

// biz
const (
	SignUpSuccess = 2001
	SignInSuccess = 2002

	NameAlreadyExists = 4001
	UserNotExists     = 4002
	PasswordWrong     = 4003
)

var Message = map[int64]string{
	Success:           "成功",
	BadRequest:        "请求错误",
	Unauthorized:      "未认证",
	PermissionDenied:  "未授权",
	NotFound:          "记录不存在",
	InternalError:     "系统内部错误, 请重试",
	NameAlreadyExists: "用户名存在",
	UserNotExists:     "用户不存在",
	PasswordWrong:     "密码验证不通过",
	SignUpSuccess:     "注册成功",
	SignInSuccess:     "登录成功",
}

func Rsp(code int64) error {
	return status.Error(codes.Code(code), Message[code])
}

func From(err error) (int64, string) {
	s := status.Convert(err)
	return int64(s.Code()), s.Message()
}
