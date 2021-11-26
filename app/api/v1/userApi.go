package v1

import (
	"efeasy-gin/app/request"
	"efeasy-gin/app/service"
	"efeasy-gin/utils"
	"efeasy-gin/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)


type userApi struct {
}

var UserApi = new(userApi)

func (userApi *userApi)  UserCreate(c *gin.Context) {
	var form request.RegisterRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	err, user := service.UserService.Register(form)
	if err != nil {
		response.Fail(c,500, err.Error())
		return
	}
	response.Success(c, user)
	return
}

func (userApi *userApi)  GetUserList(c *gin.Context) {

	total, userList := service.UserService.GetUserList()
	// 判断
	if (total + len(userList)) == 0 {
		response.BusinessFail(c, "未获取到到数据")
		return
	}
	response.Success(c,  map[string]interface{}{
		"total":    total,
		"userList": userList,
	})
	return
}

// PutHeaderImage 上传用户头像
func (userApi *userApi) PutHeaderImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileObj, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 把文件上传到minio对应的桶中
	ok := utils.UploadFile("user_header", file.Filename, fileObj, file.Size)
	if !ok {
		response.Fail(c, 401, "头像上传失败")
		return
	}
	headerUrl := utils.GetFileUrl("user_header", file.Filename, time.Second*24*60*60)
	if headerUrl == "" {
		response.Fail(c, 400, "获取用户头像失败")
		return
	}
	//TODO 把用户的头像地址存入到对应user表中head_url 中
	response.Success(c,  map[string]interface{}{
		"userHeaderUrl": headerUrl,
	})
	return
}

