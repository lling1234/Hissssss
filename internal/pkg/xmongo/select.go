package xmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SelectModel struct {
	ctx       context.Context
	database  *mongo.Database
	col       *mongo.Collection
	filter    bson.M
	multi     bool
	oneOpts   []*options.FindOneOptions
	manyOpts  []*options.FindOptions
	countOpts []*options.CountOptions
	rows      any
	sort      bson.D
}

func (x *XMongo) SelectModel(db string, col string) *SelectModel {
	database := x.database[db]
	return &SelectModel{
		ctx:      context.TODO(),
		database: database,
		col:      database.Collection(col),
		sort:     make(bson.D, 0),
	}
}

func (s *SelectModel) Session(ctx context.Context) *SelectModel {
	s.ctx = ctx
	return s
}

func (s *SelectModel) Filter(filter map[string]any) *SelectModel {
	s.filter = filter
	return s
}

func (s *SelectModel) OneOpts(opts ...*options.FindOneOptions) *SelectModel {
	s.oneOpts = opts
	return s
}

func (s *SelectModel) ManyOpts(opts ...*options.FindOptions) *SelectModel {
	s.manyOpts = append(s.manyOpts, opts...)
	return s
}

func (s *SelectModel) CountOpts(opts ...*options.CountOptions) *SelectModel {
	s.countOpts = opts
	return s
}

func (s *SelectModel) Multi(multi bool) *SelectModel {
	s.multi = multi
	return s
}

func (s *SelectModel) Rows(rows any) *SelectModel {
	s.rows = rows
	return s
}

func (s *SelectModel) Asc(key string) *SelectModel {
	s.sort = append(s.sort, bson.E{Key: key, Value: true})
	return s
}

func (s *SelectModel) Desc(key string) *SelectModel {
	s.sort = append(s.sort, bson.E{Key: key, Value: true})
	return s
}

func (s *SelectModel) Skip(skip int64) *SelectModel {
	s.manyOpts = append(s.manyOpts, options.Find().SetSkip(skip))
	return s
}

func (s *SelectModel) Limit(limit int64) *SelectModel {
	s.manyOpts = append(s.manyOpts, options.Find().SetLimit(limit))
	return s
}

func (s *SelectModel) Do() error {
	if s.multi {
		s.manyOpts = append(s.manyOpts, options.Find().SetSort(s.sort))
		cursor, err := s.col.Find(s.ctx, s.filter, s.manyOpts...)
		if err != nil {
			return err
		}
		return cursor.All(s.ctx, s.rows)
	}
	s.oneOpts = append(s.oneOpts, options.FindOne().SetSort(s.sort))
	return s.col.FindOne(s.ctx, s.filter, s.oneOpts...).Decode(s.rows)
}

func (s *SelectModel) Count() (int64, error) {
	return s.col.CountDocuments(s.ctx, s.filter, s.countOpts...)
}
