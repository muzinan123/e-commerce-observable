package handler

import (
	"context"

	"git.imooc.com/coding-447/cart/domain/model"
	"git.imooc.com/coding-447/cart/domain/service"
	cart "git.imooc.com/coding-447/cart/proto/cart"
	"git.imooc.com/coding-447/common"
)

type Cart struct {
	CartDataService service.ICartDataService
}

// Add item to cart
func (h *Cart) AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) (err error) {
	cart := &model.Cart{}
	common.SwapTo(request, cart)
	response.CartId, err = h.CartDataService.AddCart(cart)
	return err
}

// Clear cart
func (h *Cart) CleanCart(ctx context.Context, request *cart.Clean, response *cart.Response) error {
	if err := h.CartDataService.CleanCart(request.UserId); err != nil {
		return err
	}
	response.Meg = "Cart cleared successfully"
	return nil
}

// Increase item quantity in cart
func (h *Cart) Incr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := h.CartDataService.IncrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Meg = "Item quantity increased successfully"
	return nil
}

// Decrease item quantity in cart
func (h *Cart) Decr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := h.CartDataService.DecrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Meg = "Item quantity decreased successfully"
	return nil
}

// Delete item from cart
func (h *Cart) DeleteItemByID(ctx context.Context, request *cart.CartID, response *cart.Response) error {
	if err := h.CartDataService.DeleteCart(request.Id); err != nil {
		return err
	}
	response.Meg = "Item deleted from cart successfully"
	return nil
}

// Get all cart items for a user
func (h *Cart) GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) error {
	cartAll, err := h.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}

	for _, v := range cartAll {
		cart := &cart.CartInfo{}
		if err := common.SwapTo(v, cart); err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, cart)
	}
	return nil
}
