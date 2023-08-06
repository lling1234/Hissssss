package xgorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type Config struct {
	User    string        `yaml:"user"`
	Pwd     string        `yaml:"pwd"`
	Addr    string        `yaml:"addr"`
	Port    string        `yaml:"port"`
	DB      string        `yaml:"db"`
	MaxIdle int           `yaml:"maxIdle"`
	MaxOpen int           `yaml:"maxOpen"`
	MaxLife time.Duration `yaml:"maxLife"`
}

func New(config Config) *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		config.User,
		config.Pwd,
		config.Addr,
		config.Port,
		config.DB,
	)
	db_, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NameReplacer:  nil,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	sqlDb, _ := db_.DB()
	sqlDb.SetMaxIdleConns(config.MaxIdle)    // must < max_open_conn
	sqlDb.SetMaxOpenConns(config.MaxOpen)    // must < max_connections | show variables like ‘max_connections’
	sqlDb.SetConnMaxLifetime(config.MaxLife) // must < wait_timeout | show variables like 'wait_timeout'
	return db_
}
