package repository

import (
	"go-mysql-crud/pkg/models"
)

// QueryRepo explain...
type QueryRepo interface {
	Fetch() []*models.Column
}
