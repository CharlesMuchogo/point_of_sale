package mystructs

type Product struct {
	Id                 uint   `gorm:"primaryKey" json:"id"`
	ProductName        string `json:"name"`
	SerialNumber       string `json:"serial_number"`
	ProductQuantity    int    `json:"quantity"`
	ProductPrice       int    `json:"price"`
	ProductImage       string `json:"image"`
	CategoryId         int    `json:"category"`
	ProductDescription string `json:"description"`
}

type ProductCategories struct {
	Id            uint   `gorm:"primaryKey" json:"id"`
	CategoryName  string `json:"name"`
	CategoryImage string `json:"image"`
}
