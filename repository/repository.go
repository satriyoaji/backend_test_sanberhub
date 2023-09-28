package repository

import (
	"backend_test/entity"
	"backend_test/model"
	"backend_test/pkg/db"
	"context"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"strings"
)

type Repository interface {
	TxBegin() error
	TxCommit() error
	TxRollback() error

	// Person
	FindPersons(ctx context.Context, filter model.GetPersonsFilter) ([]entity.Person, error)
	CountPersons(ctx context.Context, filter model.GetPersonsFilter) (int, error)
	CreatePerson(ctx context.Context, merchant *entity.Person) error
	FindPersonByID(ctx context.Context, id uint) (entity.Person, error)
	FindPersonByName(ctx context.Context, name string) (entity.Person, error)
	UpdatePerson(ctx context.Context, merchant *entity.Person) error

	// User
	FindUsers(ctx context.Context, filter model.GetUsersFilter) ([]entity.User, error)
	CountUsers(ctx context.Context, filter model.GetUsersFilter) (int, error)
	CreateUser(ctx context.Context, merchant *entity.User) error
	FindUserByID(ctx context.Context, id uint) (entity.User, error)
	FindUserByColumnValue(ctx context.Context, columnName string, search interface{}) (entity.User, error)
	UpdateUser(ctx context.Context, merchant *entity.User) error
}

type DefaultRepository struct {
	handler *db.Handler
}

func Default(handler *db.Handler) *DefaultRepository {
	return &DefaultRepository{
		handler: handler,
	}
}

func (d DefaultRepository) TxBegin() error {
	log.Debug("Start db transaction")
	d.handler.Tx = d.handler.DB.Begin()
	return d.handler.Tx.Error
}

func (d DefaultRepository) TxCommit() error {
	log.Debug("Commit db transaction")
	err := d.handler.Tx.Commit().Error
	d.handler.Tx = d.handler.DB
	return err
}

func (d DefaultRepository) TxRollback() error {
	log.Debug("Rollback db transaction")
	err := d.handler.Tx.Rollback().Error
	d.handler.Tx = d.handler.DB
	return err
}

func paginate(pageNum, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum <= 0 {
			pageNum = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func withAlias(column, alias string) string {
	if alias == "" {
		return column
	}
	return alias + "." + column
}

func withPercentAround(val string) string {
	return "%" + val + "%"
}

func withPercentAfter(val string) string {
	return val + "%"
}

func withPercentBefore(val string) string {
	return "%" + val
}

func getSortDir(sortDir string) string {
	if strings.ToLower(sortDir) == "asc" {
		return strings.ToLower(sortDir)
	}
	return "desc"
}
