package user

import (
	myJwt "firstGin/middleware/jwt"
	"firstGin/models/request"
	"firstGin/models/response"
	"firstGin/pkg/gredis"
	userService2 "firstGin/service/userService"
	"fmt"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func userLogin(c *gin.Context) {
	json := request.User{}
	_ = c.Bind(&json)
	userService := userService2.User{Username: json.Name, Password: json.Password}
	isExist, err := userService.Check()
	if err == nil && isExist {
		role, err := userService.GetRole()
		if err == nil {
			generateToken(c, json, role)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 401,
				"msg":    "其他错误",
				"data":   err,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"msg":    "账号或密码错误",
			"data":   "",
		})
	}

}

func generateToken(c *gin.Context, user request.User, role int) {
	j := &myJwt.JWT{
		SigningKey: []byte("GPK&ZMZ"),
	}
	claims := myJwt.CustomClaims{
		Name: user.Name,
		Role: role,
		StandardClaims: jwtGo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600, // 过期时间 一小时
			Issuer:    "GPK",                    //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	err1 := gredis.Set(user.Name, token, 1800000)
	if err1 != nil {
		fmt.Println("redis操作错误")
	}
	log.Println(token)
	//data := LoginResult{
	//	User:  user,
	//	Token: token,
	//}
	data := response.LoginResult{
		User: user, Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}
