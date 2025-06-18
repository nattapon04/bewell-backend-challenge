package usecase

import (
	"bewell-backend-challenge/internal/model"
	"bewell-backend-challenge/util/response"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectStringMapping(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Test Case 1: Empty string input
		{
			name:     "Should handle empty string input",
			input:    "",
			expected: "",
		},
		// Test Case 2: Input does NOT exist and is a simple word
		{
			name:     "Should return original input for 'randomword'",
			input:    "randomword",
			expected: "randomword",
		},
		// Test Case 3: Input exists in the replacementMap
		{
			name:     "Should return mapped value for 'FG0A'",
			input:    "FG0A",
			expected: "WIPING-CLOTH",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := directStringMapping(tt.input)
			if got != tt.expected {
				assert.Equal(t, tt.expected, got, "directStringMapping(%q)", tt.input)
			}
		})
	}
}

func TestExtractMaterialIDAndModelID(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantMatID   string
		wantModelID string
		wantErr     error
	}{
		// Test Case 1: Input contains too few parts
		{
			name:        "Should return an invalid input error when there are too few input parts",
			input:       "",
			wantMatID:   "",
			wantModelID: "",
			wantErr:     errors.New(response.ErrInvalidInput),
		},
		// Test Case 2: Input contains too many parts
		{
			name:        "Should return an invalid input error when there are too many input parts",
			input:       "A-B-C-D-E",
			wantMatID:   "",
			wantModelID: "",
			wantErr:     errors.New(response.ErrInvalidInput),
		},
		// Test Case 3: Input contains 3 parts
		{
			name:        "Should return n MaterialID and a ModelID when input contains 3 parts",
			input:       "FG0A-CLEAR-OPPOA3",
			wantMatID:   "FG0A-CLEAR",
			wantModelID: "OPPOA3",
			wantErr:     nil,
		},
		// Test Case 4: Input contains 4 parts
		{
			name:        "Should return a MaterialID and a ModelID when input contains 4 parts",
			input:       "FG0A-CLEAR-OPPOA3-B",
			wantMatID:   "FG0A-CLEAR",
			wantModelID: "OPPOA3-B",
			wantErr:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matID, modelID, err := extractMaterialIDAndModelID(tt.input)
			assert.Equal(t, tt.wantMatID, matID, "materialID mismatch")
			assert.Equal(t, tt.wantModelID, modelID, "modelID mismatch")
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}

func TestCleanData(t *testing.T) {
	tests := []struct {
		name                   string
		input                  []string
		wantCleanedPlatformIDs []string
		wantQuantities         []int
		wantTotalQuantity      int
	}{
		{
			name:                   "Empty input returns empty slices and zero total",
			input:                  []string{},
			wantCleanedPlatformIDs: []string{},
			wantQuantities:         []int{},
			wantTotalQuantity:      0,
		},
		{
			name:                   "Single valid FG with no quantity",
			input:                  []string{"  FG0A-CLEAR-OPPOA3 "},
			wantCleanedPlatformIDs: []string{"FG0A-CLEAR-OPPOA3"},
			wantQuantities:         []int{1},
			wantTotalQuantity:      1,
		},
		{
			name:                   "Single valid FG with quantity",
			input:                  []string{"FG0A-CLEAR-OPPOA3*3"},
			wantCleanedPlatformIDs: []string{"FG0A-CLEAR-OPPOA3"},
			wantQuantities:         []int{3},
			wantTotalQuantity:      3,
		},
		{
			name:                   "Multiple valid FG with mixed quantities",
			input:                  []string{"FG0A-CLEAR-OPPOA3", "FG0A-CLEAR-OPPOA4*5", "--FG0A-CLEAR-OPPOA6*2", "FG0A-CLEAR-OPPOA7*abc"},
			wantCleanedPlatformIDs: []string{"FG0A-CLEAR-OPPOA3", "FG0A-CLEAR-OPPOA4", "FG0A-CLEAR-OPPOA6", "FG0A-CLEAR-OPPOA7"},
			wantQuantities:         []int{1, 5, 2, 1},
			wantTotalQuantity:      9,
		},
		{
			name:                   "Ignore inputs without FG",
			input:                  []string{"ABC123", "NOFHERE", "12345"},
			wantCleanedPlatformIDs: []string{},
			wantQuantities:         []int{},
			wantTotalQuantity:      0,
		},
		{
			name:                   "Ignore inputs with more than two parts separated by multiplication",
			input:                  []string{"FG123*2*3", "FG0A-CLEAR-OPPOA4*1"},
			wantCleanedPlatformIDs: []string{"FG0A-CLEAR-OPPOA4"},
			wantQuantities:         []int{1},
			wantTotalQuantity:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIDs, gotQtys, gotTotal := cleanData(tt.input)
			assert.Equal(t, tt.wantCleanedPlatformIDs, gotIDs, "cleanedPlatformProductIDs mismatch")
			assert.Equal(t, tt.wantQuantities, gotQtys, "quantities mismatch")
			assert.Equal(t, tt.wantTotalQuantity, gotTotal, "totalQuantity mismatch")
		})
	}
}

