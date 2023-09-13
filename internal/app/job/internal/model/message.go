package model

import (
	"github.com/cd-home/Hissssss/api/pb/common"
	"time"
)

// OfflineMessage 离线消息表存储mongo
type OfflineMessage struct {
	MsgId    int64           `bson:"column:MsgId"`
	MsgType  common.PushType `bson:"column:MsgType"`
	Sender   int64           `bson:"column:Sender"`
	Receiver int64           `bson:"column:Receiver"`
	SendTime time.Time       `bson:"column:SendTime"`
	Received time.Time       `bson:"column:ReceiveTime"`
	Receive  bool            `bson:"column:Received"`
}
