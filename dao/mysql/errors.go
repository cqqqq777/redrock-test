package mysql

import "errors"

var (
	ErrorUserExist     = errors.New("用户已存在")
	ErrorUserNotExist  = errors.New("用户不存在")
	ErrorEmailExist    = errors.New("邮箱已注册")
	ErrorEmailNotExist = errors.New("邮箱未注册")
	ErrorWrongPassword = errors.New("密码错误")
	ErrorInvalidId     = errors.New("无效的id")
	ErrorQuestionExist = errors.New("问题已存在")
	ErrorNoPermission  = errors.New("无权限")
	ErrorRepeatOperate = errors.New("重复操作")
)
