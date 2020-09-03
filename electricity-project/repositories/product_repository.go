package repositories

import (
	"database/sql"
	"go-practice/electricity-project/common"
	"go-practice/electricity-project/datamodels"
	"strconv"
)

// 第一步，开发接口
// 第二步，实现接口

type IProduct interface {
	Conn() (error)
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) (bool)
	Update(product *datamodels.Product) (error)
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductMananger struct {
	table string
	mysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductMananger{table:table, mysqlConn:db}
}

func (p *ProductMananger) Conn() (err error) {
	if p.Conn() == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}

	if p.table == "" {
		p.table = "product"
	}
	return
}

// 插入
func (p *ProductMananger) Insert(product *datamodels.Product) (productId int64, err error) {
	if err := p.Conn(); err != nil {
		return 0, err
	}
	sql := "insert product set productName=?, productNum=?,productImage=?,productUrl=?"
	stmt, errSql := p.mysqlConn.Prepare(sql)
	if errSql != nil {
		return 0, errSql
	}

	result, errStmt := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if errStmt != nil {
		return 0, errStmt
	}

	return result.LastInsertId()
}

func (p *ProductMananger) Delete(productId int64) bool {
	if err := p.Conn(); err != nil {
		return false
	}

	sql := "delete from product where ID = ?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}

	_, err = stmt.Exec(productId)

	if err != nil {
		return false
	}

	return true
}

func (p *ProductMananger) Update(product *datamodels.Product) error {
	if err := p.Conn(); err != nil{
		return err
	}

	sql := "update product set productName=?, productNum=?, productImage=?, productUrl=? where ID" + strconv.FormatInt(product.ID, 10)
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProductMananger) SelectByKey(productId int64) (product *datamodels.Product, err error)  {
	if err := p.Conn(); err != nil {
		return &datamodels.Product{}, err
	}

	sql := "select * from product where ID = ?"
	rows, err := p.mysqlConn.Query(sql, strconv.FormatInt(productId, 10))
	defer rows.Close()
	if err != nil {
		return
	}
	result := common.GetResultRow(rows)
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}

	common.DataToStructByTagSql(result, product)
	return
}

func (p *ProductMananger) SelectAll()(products []*datamodels.Product, err error) {
	if err := p.Conn(); err != nil {
		return nil, err
	}

	sql := "select * from product"
	rows, err := p.mysqlConn.Query(sql)

	defer rows.Close()
	if err != nil {
		return
	}
	result := common.GetResultRows(rows)
	if (len(result) == 0) {
		return nil, nil
	}

	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		products = append(products, product)
	}
	return
}
