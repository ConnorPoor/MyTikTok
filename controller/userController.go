package controller

import (
	"github.com/ConnorPoor/MyTikTok/common"
	"github.com/ConnorPoor/MyTikTok/model"
	"github.com/ConnorPoor/MyTikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	MaxUsernameLength = 32 //用户名最大长度
	MaxPasswordLength = 32 //密码最大长度
	MinPasswordLength = 6  //密码最小长度
)

type UserLoginResponse struct {
	common.Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoQueryResponse struct {
	UserId         uint   `json:"user_id"`
	UserName       string `json:"name"`
	FollowCount    uint   `json:"follow_count"`
	FollowerCount  uint   `json:"follower_count"`
	IsFollow       bool   `json:"is_follow"`
	TotalFavorited uint   `json:"total_favorited"`
	FavoriteCount  uint   `json:"favorite_count"`
}

type UserInfoResponse struct {
	common.Response
	UserList UserInfoQueryResponse `json:"user"`
}

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
		Response: common.Response{StatusCode: 0, StatusMsg: "注册成功"},
		UserId:   newUser.ID,
		Token:    userName,
	})
}

func Login(c *gin.Context) {
	userName := c.Query("username")
	passWord := c.Query("password")
	login := model.User{}
	if err := service.IsUserExit(userName, passWord, &login); err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: common.Response{StatusCode: 0, StatusMsg: "登录成功"},
		UserId:   login.ID,
		Token:    userName,
	})
}

func UserInfo(c *gin.Context) {
	rawID := c.Query("user_id")
	token := c.Query("token")
	userID, _ := strconv.ParseUint(rawID, 10, 64)
	var targetInfo = model.User{}
	var hostInfo = model.User{}
	if err := service.GetUserByID(uint(userID), &targetInfo); err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
	if err := service.GetUserByToken(token, &hostInfo); err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
	var userInfoQueryResponse = UserInfoQueryResponse{
		UserId:         hostInfo.ID,
		UserName:       hostInfo.Name,
		FollowCount:    hostInfo.FollowCount,
		FollowerCount:  hostInfo.FollowerCount,
		IsFollow:       service.CheckIsFollow(targetInfo.ID, hostInfo.ID),
		TotalFavorited: hostInfo.TotalFavorited,
		FavoriteCount:  hostInfo.FavoriteCount,
	}

	c.JSON(http.StatusOK, UserInfoResponse{
		Response: common.Response{StatusCode: 0, StatusMsg: "用户信息查看成功"},
		UserList: userInfoQueryResponse,
	})
}
