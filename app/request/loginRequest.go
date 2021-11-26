package request

import "efeasy-gin/utils"

type LoginRequest struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login LoginRequest) GetMessages() utils.ValidatorMessages {
	return utils.ValidatorMessages{
		"mobile.required": "手机号码不能为空",
		"mobile.mobile": "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}