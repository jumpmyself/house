package model

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type UserToken struct {
	Uid int64

	jwt.RegisteredClaims
}

// 签名密钥

func GetJwt(Uid int64) (string, error) {
	if Uid < 0 {
		return "", errors.New("参数错误")
	}
	//读取签名密钥配置
	signKey := viper.GetString("jwt.signKey")
	token := &UserToken{
		Uid: Uid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "可求帅图书馆",                                            // 签发者
			Subject:   "名流张三",                                              // 签发对象
			Audience:  jwt.ClaimStrings{"Android", "IOS", "H5"},            //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),       //过期时间 1小时
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second * 0)), //最早使用时间 1秒之后
			IssuedAt:  jwt.NewNumericDate(time.Now()),                      //签发时间 当前时间
			ID:        "Test-1",                                            // jwt ID,类似于盐值 最好是每次都随机
		},
	}
	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, token).SignedString([]byte(signKey))
	//fmt.Println("生成的", tokenStr)
	return tokenStr, err
}

func CheckJwt(tokenStr string) (*UserToken, error) {
	//fmt.Println("收到的", tokenStr)
	signKey := viper.GetString("jwt.signKey")
	token, err := jwt.ParseWithClaims(tokenStr, &UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil //返回签名密钥
	})
	if err != nil || !token.Valid {
		fmt.Println("TOKEN不合格")
		return nil, errors.New("校验失败，TOKEN不合格")
	}

	claims, ok := token.Claims.(*UserToken)
	if !ok {
		fmt.Println("TOKEN转义失败")
		return nil, errors.New("TOKEN转义失败！")
	}

	return claims, nil
}

//func init() {
//	viper.SetConfigName("viper") // Name of the configuration file (without extension)
//	viper.SetConfigType("yaml")  // Type of the configuration file
//	viper.AddConfigPath(".")     // Path to the directory containing the configuration file
//
//	err := viper.ReadInConfig() // Read the configuration file
//	if err != nil {
//		panic(fmt.Errorf("failed to read configuration file: %s", err))
//	}
//}
