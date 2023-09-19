package repo

import (
	"errors"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/adapter"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/pkg/model"
	"github.com/cd-home/Hissssss/internal/pkg/xmongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	AccountDatabase = "Account"
	User            = "User"
)

type AccountRepo struct {
	logger *zap.Logger
	mongo  *xmongo.XMongo
}

func NewAccountRepo(logger *zap.Logger, mongo *xmongo.XMongo) adapter.AccountRepo {
	return &AccountRepo{
		logger: logger.WithOptions(zap.Fields(zap.String("module", "account repo"))),
		mongo:  mongo,
	}
}

func (a *AccountRepo) Create(doc *model.User) error {
	insert := a.mongo.InsertModel(AccountDatabase, User)
	doc.ID = insert.Unique()
	doc.CreatHook().Encrypt(doc.Password)
	return insert.Multi(false).Doc(doc).Do()
}

func (a *AccountRepo) Retrieve(name string) (*model.User, bool, error) {
	var u model.User
	selectx := a.mongo.SelectModel(AccountDatabase, User)
	err := selectx.Multi(false).Filter(map[string]any{"Name": name}).Rows(&u).Do()
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, false, err
	}
	return &u, u.ID > 0, nil
}
