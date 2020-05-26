package model

type Product struct {
	Id           uint16 `json:"id"`
	Productname  string `json:"productName"`
	Productnum   uint16 `json:"productNum"`
	Productimage string `json:"productImage"`
	Producturl   string `json:"productUrl"`
}
