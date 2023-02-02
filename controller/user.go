package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"redrock-test/dao/mysql"
	"redrock-test/dao/redisdao"
	g "redrock-test/global"
	"redrock-test/model"
	"redrock-test/services"
	"redrock-test/utils"
	"strconv"
)

// PostVerification 发送验证码
func PostVerification(c *gin.Context) {
	//获取邮箱并校验
	email := c.PostForm("email")
	//判断邮箱格式是否正确
	if email == "" || !utils.CheckEmail(email) {
		RespFailed(c, CodeInvalidParam)
		return
	}
	//给邮箱发验证码
	if err := services.PostVerification(email); err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Debug(err.Error())
		return
	}
	//返回响应
	RespSuccess(c, nil)
}

// Register 注册
func Register(c *gin.Context) {
	ParamUser := new(model.ParamRegisterUser)
	if err := c.ShouldBindJSON(ParamUser); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if ParamUser.Password != ParamUser.RePassword {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if err := services.Register(ParamUser); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			RespFailed(c, CodeUserExist)
			return
		}
		if errors.Is(err, mysql.ErrorEmailExist) {
			RespFailed(c, CodeEmailExist)
			return
		}
		if errors.Is(err, redisdao.ErrorInvalidVerification) {
			RespFailed(c, CodeWrongVerification)
			return
		}
		RespFailed(c, CodeServiceBusy)
		g.Logger.Debug(err.Error())
		return
	}
	RespSuccess(c, nil)
}

// Login 登录
func Login(c *gin.Context) {
	ParamUser := new(model.ParamLoginUser)
	if err := c.ShouldBindJSON(ParamUser); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	ParamUser.Password = utils.Md5(ParamUser.Password)
	//判断通过邮箱还是用户名登录
	var (
		token string
		err   error
		uid   int
	)
	if utils.CheckEmail(ParamUser.UsernameOrEmail) {
		uid, token, err = services.LoginByEmail(ParamUser)
	} else {
		uid, token, err = services.LoginByUsername(ParamUser)
	}
	//判断错误类型
	if err != nil {
		if errors.Is(err, mysql.ErrorUserNotExist) {
			RespFailed(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, mysql.ErrorEmailNotExist) {
			RespFailed(c, CodeEmailNotExist)
			return
		}
		if errors.Is(err, mysql.ErrorWrongPassword) {
			RespFailed(c, CodeWrongPassword)
			return
		}
		RespFailed(c, CodeServiceBusy)
		return
	}
	//返回token
	RespSuccess(c, &model.ApiUser{
		Uid:   uid,
		Token: token,
	})
}

// LoginByVerification 通过验证码登录
func LoginByVerification(c *gin.Context) {
	ParamUser := new(model.ParamLoginByVerificationUser)
	if err := c.ShouldBindJSON(ParamUser); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if !utils.CheckEmail(ParamUser.Email) {
		RespFailed(c, CodeInvalidParam)
		return
	}
	data, err := services.LoginByVerification(ParamUser)
	if err != nil {
		if errors.Is(err, mysql.ErrorEmailNotExist) {
			RespFailed(c, CodeEmailNotExist)
			return
		}
		if errors.Is(err, redisdao.ErrorInvalidVerification) {
			RespFailed(c, CodeWrongVerification)
			return
		}
		RespFailed(c, CodeServiceBusy)
		g.Logger.Debug(err.Error())
		return
	}
	RespSuccess(c, data)
}

// RevisePassword 修改密码
func RevisePassword(c *gin.Context) {
	ParamUser := new(model.ParamReviseUser)
	if err := c.ShouldBindJSON(ParamUser); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if ParamUser.OriPassword == "" || ParamUser.NewPassword == "" || ParamUser.RePassword == "" || ParamUser.NewPassword != ParamUser.RePassword || ParamUser.OriPassword == ParamUser.NewPassword {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	ParamUser.Uid = uid
	ParamUser.NewPassword = utils.Md5(ParamUser.NewPassword)
	ParamUser.OriPassword = utils.Md5(ParamUser.OriPassword)
	if err := services.RevisePassword(ParamUser); err != nil {
		if errors.Is(err, mysql.ErrorWrongPassword) {
			RespFailed(c, CodeWrongPassword)
			return
		}
		RespFailed(c, CodeServiceBusy)
		return
	}
	RespSuccess(c, nil)
}

// ReviseUsername 修改用户名
func ReviseUsername(c *gin.Context) {
	ParamUser := new(model.ParamReviseUser)
	if err := c.ShouldBindJSON(ParamUser); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if ParamUser.NewUsername == "" {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	ParamUser.Uid = uid
	if err := services.ReviseUsername(ParamUser); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			RespFailed(c, CodeUserExist)
			return
		}
		RespFailed(c, CodeServiceBusy)
		return
	}
	RespSuccess(c, nil)
}

// ForgetPassword 忘记密码
func ForgetPassword(c *gin.Context) {
	paramUser := new(model.ParamRegisterUser)
	if err := c.ShouldBindJSON(paramUser); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if paramUser.Password != paramUser.RePassword {
		RespFailed(c, CodeInvalidParam)
		return
	}
	paramUser.Password = utils.Md5(paramUser.Password)
	if err := services.ForgetPassword(paramUser); err != nil {
		if errors.Is(err, mysql.ErrorEmailNotExist) {
			RespFailed(c, CodeEmailNotExist)
			return
		}
		if errors.Is(err, redisdao.ErrorInvalidVerification) {
			RespFailed(c, CodeWrongVerification)
			return
		}
		RespFailed(c, CodeServiceBusy)
		return
	}
	RespSuccess(c, nil)
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	data, err := services.GetUserInfo(uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			RespFailed(c, CodeInvalidId)
			return
		}
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, data)
}

// UpdateUserInfo 修改用户信息
func UpdateUserInfo(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	reallyUid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	user := new(model.User)
	if err = c.ShouldBindJSON(user); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	user.Uid = reallyUid
	err = services.UpdateUserInfo(int(uid), user)
	if err != nil {
		if errors.Is(err, mysql.ErrorNoPermission) {
			RespFailed(c, CodeNoPermission)
			return
		}
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, nil)
}
