package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	ctx "github.com/yeyee2901/gin-web/src/context"
	m "github.com/yeyee2901/gin-web/src/model"
	rc "github.com/yeyee2901/gin-web/src/response-code"
	utils "github.com/yeyee2901/gin-web/src/utils"
)

// RE EXPORT
// ini perlu agar method nya bisa di definisikan
type AppContext ctx.AppContext

// @Tags login
// @Router /login [post]
// @Summary Login user
// @Description Login user dengan informasi dari JSON input.
// @Description Jika sukses, pada response header akan disertakan token login (key: X-TOKEN)
//
// @Accepts json
// @Produce json
//
// @Param   body    body    model.UserData    true      "Body JSON input"
//
// @Success 200 {object} model.LoginResponse            "Successful Operation"
// @Header  200 {string} X-TOKEN                        "Token login"
//
// @Failure 206 {object} model.FailResponse             "Error lain"
// @Failure 207 {object} model.FailResponse             "Gagal Login"
// @Failure 209 {object} model.FailResponse             "Invalid JSON input"
// @Failure 210 {object} model.FailResponse             "Error Middleware Internal"
// @Failure 214 {object} model.FailResponse             "Data Pengguna Tidak Ditemukan"
// @Failure 500 {string} Internal Server Fatal Error
func (appCtx *AppContext) Login(reqContext *gin.Context) {
    // validasi json input
	userData := new(m.UserData)
	err := reqContext.ShouldBindJSON(&userData)

	if err != nil {
		failRes := m.FailResponse{
			Msg: "Gagal login. Struktur data input tidak valid",
		}

		reqContext.JSON(rc.GagalLogin, failRes)
		return
	}

    // get config context
	k := ctx.Key("config")
	config, ok := appCtx.Config.Value(k).(*ctx.Config)

	// cek type cast
	if !ok {
		panic("[FATAL] Gagal load config context")
	}


    // generate token
	token, err := utils.ComposeToken(config, userData)
	if err != nil {
		msg := fmt.Sprintf("Gagal Login: Token gagal dibuat. %s", err.Error())
		failRes := &m.FailResponse{
			Msg: msg,
		}

		reqContext.JSON(rc.GagalLogin, failRes)
		return
	}

    res := &m.LoginResponse{
        Nama: userData.Nama,
        Token: token,
    }

    reqContext.Header("X-TOKEN", token) // set token di header
    reqContext.JSON(rc.Sukses, res)
    return
}


// @Tags test
// @Router /cek_middleware [get]
// @Summary Cek kondisi middleware yang digunakan untuk authorization
//
// @Accepts json
// @Produce json
//
// @Param   X-TOKEN header  string              true    "Token Login"
//
// @Success 200 {object} model.UserData                 "Successful Operation"
// @Header  200 {string} X-TOKEN                        "Token login"
//
// @Failure 206 {object} model.FailResponse             "Error lain"
// @Failure 207 {object} model.FailResponse             "Gagal Login"
// @Failure 209 {object} model.FailResponse             "Invalid JSON input"
// @Failure 210 {object} model.FailResponse             "Error Middleware Internal"
// @Failure 214 {object} model.FailResponse             "Data Pengguna Tidak Ditemukan"
// @Failure 500 {string} Internal Server Fatal Error
func (appCtx *AppContext) CekMiddleware(reqContext *gin.Context) {
	dataUser, exist := reqContext.Get("userToken")
	if !exist {
        failRes := &m.FailResponse{
            Msg: "Data dari token tidak ada",
        }
        reqContext.JSON(rc.TokenInvalid, failRes)
        return
	}

    // validasi bentuk data yang ada di token
    dataUser, ok := dataUser.(*m.UserData)
    if !ok {
        failRes := &m.FailResponse{
            Msg: "Bentuk data di token tidak valid",
        }
        reqContext.JSON(rc.TokenInvalid, failRes)
        return
    }

	reqContext.JSON(rc.Sukses, dataUser)
}
