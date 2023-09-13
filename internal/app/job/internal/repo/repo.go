package repo

import (
	"github.com/cd-home/Hissssss/internal/app/job/internal/adapter"
	"github.com/cd-home/Hissssss/internal/app/job/internal/model"
	"github.com/cd-home/Hissssss/internal/pkg/xmongo"
)

type JobRepo struct {
	mongo *xmongo.XMongo
}

func New(mongo *xmongo.XMongo) adapter.JobRepo {
	return &JobRepo{mongo: mongo}
}

func (j *JobRepo) CreateOfflineMessage(msg *model.OfflineMessage) error {
	return j.mongo.InsertModel("test", "offlineMessage").Multi(false).Doc(msg).Do()
}
