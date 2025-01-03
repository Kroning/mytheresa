package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Kroning/mytheresa/internal/domain"
	"github.com/Kroning/mytheresa/internal/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ProductService interface {
	GetProducts(ctx context.Context, category string, price int) ([]*domain.Product, error)
}

type DiscountService interface {
	GetDiscounts(ctx context.Context) (domain.Discounts, error)
}

const (
	paramCategory = "category"
	paramPrice    = "priceLessThan"

	ErrNoCategory   = "no category in request"
	ErrInvalidPrice = "invalid price in request"
	ErrGetProducts  = "cannot get products"
	ErrGetDiscounts = "cannot get discounts"
)

type ApiHandler struct {
	productService  ProductService
	discountService DiscountService
}

func NewApiHandler(productService ProductService, discountService DiscountService) *ApiHandler {
	return &ApiHandler{
		productService:  productService,
		discountService: discountService,
	}
}

func (h *ApiHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	category := r.URL.Query().Get(paramCategory)
	if category == "" {
		logger.Error(ctx, ErrNoCategory)
		ErrorJSON(w, r, http.StatusBadRequest, errors.New(ErrNoCategory))
		return
	}
	priceStr := r.URL.Query().Get(paramPrice)
	var price int
	var err error
	if priceStr != "" {
		price, err = strconv.Atoi(priceStr)
		if err != nil {
			logger.Error(ctx, ErrInvalidPrice, zap.Error(err))
			ErrorJSON(w, r, http.StatusBadRequest, errors.New(ErrInvalidPrice))
			return
		}
	}

	products, err := h.productService.GetProducts(ctx, category, price)
	if err != nil {
		logger.Error(ctx, ErrGetProducts, zap.Error(err))
		ErrorJSON(w, r, http.StatusInternalServerError, errors.New(ErrGetProducts))
		return
	}

	discounts, err := h.discountService.GetDiscounts(ctx)
	// TODO: does not produce any error right now. Add test when it will be
	if err != nil {
		logger.Error(ctx, ErrGetDiscounts, zap.Error(err))
		ErrorJSON(w, r, http.StatusInternalServerError, errors.New(ErrGetDiscounts))
		return
	}

	productsWithDiscounts := domain.AddDiscountsToProduct(products, discounts)

	ResponseJSON(w, r, MapProductsWithDiscountResponse(productsWithDiscounts))
}
