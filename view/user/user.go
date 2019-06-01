package user

import (
	"net/http"

	"github.com/dovics/catapult/controller"
	"github.com/dovics/catapult/model"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
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
	dao.CreateTable()
	return &Controller{
		controller.New(valid),
		dao,
	}
}

// Register -
func (controller *Controller) Register(c echo.Context) error {
	var req struct {
		name     string `json: "name" validate:"required, alphanum, min=2, max=30"`
		password string `json: "pwd"  validete:"printascii, min=6, max=30"`
	}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := controller.Validator().Struct(req); err != nil {
		return err
	}

	if err := controller.dao.InsertUser(req.name, req.password); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
