package CartHandler

import (
	"FinalProjectGO/Models/cart"
	"FinalProjectGO/Models/product"
	"errors"
)

func GetCartList(userId uint) (*[]cart.CartDetails, error) {
	allCartDetails := cart.GetAllCartDetailsOfUser(userId)
	if len(*allCartDetails) == 0 {
		return nil, errors.New("cart is empty")
	}
	return allCartDetails, nil
}

func AddProductToCart(productId uint, amount int, userId uint) (*product.Product, bool, error) {

	chosenProduct := product.SearchById(productId)
	if chosenProduct.ID == 0 {
		return nil, false, errors.New("ProductNotFoundError")
	}

	if chosenProduct.Stock <= amount {
		return nil, false, errors.New("ProductNotEnoughStockError")
	}

	if amount <= 0 {
		return nil, false, errors.New("InvalidNumberOfProductsError")
	}

	if cart.IsProductExist(userId, productId) {
		return nil, false, errors.New("ProductAlreadyExistInCart")
	}

	return chosenProduct, true, nil
}

func UpdateProductInCart(productId uint, userId uint, amount int) (*cart.CartDetails, error) {
	chosenProduct := product.SearchById(productId)
	if chosenProduct.ID == 0 {
		return nil, errors.New("ProductNotFoundError")
	}

	if chosenProduct.Stock <= amount {
		return nil, errors.New("ProductNotEnoughStockError")
	}

	if chosenProduct.Stock <= 0 {
		return nil, errors.New("InvalidNumberOfProductsError")
	}
	if !cart.IsProductExist(userId, productId) {
		return nil, errors.New("ProductNotFoundErrorInCart")
	}

	cartDetails := cart.GetCartDetailsByCartIdAndProductId(userId, productId)
	return cartDetails, nil
}

func DeleteProductFromCart(userId, productId uint) (*cart.CartDetails, error) {
	chosenProduct := product.SearchById(productId)
	if chosenProduct.ID == 0 {
		return nil, errors.New("ProductNotFoundError")
	}
	if !cart.IsProductExist(userId, productId) {
		return nil, errors.New("ProductNotFoundErrorInCart")
	}
	cartDetails := cart.GetCartDetailsByCartIdAndProductId(userId, productId)
	return cartDetails, nil
}
