package xmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeleteModel struct {
	ctx      context.Context
	database *mongo.Database
	col      *mongo.Collection
	filter   bson.M
	multi    bool
	opts     []*options.DeleteOptions
}

func (x *XMongo) DeleteModel(db string, col string) *DeleteModel {
	database := x.database[db]
	return &DeleteModel{
		ctx:      context.TODO(),
		database: database,
		col:      database.Collection(col),
	}
}

func (d *DeleteModel) Session(ctx context.Context) *DeleteModel {
	d.ctx = ctx
	return d
}

func (d *DeleteModel) Filter(filter map[string]any) *DeleteModel {
	d.filter = filter
	return d
}

func (d *DeleteModel) Multi(multi bool) *DeleteModel {
	d.multi = multi
	return d
}

func (d *DeleteModel) Opts(opts ...*options.DeleteOptions) *DeleteModel {
	d.opts = opts
	return d
}

func (d *DeleteModel) Do() error {
	if d.multi {
		_, err := d.col.DeleteMany(d.ctx, d.filter, d.opts...)
		return err
	}
	_, err := d.col.DeleteOne(d.ctx, d.filter, d.opts...)
	return err
}
