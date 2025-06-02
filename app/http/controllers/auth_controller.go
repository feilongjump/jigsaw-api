package controllers

import (
	"errors"
	"github.com/feilongjump/jigsaw-api/app/http/requests"
	"github.com/feilongjump/jigsaw-api/app/http/responses"
	UserModel "github.com/feilongjump/jigsaw-api/app/models/user"
	"github.com/feilongjump/jigsaw-api/plugins/jwt"
	"github.com/feilongjump/jigsaw-api/plugins/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct{}

func (ac *AuthController) Login(ctx *gin.Context) {

	params := requests.LoginRequest{}
	if ok := requests.ValidateJSON(ctx, &params, params.GetErrMessage); !ok {
		return
	}

	user, err := UserModel.FindByMulti(params.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.NotFound(ctx, "仔细想想这个账号注册了没？")
			return
		}

		logger.Error("登录失败---" + err.Error())
		responses.Abort500(ctx, "这都能登陆失败！扣他工资！")
		return
	}

	if ok := user.ComparePassword(params.Password); !ok {
		responses.ValidatorError(ctx, map[string]string{
			"password": "密码错误",
		})
		return
	}

	authToken(ctx, user)
}

func (ac *AuthController) SignUp(ctx *gin.Context) {

	params := requests.SignUpRequest{}
	if ok := requests.ValidateJSON(ctx, &params, params.GetErrMessage); !ok {
		return
	}

	user := UserModel.User{
		Name:     params.Username,
		Email:    params.Email,
		Password: params.Password,
	}
	if err := user.Create(); err != nil {
		logger.Error("注册失败---" + err.Error())
		responses.Abort500(ctx, "注册失败？看来是不想要我这个用户了！")
		return
	}

	authToken(ctx, user)
}

func authToken(ctx *gin.Context, user UserModel.User) {

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		logger.Error("生成 token 失败---" + err.Error())
		responses.Abort500(ctx, "这只 🐛 需要被逮捕，麻烦尽快联系开发人员！")
		return
	}

	responses.Success(ctx, map[string]string{
		"token":      token,
		"token_type": "Bearer",
	})
}
