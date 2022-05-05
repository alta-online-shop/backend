package entity

type Comment struct {
	ID      uint
	Content string
	User    *User
	Product *Product
	Comment *Comment
}
