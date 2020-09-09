package main

import (
	"errors"
	"fmt"
	"go-practice/electricity-project/common"
	"go-practice/electricity-project/encrypt"
	"net/http"
	"sync"
)

var hostArray = []string{"ip"}

// 用来存放控制信息
type AccessControl struct {
	// 用来存放用户想要存放的信息
	sourceArray map[int]interface{}
	*sync.RWMutex
}

var accessControl = &AccessControl{sourceArray: make(map[int]interface{})}

// 获取指定数据
func (m *AccessControl) GetNewRecord(uid int) interface{} {
	m.RLock()
	defer m.RUnlock()
	data := m.sourceArray[uid]
	return data
}

// 设置记录
func (m *AccessControl) SetNewRecord(uid int) {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	m.sourceArray[uid] = "hello imooc"
}

func (m *AccessControl) GetDistributeRight(req *http.Request) bool {
	uid, err := req.Cookie("uid")
	if err != nil {
		return false
	}
	fmt.Println(uid)
	return true
}

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
	hashConsistent := common.NewConsistent()
	// 采用一致性hash，添加节点
	for _, v := range hostArray {
		hashConsistent.Add(v)
	}

	filter := common.NewFilter()
	filter.RegisterFilterUri("/check", Auth)

	http.HandleFunc("/check", filter.Handle(Check))

	http.ListenAndServe(":8083", nil)
}
