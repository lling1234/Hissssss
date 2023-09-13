package xmongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestXMongo_InsertModel(t *testing.T) {
	mgo := New(Config{
		Host: "10.211.55.18",
		Port: "27017",
		User: "mongo",
		Pwd:  "mongo",
		DB:   []string{"test"},
	})
	err := mgo.InsertModel("test", "api").Doc(map[string]any{"router": "/c"}).Do()
	if err != nil {
		t.Fatal(err)
		return
	}
}

func TestXMongo_UpdateModel(t *testing.T) {
	mgo := New(Config{
		Host: "10.211.55.18",
		Port: "27017",
		User: "mongo",
		Pwd:  "mongo",
		DB:   []string{"test"},
	})
	err := mgo.UpdateModel("test", "api").
		Filter(map[string]any{"router": "/b"}).
		Inc(map[string]any{"num": 1}).Multi(false).Do()
	if err != nil {
		return
	}
}

func TestXMongo_DeleteModel(t *testing.T) {
	mgo := New(Config{
		Host: "10.211.55.18",
		Port: "27017",
		User: "mongo",
		Pwd:  "mongo",
		DB:   []string{"test"},
	})
	err := mgo.DeleteModel("test", "api").Filter(map[string]any{"router": "/b"}).Multi(false).Do()
	if err != nil {
		t.Fatal(err)
		return
	}
}

func TestXMongo_SelectModel(t *testing.T) {
	mgo := New(Config{
		Host: "10.211.55.18",
		Port: "27017",
		User: "mongo",
		Pwd:  "mongo",
		DB:   []string{"test"},
	})
	var m map[string]any
	err := mgo.SelectModel("test", "api").
		Filter(map[string]any{"router": "/a"}).
		Multi(false).
		Rows(&m).
		Do()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(m)
}

func TestXMongo_SessionModel(t *testing.T) {
	mgo := New(Config{
		Host: "10.211.55.18",
		Port: "27017",
		User: "mongo",
		Pwd:  "mongo",
		DB:   []string{"test"},
	})
	err := mgo.SessionModel().WithSession(func(session mongo.SessionContext) error {
		err := mgo.InsertModel("test", "api").
			Session(session).Multi(false).Doc(map[string]any{"router": "/a"}).Do()
		if err != nil {
			return err
		}
		err = mgo.DeleteModel("test", "api").Session(session).Multi(false).Filter(map[string]any{"router": "/f"}).Do()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Log(err)
	}
}

func TestXMongo_SessionModel2(t *testing.T) {
	mgo := New(Config{
		Host: "10.211.55.18",
		Port: "27017",
		User: "mongo",
		Pwd:  "mongo",
		DB:   []string{"test"},
	})
	err := mgo.SessionModel().WithTransaction(func(session mongo.SessionContext) error {
		err := mgo.InsertModel("test", "api2").
			Session(session).Multi(false).Doc(map[string]any{"router": "/f"}).Do()
		if err != nil {
			return err
		}
		err = mgo.DeleteModel("test", "api").Session(session).Multi(false).Filter(map[string]any{"router": "/f"}).Do()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Log(err)
	}
}
