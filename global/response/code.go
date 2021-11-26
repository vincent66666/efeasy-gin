package response

type Respond struct {
	Code    int
	Message string
}

type CustomCode struct {
	BusinessError Respond
	ValidateError Respond
	TokenError    Respond
	Success       Respond
}

var Enum = CustomCode{
	BusinessError: Respond{40000, "业务错误"},
	ValidateError: Respond{42200, "请求参数错误"},
	Success:       Respond{200, "请求成功"},
	TokenError:    Respond{40100, "登录授权失效"},
}
