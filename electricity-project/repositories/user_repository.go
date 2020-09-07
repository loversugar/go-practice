package repositories

import (
	"database/sql"
	"errors"
	"go-practice/electricity-project/common"
	"go-practice/electricity-project/datamodels"
	"strconv"
)

type IUserRepository interface {
	Conn() (err error)
	Select(username string) (user *datamodels.User, err error)
	Insert(user *datamodels.User) (userId int64, err error)
}

func NewUserRepository(table string, db *sql.DB) IUserRepository {
	return &UserManagerRepository{table: table, mysqlConn: db}
}

type UserManagerRepository struct {
	table     string
	mysqlConn *sql.DB
}

func (u UserManagerRepository) Conn() (err error) {
	if u.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		u.mysqlConn = mysql
	}

	if u.table == "" {
		u.table = "user"
	}
	return
}

func (u UserManagerRepository) Select(username string) (user *datamodels.User, err error) {
	if username == "" {
		return &datamodels.User{}, errors.New("条件不能为空！")
	}
	if err = u.Conn(); err != nil {
		return
	}

	sql := "select * from `" + u.table + "` where userName = ?"
	row, err := u.mysqlConn.Query(sql, username)
	defer row.Close()
	if err != nil {
		return
	}
	data := common.GetResultRow(row)
	if len(data) <= 0 {
		return &datamodels.User{}, err
	}

	user = &datamodels.User{}
	common.DataToStructByTagSql(data, user)
	return user, nil
}

func (u UserManagerRepository) Insert(user *datamodels.User) (userId int64, err error) {
	if err = u.Conn(); err != nil {
		return
	}

	sql := "insert `" + u.table + "` set nickName = ?, userName = ?, `password` = ?"
	stmt, err := u.mysqlConn.Exec(sql, user.NickName, user.UserName, user.Password)

	if err != nil {
		return
	}

	return stmt.LastInsertId()
}

func (u UserManagerRepository) SelectById(userId int64) (user *datamodels.User, err error) {
	if err = u.Conn(); err != nil {
		return
	}
	sql := "select * from `" + u.table + "` where ID = " + strconv.FormatInt(userId, 10)
	row, err := u.mysqlConn.Query(sql)
	if err != nil {
		return
	}

	result := common.GetResultRow(row)
	if len(result) <= 0 {
		return
	}

	user = &datamodels.User{}
	common.DataToStructByTagSql(result, user)
	return
}
