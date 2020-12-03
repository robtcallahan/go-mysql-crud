package mysql

import (
	"go-mysql-crud/pkg/models"
	repo "go-mysql-crud/pkg/repository"

	"gorm.io/gorm"
)

type mysqlQueryRepo struct {
	Conn *gorm.DB
}

// NewSQLQueryRepo retunrs implement of post repository interface
func NewSQLQueryRepo(Conn *gorm.DB) repo.QueryRepo {
	return &mysqlQueryRepo{
		Conn: Conn,
	}
}

func (m *mysqlQueryRepo) Fetch() []*models.Column {
	var cols []*models.Column

	m.Conn.Order("column_index").Find(&cols)
	return cols
}
