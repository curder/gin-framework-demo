package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required,regexRule"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 自定义验证规则
var regexRule validator.Func = func(fl validator.FieldLevel) (ok bool) {
	var (
		value string
	)

	if value, ok = fl.Field().Interface().(string); !ok {
		return ok
	}

	ok, _ = regexp.MatchString("^[0-9A-Za-z]{6,8}$", value)

	return
}

// 绑定模型获取验证错误的方法
func (r *Login) GetError(errors validator.ValidationErrors) string {
	var (
		val validator.FieldError
	)

	for _, val = range errors {
		if val.Field() == "User" {
			switch val.Tag() {
			case "required":
				return "请输入用户名"
			case "regex":
				return "用户名规则有误"

			}
		}

		if val.Field() == "Password" {
			switch val.Tag() {
			case "required":
				return "请输入密码"
			}
		}
	}

	return "参数错误"
}

func main() {
	var (
		r   *gin.Engine
		err error
	)
	r = gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("regexRule", regexRule)
	}

	r.POST("/", func(ctx *gin.Context) {
		var (
			json Login
			err  error
		)

		if err = ctx.BindJSON(&json); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": json.GetError(err.(validator.ValidationErrors)),
			})
			return
		}

		if json.User == "manu" && json.Password == "123" {
			ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	})

	if err = r.Run(); err != nil {
		panic(err)
	}
}
