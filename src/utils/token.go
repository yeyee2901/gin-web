package utils

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	ctx "github.com/yeyee2901/gin-web/src/context"
	m "github.com/yeyee2901/gin-web/src/model"
)

type TokenClaim struct {
	jwt.StandardClaims
	*m.UserData
}

var TOKEN_SIGN_METHOD = jwt.SigningMethodHS256

// ComposeToken godoc
// Helper function untuk membuat token berdasarkan struct userData
func ComposeToken(config *ctx.Config, userData *m.UserData) (signedToken string, err error) {
	apiKey := []byte(config.ApiKey.Key)
	expired := time.Duration(config.Token.ExpTimeAccess) * time.Second

	claims := TokenClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Artaku MyClient",
			ExpiresAt: time.Now().Add(expired).Unix(),
		},
		UserData: userData,
	}

	_token := jwt.NewWithClaims(TOKEN_SIGN_METHOD, claims)

	// sign token nya
	signedToken, err = _token.SignedString(apiKey)

	return
}

// DecomposeToken godoc
// Helper function untuk mengekstrak data dari token string yang dibuat dengan ComposeToken()
func DecomposeToken(config *ctx.Config, tokenString string) (claims *m.UserData, err error) {
	_token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		method, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("Token invalid. Signing method tidak sesuai")
		}

		if method != TOKEN_SIGN_METHOD {
			return nil, fmt.Errorf("Token invalid. Signing method tidak sesuai")
		}

		// callback ini harus return api key nya (dalam byte)
		return []byte(config.ApiKey.Key), nil
	})

	if err != nil {
		return nil, err
	}

	_claims, ok := _token.Claims.(jwt.MapClaims)

	if !ok || !_token.Valid {
		return nil, fmt.Errorf("Token tidak valid")
	}

    // parse claim data nya
    fmt.Println(_claims)
    claims = &m.UserData{
        Nama: _claims["nama"].(string),
    }

	return
}
