package model

import (
	"github.com/cd-home/Hissssss/api/pb/common"
	"time"
)

// OfflineMessage 离线消息表存储mongo
type OfflineMessage struct {
	MsgId    int64           `bson:"MsgId"`
	MsgType  common.PushType `bson:"MsgType"`
	From     int64           `bson:"From"`
	To       int64           `bson:"To"`
	SendTime time.Time       `bson:"SendTime"`
	Received time.Time       `bson:"ReceiveTime"`
	Receive  bool            `bson:"Received"`
}
