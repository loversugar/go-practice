package repositories

import (
	"database/sql"
	"fmt"
	"go-practice/electricity-project/common"
	"go-practice/electricity-project/datamodels"
	"strconv"
)

// 第一步，开发接口
// 第二步，实现接口

type IProduct interface {
	Conn() error
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(product *datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductManager struct {
	Table     string
	MysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{Table: table, MysqlConn: db}
}

func (p *ProductManager) Conn() (err error) {
	if p.MysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.MysqlConn = mysql
	}

	if p.Table == "" {
		p.Table = "product"
	}
	return
}

func (p *ProductManager) Insert(product *datamodels.Product) (id int64, err error) {
	//判断连接是否存在
	if err = p.Conn(); err != nil {
		return 0, err
	}
	//准备sql
	sql := fmt.Sprintf("INSERT %s SET productName=?,productNum=?,productImage=?,productUrl=?", p.Table)
	if stmt, err := p.MysqlConn.Prepare(sql); err != nil {
		return 0, err
	} else if result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl); err != nil {
		return 0, err
	} else {
		//defer stmt.Close()
		return result.LastInsertId()
	}

}

func (p *ProductManager) Delete(productId int64) bool {
	if err := p.Conn(); err != nil {
		return false
	}

	sql := "delete from product where ID = ?"
	stmt, err := p.MysqlConn.Prepare(sql)
	if err != nil {
		return false
	}

	_, err = stmt.Exec(productId)

	if err != nil {
		return false
	}

	return true
}

func (p *ProductManager) Update(product *datamodels.Product) error {
	if err := p.Conn(); err != nil {
		return err
	}

	sql := "update product set productName=?, productNum=?, productImage=?, productUrl=? where ID = " + strconv.FormatInt(product.ID, 10)
	stmt, err := p.MysqlConn.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProductManager) SelectByKey(productId int64) (product *datamodels.Product, err error) {
	if err := p.Conn(); err != nil {
		return &datamodels.Product{}, err
	}

	sql := "select * from product where ID = ?"
	rows, err := p.MysqlConn.Query(sql, strconv.FormatInt(productId, 10))
	defer rows.Close()
	if err != nil {
		return
	}
	result := common.GetResultRow(rows)
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}

	product = &datamodels.Product{}
	common.DataToStructByTagSql(result, product)
	return
}

func (p *ProductManager) SelectAll() (products []*datamodels.Product, err error) {
	if err := p.Conn(); err != nil {
		return nil, err
	}

	sql := "select * from product"
	rows, err := p.MysqlConn.Query(sql)

	defer rows.Close()
	if err != nil {
		return
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}

	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		products = append(products, product)
	}
	return
}
