package product

import (
	"pos/model/category"
	"pos/model/vendor"
	"time"
)

type Product struct {
	ID				  int `json:"id"`
	CategoryID	int `json:"category_id"`
	VendorID		int `json:"vendor_id"`
	Name 				string `json:"name"`
	Describtion 				string `json:"describtion"`
	Sku 				string `json:"sku"`
	Price 			float64 `json:"price"`
	Stock 			int `json:"stock"`
	Status 			string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Vendor vendor.Vendor
	Category category.Category
}




