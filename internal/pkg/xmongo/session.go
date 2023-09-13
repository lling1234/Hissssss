package xmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
)

type SessionModel struct {
	ctx      context.Context
	client   *mongo.Client
	database map[string]*mongo.Database
	opts     []*options.SessionOptions
}

func (x *XMongo) SessionModel() *SessionModel {
	return &SessionModel{
		ctx:      context.TODO(),
		client:   x.client,
		database: x.database,
	}
}

func (s *SessionModel) Opts(opts ...*options.SessionOptions) *SessionModel {
	s.opts = opts
	s.opts = append(s.opts, options.Session().SetDefaultReadConcern(readconcern.Majority()))
	return s
}

func (s *SessionModel) WithSession(f func(session mongo.SessionContext) error) error {
	session, _ := s.client.StartSession(s.opts...)
	defer session.EndSession(s.ctx)
	err := mongo.WithSession(s.ctx, session, func(sessionContext mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			return err
		}
		err = f(sessionContext)
		if err != nil {
			_ = session.AbortTransaction(s.ctx)
			return err
		}
		return session.CommitTransaction(s.ctx)
	})
	return err
}

func (s *SessionModel) WithTransaction(f func(session mongo.SessionContext) error) error {
	session, _ := s.client.StartSession(s.opts...)
	defer session.EndSession(s.ctx)
	_, err := session.WithTransaction(s.ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		return nil, f(ctx)
	})
	return err
}
