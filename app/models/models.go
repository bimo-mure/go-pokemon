package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Book struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Date      string `json:"tanggal"`
}

func (m *DBModel) GetAllBooks() (Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var book Book

	row := m.DB.QueryRowContext(ctx, "SELECT id, name, publisher, date FROM books")
	err := row.Scan(&book.ID, &book.Name, &book.Publisher, &book.Date)
	if err != nil {
		return book, err
	}

	return book, err
}
