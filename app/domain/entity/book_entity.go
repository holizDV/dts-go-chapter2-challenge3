package entity

// table of books
type Book struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null;type:varchar(200)"`
	Author      string `gorm:"not null;type:varchar(200)"`
	Description string `gorm:"not null;type:varchar(300)"`
}

type BookRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=200"`
	Author      string `json:"author" binding:"required,min=3,max=200"`
	Description string `json:"description" binding:"required,min=5,max=300"`
}

type BookResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type BookDeleteResponse struct {
	ID uint `json:"id"`
}
