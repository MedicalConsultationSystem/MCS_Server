package middleware

import (
	"MCS_Server/global"
	"MCS_Server/model/request"
	"MCS_Server/model/response"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithAll(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		j := &JWT{[]byte(global.MCS_Config.JWT.SigningKey)}
		_, err := j.ParseToken(token)
		if err != nil {
			response.FailWithAll(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		c.Next()
	}
}

var (
	TokenExpired = errors.New("令牌过期")
	TokenNotValidYet = errors.New("令牌不活跃")
	TokenMalformed   = errors.New("该字符串不属于令牌")
	TokenInvalid     = errors.New("无法处理这个令牌")
)

type JWT struct {
	SigningKey []byte
}

func (j *JWT) CreateToken(claims request.Claims) (string, error) {
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*request.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