func TestUsecase_New(t *testing.T) {
	t.Run("should return struct clean order usecase not nil when call new usecase", func(t *testing.T) {
		usecase := &cleanOrderUsecase{}

		assert.NotNil(t, usecase)
	})
}

func TestCleanOrders(t *testing.T) {
	usecase := &cleanOrderUsecase{}

	t.Run("should return empty array response when call clean orders with platform product id contains too few parts", func(t *testing.T) {
		inputOrders := model.OrderRequest{
			Orders: []model.Order{
				{
					PlatformProductId: "AAAA",
					UnitPrice:         10,
					TotalPrice:        10,
					Qty:               1,
				},
			},
		}

		resp, err := usecase.CleanOrders(inputOrders)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Empty(t, resp.CleanedOrders)
	})

	t.Run("should return empty array response when call clean orders with film type is not start with fg", func(t *testing.T) {
		inputOrders := model.OrderRequest{
			Orders: []model.Order{
				{
					PlatformProductId: "AAAA-CLEAR-OPPOA3",
					UnitPrice:         10,
					TotalPrice:        10,
					Qty:               1,
				},
			},
		}

		resp, err := usecase.CleanOrders(inputOrders)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Empty(t, resp.CleanedOrders)
	})

	t.Run("should return empty array response when call clean orders with film type is not start with fg", func(t *testing.T) {
		inputOrders := model.OrderRequest{
			Orders: []model.Order{
				{
					PlatformProductId: "AAAA-CLEAR-OPPOA3",
					UnitPrice:         10,
					TotalPrice:        10,
					Qty:               1,
				},
			},
		}

		resp, err := usecase.CleanOrders(inputOrders)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Empty(t, resp.CleanedOrders)
	})

	t.Run("should return struct cleaned order response when call clean orders with only one product", func(t *testing.T) {
		inputOrders := model.OrderRequest{
			Orders: []model.Order{
				{
					PlatformProductId: "FG0A-CLEAR-IPHONE16PROMAX",
					UnitPrice:         50,
					TotalPrice:        100,
					Qty:               2,
				},
			},
		}

		mockResponse := model.CleanedOrderResponse{
			CleanedOrders: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					No:         2,
					ProductId:  "WIPING-CLOTH",
					Qty:        2,
					UnitPrice:  0.00,
					TotalPrice: 0.00,
				},
				{
					No:         3,
					ProductId:  "CLEAR-CLEANER",
					Qty:        2,
					UnitPrice:  0.00,
					TotalPrice: 0.00,
				},
			},
		}

		resp, err := usecase.CleanOrders(inputOrders)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.CleanedOrders)
		assert.Equal(t, len(resp.CleanedOrders), 3)
		assert.Equal(t, resp.CleanedOrders, mockResponse.CleanedOrders)
	})

	t.Run("should return struct cleaned order response when call clean orders with one product that has wrong prefix", func(t *testing.T) {
		inputOrders := model.OrderRequest{
			Orders: []model.Order{
				{
					PlatformProductId: "x2-3&FG0A-CLEAR-IPHONE16PROMAX",
					UnitPrice:         50,
					TotalPrice:        100,
					Qty:               2,
				},
			},
		}

		mockResponse := model.CleanedOrderResponse{
			CleanedOrders: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					No:         2,
					ProductId:  "WIPING-CLOTH",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         3,
					ProductId:  "CLEAR-CLEANER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		}

		resp, err := usecase.CleanOrders(inputOrders)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.CleanedOrders)
		assert.Equal(t, len(resp.CleanedOrders), 3)
		assert.Equal(t, resp.CleanedOrders, mockResponse.CleanedOrders)
	})

	t.Run("should return struct cleaned order response when call clean orders with one product that has wrong prefix and multiplication symbol", func(t *testing.T) {
		inputOrders := model.OrderRequest{
			Orders: []model.Order{
				{
					PlatformProductId: "x2-3&FG0A-MATTE-IPHONE16PROMAX*3",
					UnitPrice:         90,
					TotalPrice:        90,
					Qty:               1,
				},
			},
		}

		mockResponse := model.CleanedOrderResponse{
			CleanedOrders: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-MATTE-IPHONE16PROMAX",
					MaterialId: "FG0A-MATTE",
					ModelId:    "IPHONE16PROMAX",
					Qty:        3,
					UnitPrice:  30.00,
					TotalPrice: 90.00,
				},
				{
					No:         2,
					ProductId:  "WIPING-CLOTH",
					Qty:        3,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         3,
					ProductId:  "MATTE-CLEANER",
					Qty:        3,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		}

		resp, err := usecase.CleanOrders(inputOrders)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.CleanedOrders)
		assert.Equal(t, len(resp.CleanedOrders), 3)
		assert.Equal(t, resp.CleanedOrders, mockResponse.CleanedOrders)
	})

	t.Run("should return struct cleaned order response when call clean orders with one bundle product that has wrong prefix and split by forward slash symbol", func(t *testing.T) {
		inputOrders := model.OrderRequest{
			Orders: []model.Order{
				{
					PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B",
					UnitPrice:         80,
					TotalPrice:        80,
					Qty:               1,
				},
			},
		}

		mockResponse := model.CleanedOrderResponse{
			CleanedOrders: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					Qty:        1,
					UnitPrice:  40.00,
					TotalPrice: 40.00,
				},
				{
					No:         2,
					ProductId:  "FG0A-CLEAR-OPPOA3-B",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3-B",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:         3,
					ProductId:  "WIPING-CLOTH",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         4,
					ProductId:  "CLEAR-CLEANER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		}

		resp, err := usecase.CleanOrders(inputOrders)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.CleanedOrders)
		assert.Equal(t, len(resp.CleanedOrders), 4)
		assert.Equal(t, resp.CleanedOrders, mockResponse.CleanedOrders)
	})

	t.Run("should return struct cleaned order response when call clean orders with one bundle product that has wrong prefix with split by forward slash symbol and multiplication symbol", func(t *testing.T) {
		inputOrders := model.OrderRequest{
			Orders: []model.Order{
				{
					PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3",
					UnitPrice:         120,
					TotalPrice:        120,
					Qty:               1,
				},
			},
		}

		mockResponse := model.CleanedOrderResponse{
			CleanedOrders: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					No:         2,
					ProductId:  "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId:    "OPPOA3",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:         3,
					ProductId:  "WIPING-CLOTH",
					Qty:        3,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         4,
					ProductId:  "CLEAR-CLEANER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         5,
					ProductId:  "MATTE-CLEANER",
					Qty:        1,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		}

		resp, err := usecase.CleanOrders(inputOrders)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.CleanedOrders)
		assert.Equal(t, len(resp.CleanedOrders), 5)
		assert.Equal(t, resp.CleanedOrders, mockResponse.CleanedOrders)
	})

	t.Run("should return struct cleaned order response when call clean orders with one bundle product and one product that has wrong prefix with split by forward slash symbol and multiplication symbol", func(t *testing.T) {
		inputOrders := model.OrderRequest{
			Orders: []model.Order{
				{
					PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2",
					UnitPrice:         160,
					TotalPrice:        160,
					Qty:               1,
				},
				{
					PlatformProductId: "FG0A-PRIVACY-IPHONE16PROMAX",
					UnitPrice:         50,
					TotalPrice:        50,
					Qty:               1,
				},
			},
		}

		mockResponse := model.CleanedOrderResponse{
			CleanedOrders: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					No:         2,
					ProductId:  "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId:    "OPPOA3",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					No:         3,
					ProductId:  "FG0A-PRIVACY-IPHONE16PROMAX",
					MaterialId: "FG0A-PRIVACY",
					ModelId:    "IPHONE16PROMAX",
					Qty:        1,
					UnitPrice:  50,
					TotalPrice: 50,
				},
				{
					No:         4,
					ProductId:  "WIPING-CLOTH",
					Qty:        5,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         5,
					ProductId:  "CLEAR-CLEANER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         6,
					ProductId:  "MATTE-CLEANER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         7,
					ProductId:  "PRIVACY-CLEANER",
					Qty:        1,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		}

		resp, err := usecase.CleanOrders(inputOrders)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.CleanedOrders)
		assert.Equal(t, len(resp.CleanedOrders), 7)
		assert.Equal(t, resp.CleanedOrders, mockResponse.CleanedOrders)
	})
}
