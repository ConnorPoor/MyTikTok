package controller

import (
	"github.com/ConnorPoor/MyTikTok/common"
	"github.com/ConnorPoor/MyTikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginResponse struct {
	common.Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

const (
	MaxUsernameLength = 32 //用户名最大长度
	MaxPasswordLength = 32 //密码最大长度
	MinPasswordLength = 6  //密码最小长度
)

func isUserLegal(userName string, passWord string) error {
	if userName == "" {
		return common.ErrorUserNameNull
	} else if len(userName) > MaxUsernameLength {
		return common.ErrorUserNameExtend
	}

	if passWord == "" {
		return common.ErrorPasswordNull
	} else if len(passWord) > MaxPasswordLength || len(passWord) < MinPasswordLength {
		return common.ErrorPasswordLength
	}

	return nil
}

func Register(c *gin.Context) {
	userName := c.Query("username")
	passWord := c.Query("password")

	if err := isUserLegal(userName, passWord); err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}

	newUser, err := service.CreateUser(userName, passWord)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: common.Response{StatusCode: 0},
		UserId:   newUser.ID,
		Token:    userName + passWord,
	})
}
