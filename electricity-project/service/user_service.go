package service

import (
	"database/sql"
	"fmt"
	"go-practice/electricity-project/datamodels"
	"go-practice/electricity-project/repositories"
	"strings"
)

type IUserService interface {
	IsPwdSuccess(userName string, pwd string) (user *datamodels.User, isOk bool)
	AddUser(user *datamodels.User) (userId int64, err error)
	IsLoginSuccess(userName string, uiPwd string) (user *datamodels.User, isOk bool)
}

func NewUserService(table string, db *sql.DB) IUserService {
	return &UserService{userRepo: repositories.NewUserRepository(table, db)}
}

type UserService struct {
	userRepo repositories.IUserRepository
}

func (u UserService) IsPwdSuccess(userName string, pwd string) (user *datamodels.User, isOk bool) {
	user, err := u.userRepo.Select(userName)
	if err != nil {
		return &datamodels.User{}, false
	}
	if ValidatePassword(pwd, user.Password) {
		return user, true
	} else {
		return user, false
	}
}

func ValidatePassword(userPwd, hashed string) (isOk bool) {
	return strings.EqualFold(userPwd, hashed)
	//if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPwd)); err != nil {
	//	return false, errors.New("密码对比错误")
	//}
}

func (s *UserService) IsLoginSuccess(userName string, uiPwd string) (user *datamodels.User, isOk bool) {
	if userName == "" || uiPwd == "" {
		return nil, false
	}
	user, err := s.userRepo.Select(userName)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	if ValidatePassword(user.Password, uiPwd) {
		return user, true
	}
	return
}

func (u UserService) AddUser(user *datamodels.User) (userId int64, err error) {
	return u.userRepo.Insert(user)
}
