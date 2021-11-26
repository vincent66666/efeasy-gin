package dao

import (
	"efeasy-gin/global"
	"efeasy-gin/app/model"
	"efeasy-gin/app/request"
	"efeasy-gin/utils"
)

var users []model.User

type userDao struct {
}

var UserDao = new(userDao)

func (userDao *userDao) Exists(Mobile string) bool {
	var count int64
	global.App.DB.Where("mobile = ?", Mobile).Count(&count)
	if count != 0 {
		return true
	}
	return false
}

func (userDao *userDao) Create(params request.RegisterRequest) (err error, user model.User) {
	user = model.User{
		Name:     params.Name,
		Mobile:   params.Mobile,
		Password: utils.BcryptMake([]byte(params.Password)),
	}
	err = global.App.DB.Create(&user).Error
	return
}

// GetUserListDao GetUserList 获取用户列表(page第几页,page_size每页几条数据)
func (userDao *userDao) GetUserListDao(page int, pageSize int) (int, []interface{}) {
	// 分页用户列表数据
	userList := make([]interface{}, 0, len(users))
	// 计算偏移量
	offset := (page - 1) * pageSize
	// 查询所有的user
	result := global.App.DB.Offset(offset).Limit(pageSize).Find(&users)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, userList
	}
	// 获取user总数
	total := len(users)
	// 查询数据
	result.Offset(offset).Limit(pageSize).Find(&users)
	//
	for _, useSingle := range users {
		birthday := ""
		if useSingle.Birthday == nil {
			birthday = ""
		} else {
			// 给未设置生日的初始值
			birthday = useSingle.Birthday.Format("2006-01-02")
		}
		userItemMap := map[string]interface{}{
			"id":        useSingle.ID.ID,
			"password":  useSingle.Password,
			"nick_name": useSingle.NickName,
			"birthday":  birthday,
			"address":   useSingle.Address,
			"desc":      useSingle.Desc,
			"gender":    useSingle.Gender,
			"role":      useSingle.Role,
			"mobile":    useSingle.Mobile,
		}
		userList = append(userList, userItemMap)
	}
	return total, userList
}
