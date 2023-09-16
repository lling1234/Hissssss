package model

import (
	"github.com/cd-home/Hissssss/api/pb/common"
	"time"
)

type OAuth struct {
	ID        uint             `bson:"ID"`
	UID       int64            `bson:"UID"`
	OAuthType common.OAuthType `bson:"OAuthType"`
	OpenID    string           `bson:"OpenID"`
	UnionId   string           `bson:"UnionId"`
	Token     string           `bson:"Token"`
	Bind      common.Bind      `bson:"Bind"`
	CreatedAt time.Time        `bson:"CreatedAt"`
	UpdatedAt time.Time        `bson:"UpdatedAt"`
	DeletedAt time.Time        `bson:"DeletedAt"`
}

type Account struct {
	ID        int64     `bson:"ID"`        // 账号ID
	UserID    int64     `bson:"UserID"`    // 用户ID
	LoginCode string    `bson:"LoginCode"` // 登录标识
	LoginType string    `bson:"LoginType"` // 登录类型
	CreatedAt time.Time `bson:"CreatedAt"`
	UpdatedAt time.Time `bson:"UpdatedAt"`
	DeletedAt time.Time `bson:"DeletedAt"`
}
