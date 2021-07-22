package structure

type Product struct {
	ProdId    int     `json:"prod_id"`
	ProdName  string  `json:"prod_name"`
	ProdPrice float64 `json:"prod_price"`
}

type ProductCart struct {
	ProductId       int     `json:"productid"`
	ProductName     string  `json:"productname"`
	ProductPrice    float64 `json:"productprice"`
	ProductQuantity int     `json:"productquantity"`
}
