package entity

type Rating struct {
	ID        uint
	Count     uint
	UserID    uint `json:"UserID,omitempty"`
	ProductID uint `json:"ProductID,omitempty"`

	User    *User    `json:"User,omitempty"`
	Product *Product `json:"Product,omitempty"`
}
