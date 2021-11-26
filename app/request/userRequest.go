package request

import "efeasy-gin/utils"

type RegisterRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	// 密码  binding:"required"为必填字段,长度大于3小于20
	Password  string `form:"password" json:"password" binding:"required,min=3,max=20"`

}

// GetMessages 自定义错误信息
func (register RegisterRequest) GetMessages() utils.ValidatorMessages {
	return utils.ValidatorMessages{
		"Name.required": "用户名称不能为空",
		"Mobile.required": "手机号码不能为空",
		"mobile.mobile": "手机号码格式不正确",
		"Password.required": "用户密码不能为空",
	}
}
