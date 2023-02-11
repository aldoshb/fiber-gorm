package models

type Customer struct {
	ID      int          `gorm:"primary_key" json:"id" form:"id"`
	Name    string       `json:"name" form:"name" gorm:"not null"`
	Address string       `json:"address" form:"address"`
	Phone   string       `json:"phone" form:"phone"`
	Sale    SaleResponse `json:"sale"`
}

type CustomerResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (CustomerResponse) TableName() string {
	return "Customers"
}
