package xmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertModel struct {
	ctx      context.Context
	database *mongo.Database
	col      *mongo.Collection
	docs     []any
	multi    bool
	one      []*options.InsertOneOptions
	many     []*options.InsertManyOptions
}

func (x *XMongo) InsertModel(db string, col string) *InsertModel {
	database := x.database[db]
	return &InsertModel{ctx: context.TODO(), database: database, col: database.Collection(col)}
}

func (i *InsertModel) Session(ctx context.Context) *InsertModel {
	i.ctx = ctx
	return i
}

func (i *InsertModel) Multi(multi bool) *InsertModel {
	i.multi = multi
	return i
}

func (i *InsertModel) OneOpts(opts ...*options.InsertOneOptions) *InsertModel {
	i.one = opts
	return i
}

func (i *InsertModel) ManyOpts(opts ...*options.InsertManyOptions) *InsertModel {
	i.many = opts
	return i
}

func (i *InsertModel) Doc(doc any) *InsertModel {
	i.docs = append(i.docs, doc)
	return i
}

func (i *InsertModel) Docs(docs []any) *InsertModel {
	i.docs = append(i.docs, docs...)
	return i
}

func (i *InsertModel) Do() error {
	if i.multi {
		_, err := i.col.InsertMany(i.ctx, i.docs, i.many...)
		return err
	}
	if len(i.docs) != 1 {
		return nil
	}
	_, err := i.col.InsertOne(i.ctx, i.docs[0], i.one...)
	return err
}

func (i *InsertModel) Unique() int64 {
	res := i.database.Collection("AutoInCID").FindOneAndUpdate(
		i.ctx,
		bson.M{"Col": i.col.Name()},
		bson.M{"$inc": bson.M{"Unique": 1}},
	)
	if res.Err() != nil {
		_, _ = i.database.Collection("AutoInCID").InsertOne(
			i.ctx,
			bson.M{"Col": i.col.Name(), "Unique": 1000001},
		)
		return 1000000
	}
	unique := struct {
		Unique int64 `bson:"Unique"`
	}{}
	_ = res.Decode(&unique)
	return unique.Unique
}
