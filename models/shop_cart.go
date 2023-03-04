package models

type ShopCartPrimaryKey struct {
	Id string `json:"id"`
}

type ClientTotalPrice struct {
	Number string  `json:"number"`
	Name   string  `json:"name"`
	Total  float64 `json:"total"`
	Date   string  `json:"date"`
}
type Top10SellProducts struct {
	Name  string  `json:"name"`
	Total float64 `json:"total"`
}
type ShopCartFull struct {
	Number  int     `json:"number"`
	Name    string  `json:"name"`
	Product string  `json:"product"`
	Price   float64 `json:"price"`
	Count   int     `json:"count"`
	Total   float64 `json:"total"`
	Time    string  `json:"time"`
}

type CategoryCountList struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
type ShopCart struct {
	Id        string `json:"id"`
	ProductId string `json:"productId"`
	UserId    string `json:"userID"`
	Count     int    `json:"count"`
	Date      string `json:"date"`
	Status    bool   `json:"status"`
}

type Add struct {
	ProductId string `json:"productId"`
	UserId    string `json:"userID"`
	Count     int    `json:"count"`
	Date      string `json:"date"`
}

type Remove struct {
	ProductId string `json:"productId"`
	UserId    string `json:"userID"`
}

// GetAllShopCart
type GetAllShopCartRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetAllShopCartsResponse struct {
	Count     int `json:"count"`
	ShopCarts []ShopCart
}
