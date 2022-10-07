package entity

type Item struct {
	ID          uint   `json:"id"`
	Item_code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_id    uint   `json:"order_id"`
}
