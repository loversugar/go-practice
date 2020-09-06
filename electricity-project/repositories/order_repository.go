package repositories

import (
	"database/sql"
	"go-practice/electricity-project/common"
	"go-practice/electricity-project/datamodels"
)

type IOrderRepository interface {
	Conn() error
	Insert(*datamodels.Order) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

func NewOrderManagerRepository(table string, db *sql.DB) IOrderRepository {
	return &OrderManagerRepository{table: table, mysqlConn: db}
}

type OrderManagerRepository struct {
	table     string
	mysqlConn *sql.DB
}

func (o *OrderManagerRepository) Conn() (err error) {
	if o.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}

		o.mysqlConn = mysql
	}

	return
}

func (o *OrderManagerRepository) Insert(order *datamodels.Order) (productId int64, err error) {
	if err = o.Conn(); err != nil {
		return
	}

	sql := "insert `order` set userId = ? , productId = ?, orderStatus = ?"
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}

	result, err := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if err != nil {
		return
	}

	return result.LastInsertId()
}

func (o *OrderManagerRepository) Delete(orderId int64) bool {
	if err := o.Conn(); err != nil {
		return false
	}

	sql := "delete from `" + o.table + "` where ID = ?"
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}

	_, err = stmt.Exec(orderId)

	if err != nil {
		return false
	}

	return true
}

func (o *OrderManagerRepository) Update(order *datamodels.Order) (err error) {
	if err = o.Conn(); err != nil {
		return
	}

	sql := "update `" + o.table + "` set userId = ?, productId = ?, orderStatus = ? where productId = ?"
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}

	_, err = stmt.Exec(order.UserId, order.ProductId, order.OrderStatus, order.ID)
	if err != nil {
		return
	}
	return
}

func (o *OrderManagerRepository) SelectByKey(id int64) (order *datamodels.Order, err error) {
	if err = o.Conn(); err != nil {
		return
	}

	sql := "select * from " + o.table + " where ID = ?"
	rows, err := o.mysqlConn.Query(sql, id)
	defer rows.Close()
	if err != nil {
		return
	}

	result := common.GetResultRow(rows)
	if len(result) == 0 {
		return &datamodels.Order{}, nil
	}

	common.DataToStructByTagSql(result, order)
	return
}

func (o *OrderManagerRepository) SelectAll() (orders []*datamodels.Order, err error) {
	if err = o.Conn(); err != nil {
		return
	}

	sql := "select * from `" + o.table + "`"
	rows, err := o.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return
	}

	results := common.GetResultRows(rows)
	if len(results) == 0 {
		return nil, nil
	}

	for _, value := range results {
		order := &datamodels.Order{}
		common.DataToStructByTagSql(value, order)
		orders = append(orders, order)
	}

	return
}

func (o *OrderManagerRepository) SelectAllWithInfo() (orders map[int]map[string]string, err error) {
	if err = o.Conn(); err != nil {
		return
	}

	sql := "select o.ID, p.productName, o.orderStatus from `order` as o left join product as p on o.productId=p.ID"
	rows, err := o.mysqlConn.Query(sql)
	if err != nil {
		return
	}

	results := common.GetResultRows(rows)
	if len(results) <= 0 {
		return nil, nil
	}

	return results, nil
}
