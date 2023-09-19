package model

import (
	"github.com/cd-home/Hissssss/api/pb/common"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type User struct {
	ID          int64            `bson:"ID"`
	Name        string           `bson:"Name"`
	Password    string           `bson:"Password"`
	Salt        string           `bson:"Salt"`
	Phone       string           `bson:"Phone"`
	Email       string           `bson:"Email"`
	Address     string           `bson:"Address"`
	Region      string           `bson:"Region"`
	Country     string           `bson:"Country"`
	IP          uint32           `bson:"IP"`
	Avatar      string           `bson:"Avatar"`
	Link        string           `bson:"Link"`
	SigninDays  int64            `bson:"SigninDays"`
	Level       int64            `bson:"Level"`
	Gender      common.Gender    `bson:"Gender"`
	State       common.State     `bson:"State"`
	SignupWay   common.SignupWay `bson:"SignupWay"`
	Vip         common.Vip       `bson:"Vip"`
	VipExpire   *time.Time       `bson:"VipExpire"`
	SignupAt    time.Time        `bson:"SignupAt"`
	LastLoginAt time.Time        `bson:"LastLoginAt"`
	CreatedAt   time.Time        `bson:"CreatedAt"`
	UpdatedAt   time.Time        `bson:"UpdatedAt"`
	DeletedAt   time.Time        `bson:"DeletedAt"`
}

func (u *User) CreatHook() *User {
	now := time.Now().Add(time.Hour * 8)
	u.CreatedAt = now
	u.UpdatedAt = now
	u.SignupAt = now
	u.LastLoginAt = now
	u.Vip = common.Vip_NoVip
	u.VipExpire = nil
	u.Salt = string(Lower(10))
	return u
}

func (u *User) Encrypt(Password string) {
	bcryptBytes, _ := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	u.Password = string(bcryptBytes)
}

func (u *User) Verify(Password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(Password)) == nil
}

func Lower(size int) []byte {
	if size <= 0 || size > 26 {
		size = 26
	}
	warehouse := []int{97, 122}
	result := make([]byte, 26)
	for i := 0; i < size; i++ {
		result[i] = uint8(warehouse[0] + rand.Intn(26))
	}
	return result
}
