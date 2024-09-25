package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cart "git.imooc.com/coding-447/cart/proto/cart"
	cartApi "git.imooc.com/coding-447/cartApi/proto/cartApi"
	"github.com/prometheus/common/log"
)

type CartApi struct {
	CartService cart.CartService
}

// CartApi.Call is exposed externally as /cartApi/findAll, receiving HTTP requests
// i.e., /cartApi/call request will call the CartApi.Call method of the go.micro.api.cartApi service
func (e *CartApi) FindAll(ctx context.Context, req *cartApi.Request, rsp *cartApi.Response) error {
	log.Info("Received /cartApi/findAll access request")
	if _, ok := req.Get["user_id"]; !ok {
		//rsp.StatusCode= 500
		return errors.New("parameter error")
	}
	userIdString := req.Get["user_id"].Values[0]
	fmt.Println(userIdString)
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		return err
	}
	// Get all items in the cart
	cartAll, err := e.CartService.GetAll(context.TODO(), &cart.CartFindAll{UserId: userId})
	// Data type conversion
	b, err := json.Marshal(cartAll)
	if err != nil {
		return err
	}
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}
