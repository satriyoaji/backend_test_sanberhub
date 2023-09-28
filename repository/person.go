package repository

import (
	"backend_test/entity"
	"backend_test/model"
	"context"

	"gorm.io/gorm"
)

func personTableName() string {
	return "persons"
}

func wherePersonNameContains(name string, alias string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name == "" {
			return db
		}
		sql := withAlias("name", alias) + " ilike ?"
		return db.Where(sql, withPercentAround(name))
	}
}

func wherePersonIsActiveEquals(isActive *bool, alias string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if isActive == nil {
			return db
		}
		sql := withAlias("is_active", alias) + " = ?"
		return db.Where(sql, isActive)
	}
}

func wherePersonIDIn(ids []int, alias string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(ids) == 0 {
			return db
		}
		sql := withAlias("id", alias) + " in ?"
		return db.Where(sql, ids)
	}
}

func (d DefaultRepository) FindPersons(ctx context.Context, filter model.GetPersonsFilter) ([]entity.Person, error) {
	personIDs := []int{}
	if filter.ID != nil {
		personIDs = append(personIDs, *filter.ID)
	}
	shops := []entity.Person{}
	err := d.handler.Tx.WithContext(ctx).
		Scopes(
			wherePersonNameContains(filter.Name, ""),
			wherePersonIDIn(personIDs, ""),
			paginate(filter.PageRequest.PageNum, filter.PageRequest.PageSize)).
		Order("created_at desc").Find(&shops).Error
	return shops, err
}

func (d DefaultRepository) CountPersons(ctx context.Context, filter model.GetPersonsFilter) (int, error) {
	personIDs := []int{}
	if filter.ID != nil {
		personIDs = append(personIDs, *filter.ID)
	}
	var count int64
	err := d.handler.Tx.WithContext(ctx).Model(&entity.Person{}).
		Scopes(
			wherePersonNameContains(filter.Name, ""),
			wherePersonIsActiveEquals(filter.IsActive, ""),
			wherePersonIDIn(personIDs, "")).
		Count(&count).Error
	return int(count), err
}

func (d DefaultRepository) CreatePerson(ctx context.Context, person *entity.Person) error {
	return d.handler.Tx.Create(person).Error
}

func (d DefaultRepository) FindPersonByID(ctx context.Context, id uint) (entity.Person, error) {
	person := entity.Person{}
	err := d.handler.Tx.WithContext(ctx).Where("id=?", id).First(&person).Error
	return person, err
}

func (d DefaultRepository) FindPersonByName(ctx context.Context, name string) (entity.Person, error) {
	person := entity.Person{}
	err := d.handler.Tx.WithContext(ctx).Where("name=?", name).First(&person).Error
	return person, err
}

func (d DefaultRepository) UpdatePerson(ctx context.Context, person *entity.Person) error {
	return d.handler.Tx.WithContext(ctx).Save(person).Error
}
