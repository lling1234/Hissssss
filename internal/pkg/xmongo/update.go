package xmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UpdateModel struct {
	ctx      context.Context
	database *mongo.Database
	col      *mongo.Collection
	filter   bson.M
	update   bson.M
	multi    bool
	opts     []*options.UpdateOptions
}

func (x *XMongo) UpdateModel(db string, col string) *UpdateModel {
	database := x.database[db]
	return &UpdateModel{
		ctx:      context.TODO(),
		database: database,
		col:      database.Collection(col),
		update:   make(bson.M, 1),
	}
}

func (u *UpdateModel) Session(ctx context.Context) *UpdateModel {
	u.ctx = ctx
	return u
}

func (u *UpdateModel) Filter(filter map[string]any) *UpdateModel {
	u.filter = filter
	return u
}

func (u *UpdateModel) Update(update map[string]any) *UpdateModel {
	u.update["$set"] = update
	return u
}

func (u *UpdateModel) Inc(inc map[string]any) *UpdateModel {
	u.update["$inc"] = inc
	return u
}

func (u *UpdateModel) Opts(opts ...*options.UpdateOptions) *UpdateModel {
	u.opts = opts
	return u
}

func (u *UpdateModel) Multi(multi bool) *UpdateModel {
	u.multi = multi
	return u
}

func (u *UpdateModel) Do() error {
	if u.multi {
		_, err := u.col.UpdateMany(u.ctx, u.filter, u.update, u.opts...)
		return err
	}
	_, err := u.col.UpdateOne(u.ctx, u.filter, u.update, u.opts...)
	return err
}
