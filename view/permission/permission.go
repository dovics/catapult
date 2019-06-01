package user

import (
	"github.com/dovics/catapult/model"
	"github.com/go-xorm/xorm"
	"gopkg.in/go-playground/validator.v8"
)

// Controller -
type Controller struct {
	*controller.BaseController
	dao *model.UserDao
}

// New -
func New(engine *xorm.Engine, valid *validator.Validate) *Controller {
	dao := model.NewDao(engine)
	return &Controller{
		controller.New(valid),
		dao,
	}
}
