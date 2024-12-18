package product

import (
	"github.com/Kroning/mytheresa/internal/database/postgresql"
)

type Repo struct {
	db *postgresql.Storage
}

func NewRepo(db *postgresql.Storage) *Repo {
	return &Repo{db: db}
}
