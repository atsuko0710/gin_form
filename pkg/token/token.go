package token

import (
	"gin_forum/pkg/response"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Claims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 24 * 365 // 过期时间

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return viper.GetString("jwt_secret"), nil
}

// GenToken 生成access token 和 refresh token
func GenToken(userID int64, username string) (aToken string, rToken string, err error) {
	// 创建一个我们自己的声明
	c := Claims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "atsukodan",                                // 签发人
		},
	}
	jwtSecret := viper.GetString("jwt_secret")
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString([]byte(jwtSecret))
	if err != nil {
		zap.L().Error("jwt.NewWithClaims() accesstoken failed", zap.Error(err))
	}

	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
		Issuer:    "atsukodan",                             // 签发人
	}).SignedString([]byte(jwtSecret))
	if err != nil {
		zap.L().Error("jwt.NewWithClaims() refreshtoken failed", zap.Error(err))
	}
	
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return aToken, rToken, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *Claims, resCode response.ResCode) {
	// 解析token
	var token *jwt.Token
	claims = new(Claims)
	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	resCode = response.OK
	if err != nil {
		zap.L().Error("jwt.ParseWithClaims() filed", zap.Error(err))
		resCode = response.InvalidToken
	}
	if !token.Valid { // 校验token
		resCode = response.InvalidToken
	}
	return claims, resCode
}
