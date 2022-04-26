package entity

type Order struct {
	ID       uint
	User     *User
	Product  *Product
	Quantity uint
}
