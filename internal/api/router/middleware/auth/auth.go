package auth

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/token"

	"github.com/pkg/errors"
)

func AuthHandler(ctx core.Context) (userId int64, userName string, err errno.Error) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("Header 中缺少 Authorization 参数"))

		return
	}

	cfg := configs.Get().JWT
	claims, errParse := token.New(cfg.Secret).Parse(auth)
	if errParse != nil {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errParse)

		return
	}

	userId = claims.UserID
	if userId <= 0 {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("claims.UserID <= 0 "))

		return
	}
	userName = claims.UserName
	return
}
