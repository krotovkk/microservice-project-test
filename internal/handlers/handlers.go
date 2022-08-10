package handlers

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
)

var errBadArgument = errors.New("Bad argument")

type BotHandler struct {
	productService ports.ProductService
}

func NewBotHandler(service ports.ProductService) *BotHandler {
	return &BotHandler{
		productService: service,
	}
}

func (bh *BotHandler) ListProducts(args string) string {
	var res string

	products, _ := bh.productService.GetAllProducts(context.Background(), 0, 0)

	if len(products) == 0 {
		return fmt.Sprintf("No products available")
	}

	for _, product := range products {
		res += product.String()
	}
	return res
}

func (bh *BotHandler) AddProduct(args string) string {
	params := strings.Split(args, " ")
	if len(params) != 2 {
		return errors.Wrapf(errBadArgument, "params: <%v>", params).Error()
	}
	name := params[0]
	price, err := strconv.ParseFloat(params[1], 64)
	if err != nil {
		return err.Error()
	}

	err = bh.productService.CreateProduct(context.Background(), name, price)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("product <%v> added successfuly", name)
}

func (bh *BotHandler) DeleteProduct(args string) string {
	params := strings.Split(args, " ")
	if len(params) != 1 {
		return errors.Wrapf(errBadArgument, "params: <%v>", params).Error()
	}
	id, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil {
		return err.Error()
	}
	err = bh.productService.DeleteProduct(context.Background(), uint(id))
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("product with id <%d> deleted successfuly", id)
}

func (bh *BotHandler) UpdateProduct(args string) string {
	params := strings.Split(args, " ")
	if len(params) != 3 {
		return errors.Wrapf(errBadArgument, "params: <%v>", params).Error()
	}

	id, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil {
		return err.Error()
	}
	name := params[1]
	price, err := strconv.ParseFloat(params[2], 64)
	if err != nil {
		return err.Error()
	}

	err = bh.productService.UpdateProduct(context.Background(), name, price, uint(id))
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("product with id <%d> updated successfuly", id)
}

func (bh *BotHandler) Help(args string) string {
	return "/help - list commands\n" +
		"/list - list data\n" +
		"/add <name> <price> - add new product with name and price\n" +
		"/delete <id> - add new product with name and price\n" +
		"/update <id> <name> <price> - add new product with name and price"

}
