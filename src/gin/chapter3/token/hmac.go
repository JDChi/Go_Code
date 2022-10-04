package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

type HmacUser struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

type MyClaims struct {
	UserId string
	jwt.RegisteredClaims
}

// go get -u github.com/golang-jwt/jwt/v4
// jwt 里是有结构体，我们可以在这个基础上进行拓展，然后对这些数据生成 token，下发给客户端
// 当服务端接收到客户端上传的 token 后，就将 token 还原回 jwt，取出里面的结构体数据来进行比较
func main() {
	r := gin.Default()
	// 分发 token
	r.POST("/getToken1", func(context *gin.Context) {
		var u HmacUser
		context.Bind(&u)
		token, err := hmacReleaseToken(u)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "token succeed",
			"data": token,
		})
	})
	// 验证 token
	r.POST("/checkToken1", hmacAuthMiddleware(), func(context *gin.Context) {
		context.JSON(http.StatusOK, "auth succeed")
	})
	r.Run(":9090")
}

func hmacAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 发送 token 的时候，加个 “Bear:”的前缀，后面才是实际的 token
		auth := "Bear"
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, auth+":") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "prefix error"})
			context.Abort()
			return
		}
		index := strings.Index(tokenString, auth+":")
		realTokenString := tokenString[index+len(auth)+1:]
		jwtToken, claims, err := hmacParseToken(realTokenString)
		if err != nil || !jwtToken.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "certificate invalid",
			})
			context.Abort()
			return
		}

		// 判断解析出 token 里的结构体里的用户 id 跟接口上传的 uid 是否一致
		var u HmacUser
		context.Bind(&u)
		if u.Id != claims.UserId {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "user is not exist",
			})
			context.Abort()
			return
		}
		// token 判定有效
		context.Next()
	}
}

// 解析 token
func hmacParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	jwtToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return jwtToken, claims, err
}

// 模拟证书签名的密钥
var jwtKey = []byte("a_secret_key")

// 分发 token
func hmacReleaseToken(u HmacUser) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &MyClaims{
		UserId: u.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Subject:   "user token",
			Issuer:    "Jack",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
