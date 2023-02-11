package models

type SaleItem struct {
	ID        int          `json:"id" grom:"primaryKey"`
	ProductID int          `json:"productID" form:"productID"`
	Price     string       `json:"price" form:"price"`
	Quantity  int          `json:"quantity" form:"quantity"`
	SaleID    int          `json:"saleID" form:"saleID"`
	Sale      SaleResponse `json:"sale"`
}

type SaleItemsResponse struct {
	ID        int    `json:"id"`
	ProductID int    `json:"productID" form:"productID"`
	Price     string `json:"price" form:"price"`
	Quantity  int    `json:"quantity" form:"quantity"`
	SaleID    int    `json:"-" form:"saleID"`
}

func (SaleItemsResponse) TableName() string {
	return "sale_items"
}
