package controller

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	ctx "github.com/yeyee2901/gin-web/src/context"
	m "github.com/yeyee2901/gin-web/src/model"
	rc "github.com/yeyee2901/gin-web/src/response-code"
	u "github.com/yeyee2901/gin-web/src/utils"
)

func (appCtx *AppContext) AuthMiddleware(reqContext *gin.Context) {
	// akan return empty string jika tidak ada
	token := reqContext.GetHeader("X-TOKEN")

	// Panggil handler selanjutnya untuk endpoint yang tidak perlu token
	if strings.Contains(reqContext.Request.URL.Path, "/api/doc") && gin.Mode() == gin.DebugMode {
		reqContext.Next()
		return
	}

	if strings.Contains(reqContext.Request.URL.Path, "/login") {
		reqContext.Next()
		return
	}

	if strings.Contains(reqContext.Request.URL.Path, "/config") {
		reqContext.Next()
		return
	}

	// get config dari context
	k := ctx.Key("config")
	config, ok := appCtx.Config.Value(k).(*ctx.Config)

	// cek type cast
	if !ok {
		panic("[FATAL] Gagal load config context")
	}

	// Cek token
	userData, err := u.DecomposeToken(config, token)

	if err != nil {
		reason := fmt.Sprintf("Token Error: %s", err.Error())
		failRes := &m.FailResponse{
			Msg: reason,
		}

		reqContext.AbortWithStatusJSON(rc.TokenInvalid, failRes)
		return
	}

	// set context agar middleware2 selanjutnya dapat akses ke data
	// user dari token ini.
	reqContext.Set("userToken", userData)

	// Set token di header X-TOKEN, supaya user dapat token session nya lagi
	reqContext.Header("X-TOKEN", token)

	// panggil handler selanjutnya
	reqContext.Next()
}
