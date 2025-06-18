package port

import "bewell-backend-challenge/internal/model"

type CleanOrderUsecase interface {
	CleanOrders(orders model.OrderRequest) (*model.CleanedOrderResponse, error)
}
