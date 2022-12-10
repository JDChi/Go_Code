package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var (
	eccError  error
	eccPriKey *ecdsa.PrivateKey
	eccPubKey *ecdsa.PublicKey
)

type EcdsaUser struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

type EcdsaClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func getEcdsaKey(keyType int) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	var err error
	var priKey *ecdsa.PrivateKey
	var pubKey *ecdsa.PublicKey
	var curve elliptic.Curve // 椭圆曲线
	switch keyType {
	case 1:
		curve = elliptic.P224()
	case 2:
		curve = elliptic.P256()
	case 3:
		curve = elliptic.P384()
	case 4:
		curve = elliptic.P521()
	default:
		return nil, nil, errors.New("select curve error")
	}
	// 生成公私密钥对
	priKey, err = ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	// 从私钥中拿出公钥
	pubKey = &priKey.PublicKey
	return priKey, pubKey, nil
}

func init() {
	eccPriKey, eccPubKey, eccError = getEcdsaKey(2)
	if eccError != nil {
		panic(eccError)
		return
	}

}

func main() {
	r := gin.Default()
	r.POST("/getToken3", func(context *gin.Context) {
		u := EcdsaUser{}
		err := context.Bind(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, "params error")
			return
		}
		token, err := ecdsaReleaseToken(u)
		if err != nil {
			context.JSON(http.StatusBadRequest, "gen token error")
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "token succeed",
			"data": token,
		})

	})
	r.POST("/checkToken3", ecdsaTokenMiddle(), func(context *gin.Context) {
		context.JSON(http.StatusOK, "auth succeed")
	})
	r.Run(":9090")

}

func ecdsaTokenMiddle() gin.HandlerFunc {
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
		claims, err := ecdsaJwtTokenRead(realTokenString)
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

		u := EcdsaUser{}
		context.Bind(&u)
		id := claimsValue["user_id"].(string)
		if u.ID != id {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized, "msg": "user is not exist",
			})
			context.Abort()
			return
		}
		context.Next()
	}
}

func ecdsaJwtTokenRead(tokenString string) (any, error) {

	// 解析出 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// value , ok := x.(T) 这是断言的语法
		// 校验签名的方法是不是 rsa
		if _, OK := token.Method.(*jwt.SigningMethodECDSA); !OK {
			return nil, fmt.Errorf("parse error")
		}
		// 是就返回公钥
		return eccPubKey, nil
	})
	// value , ok := x.(T) 这是断言的语法
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func ecdsaReleaseToken(u EcdsaUser) (any, error) {
	claims := &EcdsaClaims{
		UserId: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			Issuer:    "GDCA",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedString, err := token.SignedString(eccPriKey)
	return signedString, err
}
