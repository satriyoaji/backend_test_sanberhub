package repository

import (
	"backend_test/entity"
	"context"
)

func (d DefaultRepository) FindMutationsByNumber(ctx context.Context, accountNumber string) ([]entity.Mutation, error) {
	var mutations []entity.Mutation
	err := d.handler.Tx.WithContext(ctx).Where("number=?", accountNumber).Find(&mutations).Error
	return mutations, err
}

func (d DefaultRepository) CreateMutation(ctx context.Context, mutation *entity.Mutation) error {
	return d.handler.Tx.Create(mutation).Error
}

func (d DefaultRepository) UpdateMutation(ctx context.Context, mutation *entity.Mutation) error {
	return d.handler.Tx.WithContext(ctx).Save(mutation).Error
}
