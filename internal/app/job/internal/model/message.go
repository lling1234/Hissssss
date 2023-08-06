package model

import (
	"github.com/cd-home/Hissssss/api/pb/common"
	"time"
)

// OfflineMessage 离线消息表存储MySQL
type OfflineMessage struct {
	ID        uint            `gorm:"primarykey"`
	MsgId     int64           `gorm:"column:msg_id;type:bigint not null;"`
	MsgType   common.PushType `gorm:"column:msg_type"`
	Sender    int64           `gorm:"column:sender"`
	Receiver  int64           `gorm:"column:receiver"`
	SendTime  time.Time       `gorm:"column:send_time"`
	Received  time.Time       `gorm:"column:received"`
	IsReceive bool            `gorm:"column:is_receive"`
}
