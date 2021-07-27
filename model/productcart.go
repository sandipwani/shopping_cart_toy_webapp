package model

import (
	"home/webonise/go/src/shopping_cart_app/structure"
	"log"
)

type ProductCartAccessor interface {
	RetriveCartProducts() (productlist []structure.ProductCart, err error)
	InsertCartProduct(productid int, productname string, productprice float64, productquantity int) error
	DeleteCartProduct(productid int) error
	UpdateCartProduct(productid int, productquantity int) error
	RetriveCartProductByID(productid int) (product_info structure.ProductCart, err error)
}

func (w *WriteDbAccess) RetriveCartProducts() (productlist []structure.ProductCart, err error) {
	rows, err := w.Db.Query("SELECT *FROM product_cart;")
	if err != nil {
		log.Printf("failed to fetched the product data from cart : %v",err)
	}

	defer rows.Close()

	var productitem structure.ProductCart
	for rows.Next() {
		rows.Scan(&productitem.ProductId, &productitem.ProductName, &productitem.ProductPrice, &productitem.ProductQuantity)
		productlist = append(productlist, productitem)
	}
	return productlist, err
}

func (w *WriteDbAccess) InsertCartProduct(productid int, productname string, productprice float64, productquantity int) error {
	sqlStatement := `INSERT INTO product_cart (productid, productname, productprice, productquantity) VALUES ($1, $2, $3, $4);`
	_, err := w.Db.Exec(sqlStatement, productid, productname, productprice, productquantity)
	return err
}

func (w *WriteDbAccess) DeleteCartProduct(productid int) error {
	sqlStatement := `DELETE FROM product_cart WHERE productid = $1;`
	_, err := w.Db.Exec(sqlStatement, productid)
	return err
}

func (w *WriteDbAccess) UpdateCartProduct(productid int, productquantity int) error {
	sqlStatement := `UPDATE product_cart SET productquantity = $2 WHERE productid = $1;`
	_, err := w.Db.Exec(sqlStatement, productid, productquantity)
	return err
}

func (w *WriteDbAccess) RetriveCartProductByID(productid int) (product_info structure.ProductCart, err error) {
	err = w.Db.QueryRow("SELECT *FROM product_cart WHERE productid = $1;", productid).Scan(&product_info.ProductId, &product_info.ProductName, &product_info.ProductPrice, &product_info.ProductQuantity)
	return product_info, err
}
