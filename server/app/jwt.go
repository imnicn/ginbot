/*
 * @Author: yesxin(陶鑫)
 * @Github: https://github.com/yesxin
 * @Date: 2020-11-10 19:33:44
 */

package app

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

/*SigningKey 签名密匙*/
const SigningKey = "yesxin.com"

/*CreateJWT 生成一个JWT token*/
func CreateJWT(id uint32, sub string, exp int64) (string, error) {
	claims := &jwt.StandardClaims{
		Id:        fmt.Sprint(id), //用户id
		ExpiresAt: exp,            //过期时间
		Subject:   sub,            //主题
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SigningKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

/*ParseJWT 解析token*/
func ParseJWT(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
