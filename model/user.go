package model

import (
	"log"
	"time"

	"github.com/go-xorm/xorm"
)

// UserDao -
type UserDao struct {
	Engine *xorm.Engine
}

// User -
type User struct {
	ID          int       `xorm:"not null pk autoincr INTEGER" `
	Name        string    `xorm:"varchar(200)"`
	Password    string    `xorm:"varchar(200)"`
	Mobile      string    `xorm:"varchar(32)"`
	Email       string    `xorm:"varchar(200)"`
	Avatar      string    `xorm:"varchar(200)"`
	description string    `xorm:"varchar(200)"`
	CreateAt    time.Time `xorm:"created"`
	Active      bool      `xorm:"tinyint"`
}

// NewDao -
func NewDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		Engine: engine,
	}
}

// CreateTable -
func (dao *UserDao) CreateTable() {
	if err := dao.Engine.Sync2(new(User)); err != nil {
		log.Fatal("CreateDatabase Error", err)
	}
}

// InsertUser -
func (dao *UserDao) InsertUser(name string, pwd string) error {
	user := &User{
		Name:     name,
		Password: pwd,
	}

	_, err := dao.Engine.Insert(user)
	if err != nil {
		return err
	}

	return nil
}
