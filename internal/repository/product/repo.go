package product

import (
	"github.com/Kroning/mytheresa/internal/database/postgresql"
)

type ProductRepo struct {
	db *postgresql.Storage
}

func NewRepo(db *postgresql.Storage) *ProductRepo {
	return &ProductRepo{db: db}
}
