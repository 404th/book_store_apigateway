package models

type Book struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	About string `json:"about" db:"about"`
	ISBN  string `json:"isbn" db:"isbn"`
}

type CreateBook struct {
	Name  string `json:"name" db:"name"`
	About string `json:"about" db:"about"`
	ISBN  string `json:"isbn" db:"isbn"`
}
