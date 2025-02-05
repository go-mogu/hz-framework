package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtPayload struct {
	User interface{} `json:"user"`
	jwt.RegisteredClaims
}

// GenerateJwtToken 生成jwt token
func GenerateJwtToken(secret string, expire int64, user interface{}, issuer string) (string, error) {
	data := JwtPayload{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire) * time.Hour)), // 过期时间 7天  配置文件
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                        // 签名发行时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                        // 签名生效时间
			Issuer:    issuer,                                                                // 签名的发行者
		},
	}
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	token, err := j.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("jwt 生成token失败" + err.Error())
	}
	return token, nil
}

// ParseJwtToken 解析 jwt token
func ParseJwtToken(jwtToken string, secret string) (*JwtPayload, error) {
	if jwtToken == "" {
		return nil, errors.New("token 为空")
	}
	token, err := jwt.ParseWithClaims(jwtToken, &JwtPayload{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("jwt token 解析失败" + err.Error())
	}
	if claims, ok := token.Claims.(*JwtPayload); ok && token.Valid {
		return claims, nil
	} else {
		return claims, errors.New("jwt 解析验证后失败")
	}
}
