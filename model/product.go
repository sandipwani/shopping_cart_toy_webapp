package model

import (
	"home/webonise/go/src/shopping_cart_app/structure"
	"log"

	_ "github.com/lib/pq"
)

type ProductAccessor interface {
	InsertShopProduct(productid int, productname string, productprice float64) error
	RetriveShopProducts() (productlist []structure.Product, err error)
	RetriveShopProductByID(productid int) (product_info structure.ProductCart, err error)
	DeleteShopProduct(productid int) error
	UpdateShopProduct(productid int, productprice float64 ) error
}

func (w *WriteDbAccess) InsertShopProduct(productid int, productname string, productprice float64) error {
	sqlStatement := `INSERT INTO product (prod_id, prod_name, prod_price) VALUES ($1, $2, $3);`
	_, err := w.Db.Exec(sqlStatement, productid, productname, productprice)
	return err
}

func (w *WriteDbAccess) RetriveShopProducts() (productlist []structure.Product, err error) {
	rows, err := w.Db.Query("SELECT *FROM product;")
	if err != nil {
		log.Printf("failed to fetch products from shop %v", err)
	}

	defer rows.Close()
	var productitem structure.Product
	for rows.Next() {

		rows.Scan(&productitem.ProdId, &productitem.ProdName, &productitem.ProdPrice)
		productlist = append(productlist, productitem)
	}

	return productlist, err
}

func (w *WriteDbAccess) RetriveShopProductByID(productid int) (product_info structure.ProductCart, err error) {

	err = w.Db.QueryRow("SELECT *FROM product WHERE prod_id = $1;", productid).Scan(&product_info.ProductId, &product_info.ProductName, &product_info.ProductPrice)
	return product_info, err
}

func (w *WriteDbAccess) DeleteShopProduct(productid int) error {
	sqlStatement := `DELETE FROM product WHERE prod_id = $1;`
	_, err := w.Db.Exec(sqlStatement, productid)
	return err
}

func (w *WriteDbAccess) UpdateShopProduct(productid int, productprice float64 ) error {
	sqlStatement := `UPDATE product SET prod_price = $2 WHERE prod_id = $1;`
	_, err := w.Db.Exec(sqlStatement, productid, productprice)
	return err
}