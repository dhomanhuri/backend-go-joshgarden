package presentation

import (
	"backend-go/features/users"
	"backend-go/features/users/presentation/dto"
	"backend-go/helper"

	"github.com/gin-gonic/gin"
)

type UserBuss struct {
	Buss users.Bussines
}

func UserHandler(UserBus users.Bussines) UserBuss {
	return UserBuss{
		Buss: UserBus,
	}
}

func (ub UserBuss) Login(c *gin.Context) {
	userRequest := dto.LoginResquest{}
	errBind := c.Bind(&userRequest)
	if errBind != nil {
		c.JSON(helper.FailedBadRequest())
		c.Abort()
		return
	}

	userCore := dto.ToCore(userRequest)
	_, token, errLogin := ub.Buss.Login(userCore)
	if errLogin != nil {
		c.JSON(helper.FailedNotFound())
		c.Abort()
		return
	}
	// fmt.Println(a)
	c.JSON(helper.AuthOK(userCore, token))
}

func (ub UserBuss) Register(c *gin.Context) {
	userReq := dto.UserResquest{}
	err := c.Bind(&userReq)
	if err != nil {
		c.JSON(helper.BadRequest("error bidn"))
		c.Abort()
		return
	}
	core := users.Core{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
		Role: users.RoleCore{
			ID: userReq.RoleID,
		},
	}

	result := ub.Buss.Register(core)
	if result != nil {
		c.JSON(helper.BadRequest(result.Error()))
		c.Abort()
		return
	}
	c.JSON(helper.SiccessCreate())
}
