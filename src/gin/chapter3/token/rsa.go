package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type RSAUser struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

type RSAClaim struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

var (
	resPrivateKey  []byte
	resPublicKey   []byte
	priErr, pubErr error
)

func init() {

	resPrivateKey, priErr = ioutil.ReadFile("./src/chapter3/token/pri.pem")
	resPublicKey, pubErr = ioutil.ReadFile("./src/chapter3/token/pub.pem")
	if priErr != nil || pubErr != nil {
		panic(fmt.Sprintf("read pem file error: %s %s", priErr, pubErr))
		return
	}

}

func main() {
	r := gin.Default()
	r.POST("/getToken2", func(context *gin.Context) {
		u := RSAUser{}
		err := context.Bind(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, "params error")
			return
		}
		token, err := rsaReleaseToken(u)
		if err != nil {
			context.JSON(http.StatusInternalServerError, "gen token error")
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "token succeed",
			"data": token,
		})
	})
	r.POST("/checkToken2", rsaTokenMiddle(), func(context *gin.Context) {
		context.JSON(http.StatusOK, "auth succeed")
	})
	r.Run(":9090")

}

func rsaTokenMiddle() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := "Bear"
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, auth+":") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "prefix error"})
			context.Abort()
			return
		}
		index := strings.Index(tokenString, auth+":")
		realTokenString := tokenString[index+len(auth)+1:]
		claims, err := rsaJwtTokenRead(realTokenString)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "cert invalid"})
			context.Abort()
			return
		}
		claimsValue := claims.(jwt.MapClaims)
		if claimsValue["user_id"] == nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized, "msg": "user is not exist",
			})
			context.Abort()
			return
		}
		u := RSAUser{}
		context.Bind(&u)
		id := claimsValue["user_id"].(string)
		if u.Id != id {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized, "msg": "user is not exist",
			})
			context.Abort()
			return
		}
		context.Next()
	}
}

func rsaJwtTokenRead(tokenString string) (any, error) {
	// 拿出公钥
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(resPublicKey)
	if err != nil {
		return nil, err
	}

	// 解析出 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// value , ok := x.(T) 这是断言的语法
		// 校验签名的方法是不是 rsa
		if _, OK := token.Method.(*jwt.SigningMethodRSA); !OK {
			return nil, fmt.Errorf("parse error")
		}
		// 是就返回公钥
		return publicKey, err
	})
	// value , ok := x.(T) 这是断言的语法
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// 分发 token
func rsaReleaseToken(u RSAUser) (any, error) {
	return rsaJwtGen(u.Id)
}

// 生成 rsa 签名的 token
func rsaJwtGen(id string) (any, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(resPrivateKey)
	if err != nil {
		return nil, err
	}

	claim := &RSAClaim{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			Issuer:    "GDCA",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedString, err := token.SignedString(privateKey)
	return signedString, err
}
