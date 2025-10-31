package jwtUtils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jhinih/hin/hglobal"
	"time"
)

type TokenData struct {
	Userid   string `json:"user_id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}

// 统一签名
func GenToken(userid, username string, role int, exp time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"userid":   userid,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(exp).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(hglobal.Config.JWT.Secret))
}

// 封装：30 秒短效 ticket
func GenWSTicket(c *gin.Context) {
	uid := GetUserId(c)
	name := GetUserName(c)
	role := GetRole(c)

	ticket, err := GenToken(uid, name, role, 30*time.Second)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"ticket": ticket})
}

func GenAtoken(userid, username string, role int, exp time.Duration) (string, error) {
	return GenToken(userid, username, role, exp)
}

func GenRtoken(userid, username string, role int, exp time.Duration) (string, error) {
	return GenToken(userid, username, role, exp)
}

func IdentifyToken(tokenString string) (TokenData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名方法: %v", token.Header["alg"])
		}
		return []byte(hglobal.Config.JWT.Secret), nil
	})
	if err != nil {
		return TokenData{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// 验证token是否过期
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			return TokenData{}, fmt.Errorf("token已过期")
		}
	} else {
		// 解析失败
		return TokenData{}, fmt.Errorf("无效的token")
	}
	// 解析token成功
	return TokenData{
		Userid:   claims["userid"].(string),
		Username: claims["username"].(string),
		Role:     int(claims["role"].(float64)),
	}, nil
}

func GetUserId(c *gin.Context) string {
	if data, exists := c.Get(hglobal.TOKEN_USER_ID); exists {
		userId, ok := data.(string)
		if ok {
			return userId
		}
	}
	return ""
}

func GetUserName(c *gin.Context) string {
	if data, exists := c.Get(hglobal.TOKEN_USER_NAME); exists {
		userName, ok := data.(string)
		if ok {
			return userName
		}
	}
	return ""
}
func GetRole(c *gin.Context) int {
	if data, exists := c.Get(hglobal.TOKEN_ROLE); exists {
		role, ok := data.(int)
		if ok {
			return role
		}
	}
	return 0
}
