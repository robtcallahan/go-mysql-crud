package handler

import (
	"go-mysql-crud/pkg/driver"
	"go-mysql-crud/pkg/models"
	"go-mysql-crud/pkg/repository"
	query "go-mysql-crud/pkg/repository/mysql"
)

// Query ...
type Query struct {
	repo repository.QueryRepo
}

// NewQueryHandler ...
func NewQueryHandler(db *driver.DB) *Query {
	return &Query{
		repo: query.NewSQLQueryRepo(db.SQL),
	}
}

// Fetch all post data
func (q *Query) Fetch() []*models.Column {
	return q.repo.Fetch()
}
