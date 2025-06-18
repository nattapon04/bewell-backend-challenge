package usecase

import (
	"bewell-backend-challenge/internal/app/port"
	"bewell-backend-challenge/internal/constant"
	"bewell-backend-challenge/internal/model"
	"bewell-backend-challenge/util/response"
	"errors"
	"strconv"
	"strings"
)

var replacementMap = map[string]string{
	constant.FG0A: constant.WipingCloth,
	constant.FG05: constant.ETC,
}

type cleanOrderUsecase struct {
}

func New() port.CleanOrderUsecase {
	return &cleanOrderUsecase{}
}

func (u *cleanOrderUsecase) CleanOrders(orders model.OrderRequest) (*model.CleanedOrderResponse, error) {
	var cleanedOrders []model.CleanedOrder
	nextOrderNo := 1
	filmTypePartQuantities := make(map[string]int)
	texturePartQuantities := make(map[string]int)

	for _, order := range orders.Orders {
		platformProductIDs := strings.Split(order.PlatformProductId, constant.ForwardSlash)
		cleanedPlatformProductIDs, QuantityItems, totalQuantity := cleanData(platformProductIDs)
		unitPricePerItem := order.UnitPrice / float64(totalQuantity)
		totalPricePerItem := order.TotalPrice / float64(totalQuantity*order.Qty)

		for index, cleanedPlatformProductID := range cleanedPlatformProductIDs {
			productQty := QuantityItems[index] * order.Qty
			finalProductID := cleanedPlatformProductID
			materialID, modelID, err := extractMaterialIDAndModelID(finalProductID)
			if err != nil {
				continue
			}

			cleanedOrders = append(cleanedOrders, model.CleanedOrder{
				No:         nextOrderNo,
				ProductId:  finalProductID,
				MaterialId: materialID,
				ModelId:    modelID,
				Qty:        productQty,
				UnitPrice:  unitPricePerItem,
				TotalPrice: totalPricePerItem * float64(productQty),
			})
			nextOrderNo++

			materialParts := strings.Split(materialID, constant.Separator)
			filmType := materialParts[0]
			texture := materialParts[1]
			filmTypePartQuantities[filmType] += productQty
			texturePartQuantities[texture] += productQty
		}
	}

	for filmType, quantity := range filmTypePartQuantities {
		cleanedOrders = append(cleanedOrders, model.CleanedOrder{
			No:         nextOrderNo,
			ProductId:  directStringMapping(filmType),
			Qty:        quantity,
			UnitPrice:  0.00,
			TotalPrice: 0.00,
		})
		nextOrderNo++
	}

	for texture, quantity := range texturePartQuantities {
		cleanedOrders = append(cleanedOrders, model.CleanedOrder{
			No:         nextOrderNo,
			ProductId:  strings.ToUpper(texture) + constant.Separator + constant.Cleaner, // ขออนุญาติแก้จาก Cleanner เป็น Cleaner
			Qty:        quantity,
			UnitPrice:  0.00,
			TotalPrice: 0.00,
		})
		nextOrderNo++
	}

	return &model.CleanedOrderResponse{
		CleanedOrders: cleanedOrders,
	}, nil
}

func directStringMapping(input string) string {
	if val, ok := replacementMap[input]; ok {
		return val
	}

	return input
}

func extractMaterialIDAndModelID(productID string) (string, string, error) {
	parts := strings.Split(productID, constant.Separator)
	if len(parts) < 3 || len(parts) > 4 {
		return "", "", errors.New(response.ErrInvalidInput)
	}

	return parts[0] + constant.Separator + parts[1],
		strings.Join(parts[2:], constant.Separator),
		nil
}

func cleanData(platformProductIDs []string) ([]string, []int, int) {
	cleanedPlatformProductIDs := make([]string, 0, len(platformProductIDs))
	quantity := make([]int, 0, len(platformProductIDs))
	totalQuantity := 0
	for _, platformProductID := range platformProductIDs {
		fgIndex := strings.Index(platformProductID, constant.FG)
		if fgIndex != -1 {
			newPlatformProductID := strings.TrimSpace(platformProductID[fgIndex:])
			parts := strings.Split(newPlatformProductID, constant.Multiplication)
			if len(parts) > 2 {
				continue
			}

			cleanedPlatformProductIDs = append(cleanedPlatformProductIDs, parts[0])
			if len(parts) == 1 {
				quantity = append(quantity, 1)
				totalQuantity += 1
				continue
			}

			qty, err := strconv.Atoi(parts[1])
			if err != nil {
				quantity = append(quantity, 1)
				totalQuantity += 1
				continue
			}

			quantity = append(quantity, qty)
			totalQuantity += qty
		}
	}

	return cleanedPlatformProductIDs, quantity, totalQuantity
}
