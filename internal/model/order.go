package model

type Order struct {
	No                int     `json:"no"`
	PlatformProductId string  `json:"platformProductId" binding:"required"`
	Qty               int     `json:"qty" binding:"required,min=1"`
	UnitPrice         float64 `json:"unitPrice" binding:"required,min=0"`
	TotalPrice        float64 `json:"totalPrice" binding:"required,min=0"`
}

type OrderRequest struct {
	Orders []Order `json:"orders" binding:"required,dive"`
}

type CleanedOrder struct {
	No         int     `json:"no"`
	ProductId  string  `json:"productId"`
	MaterialId string  `json:"materialId,omitempty"`
	ModelId    string  `json:"modelId,omitempty"`
	Qty        int     `json:"qty"`
	UnitPrice  float64 `json:"unitPrice"`
	TotalPrice float64 `json:"totalPrice"`
}

type CleanedOrderResponse struct {
	CleanedOrders []CleanedOrder
}
