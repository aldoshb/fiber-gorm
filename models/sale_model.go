package models

type Sale struct {
	ID          int                 `json:"id" form:"id" gorm:"primary"`
	SaleCode    string              `json:"saleCode" form:"salecode" gorm:"string;unique;not null"`
	DateCreated string              `json:"dateCreated" form:"dateCreated" gorm:"string"`
	Status      string              `json:"status" form:"status" gorm:"string"` // draft/cancel/unpaid/paid
	CustomerID  int                 `json:"customerID" form:"customerID" gorm:"not null"`
	Customer    CustomerResponse    `json:"customer"`
	SaleItems   []SaleItemsResponse `json:"items"`
}

type SaleResponse struct {
	ID          int    `json:"id"`
	SaleCode    string `json:"saleCode"`
	DateCreated string `json:"dateCreated" form:"dateCreated" gorm:"string"`
	Status      string `json:"status"`
	CustomerID  int    `json:"-"`
}

func (SaleResponse) TableName() string {
	return "Sales"
}
