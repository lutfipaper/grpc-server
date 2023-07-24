package app

type ProductList struct {
	UUID        string  `json:"uuid,omitempty" gorm:"column:uuid"`
	Name        string  `json:"name,omitempty" gorm:"column:lvl_id"`
	Description string  `json:"description,omitempty" gorm:"column:description"`
	SalesPrice  float32 `json:"sales_price,omitempty" gorm:"column:sales_price"`
}
