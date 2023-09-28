package repository

import (
	"backend_test/entity"
	"backend_test/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func userTableName() string {
	return "users"
}

func whereUserNameContains(name string, alias string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name == "" {
			return db
		}
		sql := withAlias("name", alias) + " ilike ?"
		return db.Where(sql, withPercentAround(name))
	}
}

func whereUserIDIn(ids []int, alias string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(ids) == 0 {
			return db
		}
		sql := withAlias("id", alias) + " in ?"
		return db.Where(sql, ids)
	}
}

func (d DefaultRepository) FindUsers(ctx context.Context, filter model.GetUsersFilter) ([]entity.User, error) {
	userIDs := []int{}
	if filter.ID != nil {
		userIDs = append(userIDs, *filter.ID)
	}
	shops := []entity.User{}
	err := d.handler.Tx.WithContext(ctx).
		Scopes(
			whereUserNameContains(filter.Name, ""),
			whereUserIDIn(userIDs, ""),
			paginate(filter.PageRequest.PageNum, filter.PageRequest.PageSize)).
		Order("created_at desc").Find(&shops).Error
	return shops, err
}

func (d DefaultRepository) CountUsers(ctx context.Context, filter model.GetUsersFilter) (int, error) {
	userIDs := []int{}
	if filter.ID != nil {
		userIDs = append(userIDs, *filter.ID)
	}
	var count int64
	err := d.handler.Tx.WithContext(ctx).Model(&entity.User{}).
		Scopes(
			whereUserNameContains(filter.Name, ""),
			whereUserIDIn(userIDs, "")).
		Count(&count).Error
	return int(count), err
}

func (d DefaultRepository) CreateUser(ctx context.Context, user *entity.User) error {
	return d.handler.Tx.WithContext(ctx).Create(user).Error
}

func (d DefaultRepository) FindUserByID(ctx context.Context, id uint) (entity.User, error) {
	user := entity.User{}
	err := d.handler.Tx.WithContext(ctx).Where("id=?", id).First(&user).Error
	return user, err
}

func (d DefaultRepository) FindUserByColumnValue(ctx context.Context, columnName string, search interface{}) (entity.User, error) {
	user := entity.User{}
	err := d.handler.Tx.WithContext(ctx).Where(fmt.Sprintf("%s=?", columnName), search).First(&user).Error
	return user, err
}

func (d DefaultRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return d.handler.Tx.WithContext(ctx).Save(user).Error
}
