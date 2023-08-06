package model

import (
	"github.com/cd-home/Hissssss/api/pb/common"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            uint             `gorm:"primarykey"` // 采用主键做唯一的用户id, 业务id 起始值从10000开始
	UserName      string           `gorm:"column:username;type:varchar(20) not null"`
	Password      string           `gorm:"column:password;type:varchar(255)"`
	NickName      string           `gorm:"column:nickname;type:varchar(20);default:null"`
	Phone         string           `gorm:"column:phone;type:char(11);default:null"`
	Email         string           `gorm:"column:email;type:varchar(255);default:null"`
	Address       string           `gorm:"column:address;type:varchar(255);default:null"`
	Region        string           `gorm:"column:region;type:varchar(255);default:null"`
	Country       string           `gorm:"column:country;type:varchar(20);default:null"`
	IP            uint32           `gorm:"column:ip;type:uint;default:null"`
	Avatar        string           `gorm:"column:avatar;type:varchar(255);default:null"`
	Link          string           `gorm:"column:link;type:varchar(255);default:null"`
	SigninDays    int64            `gorm:"column:signin_days;type:int;default:0"`
	Level         int64            `gorm:"column:level;type:int;default:0"`
	Gender        common.Gender    `gorm:"column:gender;type:tinyint;default:0"`
	State         common.State     `gorm:"column:status;type:tinyint;default:2"`
	SignupWay     common.SignupWay `gorm:"column:signup_way;type:tinyint;default:1"`
	Vip           common.Vip       `gorm:"column:vip;type:int;default:1"`
	VipExpire     time.Time        `gorm:"vip_expire;type:datetime;default:null"`
	SignupTime    time.Time        `gorm:"column:signup_time;type:datetime;autoCreateTime"`
	LastLoginTime time.Time        `gorm:"column:last_login_time;type:datetime;default:null"`
	CreatedAt     time.Time        `gorm:"column:create_at;type:datetime;autoCreateTime"`
	UpdatedAt     time.Time        `gorm:"column:update_at;type:datetime;autoUpdateTime"`
	DeletedAt     time.Time        `gorm:"column:delete_at;type:datetime;index;default:null"`
}

type OAuth struct {
	ID        uint             `gorm:"primarykey"`
	UID       int64            `gorm:"column:uid;type:bigint not null"`
	OAuthType common.OAuthType `gorm:"column:oauth_type;type:varchar(20)"`
	OpenID    string           `gorm:"column:open_id;type:varchar(100)"`
	UnionId   string           `gorm:"column:union_id;type:varchar(100);default:null"`
	Token     string           `gorm:"column:token;type:varchar(100);default:null"`
	Bind      common.Bind      `gorm:"column:bind;type:tinyint;default:1"`
	CreatedAt time.Time        `gorm:"column:create_at;type:datetime;autoCreateTime"`
	UpdatedAt time.Time        `gorm:"column:update_at;type:datetime;autoUpdateTime"`
	DeletedAt time.Time        `gorm:"column:delete_at;type:datetime;index;default:null"`
}

func MigrateModel(db *gorm.DB, logger *zap.Logger) {
	err := db.AutoMigrate(&User{}, &OAuth{})
	if err != nil {
		logger.With(zap.String("module", "migrate")).Error(err.Error())
		return
	}
	logger.With(zap.String("module", "migrate")).Info("init model successful")
}
