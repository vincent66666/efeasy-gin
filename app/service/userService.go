package service

import (
	"efeasy-gin/app/dao"
	"efeasy-gin/app/model"
	"efeasy-gin/app/request"
	"efeasy-gin/global"
	"efeasy-gin/utils"
	"errors"
	"strconv"
)

type userService struct {
}

var UserService = new(userService)

// Login 登录
func (userService *userService) Login(params request.LoginRequest) (err error, user *model.User) {
	err = global.App.DB.Where("mobile = ?", params.Mobile).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user model.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}



// Register 注册
func (userService *userService) Register(params request.RegisterRequest) (err error, user model.User) {
	mobileExists := dao.UserDao.Exists(params.Mobile)
	if mobileExists {
		err = errors.New("手机号已存在")
		return
	}
	err, users := dao.UserDao.Create(params)
	if err != nil {
		return err, model.User{}
	}
	return nil, users
}

// GetUserList  列表
func (userService *userService) GetUserList() (int, []interface{})  {
	// 获取数据
	return dao.UserDao.GetUserListDao(1, 10)
}
