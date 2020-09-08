package main

import (
	"errors"
	"fmt"
	"go-practice/electricity-project/common"
	"go-practice/electricity-project/encrypt"
	"net/http"
)

func Check(w http.ResponseWriter, r *http.Request) {
	fmt.Println("执行check")
}

func CheckUserInfo(req *http.Request) error {
	uidCookie, err := req.Cookie("uid")
	if err != nil {
		return errors.New("用户未登录！")
	}
	signCookie, err := req.Cookie("sign")
	if err != nil {
		return errors.New("用户加密串获取失败！")
	}
	deSignByte, err := encrypt.DePwdCode(signCookie.Value)
	if err != nil {
		return errors.New("加密串已被篡改！")
	}
	fmt.Println("结果比对：uid=", uidCookie.Value, " sign: ", string(deSignByte))
	if CheckIdInfo(uidCookie.Value, string(deSignByte)) {
		return nil
	}
	return errors.New("身份校验失败!")
}

func CheckIdInfo(checkStr string, signStr string) bool {
	if checkStr == signStr {
		return true
	}
	return false
}

// 统一验证拦截器
func Auth(w http.ResponseWriter, r *http.Request) error {
	return CheckUserInfo(r)
}

func main() {
	filter := common.NewFilter()
	filter.RegisterFilterUri("/check", Auth)

	http.HandleFunc("/check", filter.Handle(Check))

	http.ListenAndServe(":8083", nil)
}
