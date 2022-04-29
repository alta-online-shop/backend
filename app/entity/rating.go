package entity

type Rating struct {
	ID      uint
	Count   uint
	User    *User
	Product *Product
}
