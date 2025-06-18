package handler

import (
	"bewell-backend-challenge/internal/app/port"
	"bewell-backend-challenge/internal/app/usecase"
	"bewell-backend-challenge/internal/model"
	"bewell-backend-challenge/util/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase port.CleanOrderUsecase
}

func New() *Handler {
	usecase := usecase.New()
	return &Handler{usecase}
}

func (h *Handler) CleanOrders(context *gin.Context) {
	var request model.OrderRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		response.HandleBadRequest(context, err)
		return
	}

	result, err := h.usecase.CleanOrders(request)
	if err != nil {
		response.HandleError(context, response.GetStatusCode(err), model.Message{
			Error: &model.Error{
				ErrorCode: response.GetErrorCode(err),
				Message:   response.GetMessage(err),
			},
		})
		return
	}

	response.HandleSuccess(context, http.StatusOK, result.CleanedOrders)
}
