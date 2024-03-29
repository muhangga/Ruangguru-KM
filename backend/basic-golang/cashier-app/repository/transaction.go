package repository

type TransactionRepository struct {
	cartItemRepository CartItemRepository
}

func NewTransactionRepository(cartItemRepository CartItemRepository) TransactionRepository {
	return TransactionRepository{cartItemRepository}
}

func (u *TransactionRepository) Pay(amount int) (int, error) {

	// cartItems, err := u.cartItemRepository.SelectAll()

	// if err != nil {
	// 	return 0, err
	// }

	// for i := 0; i < len(cartItems); i++ {
	// 	if cartItems[i].Quantity > 0 {
	// 		amount = amount - (cartItems[i].Price * cartItems[i].Quantity)
	// 	}
	// }

	totalPrice, err := u.cartItemRepository.TotalPrice()
	if err != nil {
		return 0, nil
	}

	moneyChanges := amount - totalPrice
	return moneyChanges, nil
}
