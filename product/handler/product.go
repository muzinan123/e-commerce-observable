package handler

import (
	"context"
	"fmt"

	"git.imooc.com/coding-447/product/common"
	"git.imooc.com/coding-447/product/domain/model"
	"git.imooc.com/coding-447/product/domain/service"
	. "git.imooc.com/coding-447/product/proto/product"
)

type Product struct {
	ProductDataService service.IProductDataService
}

func (h *Product) AddProduct(ctx context.Context, request *ProductInfo, response *ResponseProduct) error {
	productAdd := &model.Product{}
	fmt.Println(request)
	if err := common.SwapTo(request, productAdd); err != nil {
		return err
	}
	fmt.Println(productAdd)
	productID, err := h.ProductDataService.AddProduct(productAdd)
	if err != nil {
		return err
	}
	response.ProductId = productID
	return nil
}

func (h *Product) FindProductByID(ctx context.Context, request *RequestID, response *ProductInfo) error {
	productData, err := h.ProductDataService.FindProductByID(request.ProductId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(productData, response); err != nil {
		return err
	}
	return nil
}

func (h *Product) UpdateProduct(ctx context.Context, request *ProductInfo, response *Response) error {
	productAdd := &model.Product{}
	if err := common.SwapTo(request, productAdd); err != nil {
		return err
	}
	err := h.ProductDataService.UpdateProduct(productAdd)
	if err != nil {
		return err
	}
	response.Msg = "xxxx"
	return nil
}

func (h *Product) DeleteProductByID(ctx context.Context, request *RequestID, response *Response) error {
	if err := h.ProductDataService.DeleteProduct(request.ProductId); err != nil {
		return err
	}
	response.Msg = "xxxx"
	return nil
}

func (h *Product) FindAllProduct(ctx context.Context, request *RequestAll, response *AllProduct) error {
	productAll, err := h.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}

	for _, v := range productAll {
		productInfo := &ProductInfo{}
		err := common.SwapTo(v, productInfo)
		if err != nil {
			return err
		}
		response.ProductInfo = append(response.ProductInfo, productInfo)
	}
	return nil
}
