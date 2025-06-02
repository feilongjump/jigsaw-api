package controllers

import (
	"errors"
	"github.com/feilongjump/jigsaw-api/app/http/responses"
	UserModel "github.com/feilongjump/jigsaw-api/app/models/user"
	"github.com/feilongjump/jigsaw-api/plugins/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct{}

func (uc *UserController) Me(ctx *gin.Context) {

	user, err := UserModel.FindByID(ctx.GetUint64("user_id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.NotFound(ctx, "好奇怪，你这个账号是不是有问题？")
			return
		}

		logger.Error("获取用户信息失败---" + err.Error())
		responses.Abort500(ctx, "怎么就获取用户信息失败了呢？")
		return
	}

	responses.Success(ctx, user)
}
