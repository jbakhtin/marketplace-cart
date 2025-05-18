package domain

type CartID int64
type UserID int64
type CartItem string

type Cart struct {
	ID     CartID
	userID UserID
	items  []CartItem
}
