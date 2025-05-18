package postgres

type CartStorage struct {
}

func NewOrderStorage() (CartStorage, error) {
	return CartStorage{}, nil
}
