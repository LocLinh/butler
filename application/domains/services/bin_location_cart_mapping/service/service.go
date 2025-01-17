package service

import (
	"butler/application/domains/services/bin_location_cart_mapping/models"
	repo "butler/application/domains/services/bin_location_cart_mapping/repository"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type service struct {
	repo repo.IRepository
}

func InitService(
	repo repo.IRepository,
) IService {
	return &service{
		repo: repo,
	}
}

func (s *service) GetById(ctx context.Context, id int64) (*models.BinLocationCartMapping, error) {
	record, err := s.repo.GetById(ctx, id)
	if err != nil {
		logrus.Errorf("error when get bin location cart mapping by id: %v", err)
		return nil, fmt.Errorf("error when get cart by id: %v", err)
	}
	return record, nil
}

func (s *service) GetOne(ctx context.Context, params *models.GetRequest) (*models.BinLocationCartMapping, error) {
	record, err := s.repo.GetOne(ctx, params)
	if err != nil {
		logrus.Errorf("error when get cart: err: %v by params: %#v", err, params)
		return nil, fmt.Errorf("error when get cart: err: %v by params: %#v", err, params)
	}
	return record, nil
}

func (s *service) GetList(ctx context.Context, params *models.GetRequest) ([]*models.BinLocationCartMapping, error) {
	records, err := s.repo.GetList(ctx, params)
	if err != nil {
		logrus.Errorf("Error when get list cart: %v", err)
		return nil, fmt.Errorf("error when get list cart: %v", err)
	}

	return records, nil
}

func (s *service) Update(ctx context.Context, obj *models.BinLocationCartMapping) (*models.BinLocationCartMapping, error) {
	record, err := s.repo.Update(ctx, obj)
	if err != nil {
		logrus.Errorf("error when update cart: %v", err)
		return nil, fmt.Errorf("error when update cart: %v", err)
	}

	return record, nil
}
func (s *service) Create(ctx context.Context, obj *models.BinLocationCartMapping) (*models.BinLocationCartMapping, error) {
	record, err := s.repo.Create(ctx, obj)
	if err != nil {
		logrus.Errorf("error when create bin location: %v", err)
		return nil, fmt.Errorf("error when create bin location: %v", err)
	}

	return record, nil
}

func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		logrus.Errorf("error when delete bin location: %v", err)
		return fmt.Errorf("error when delete bin location: %v", err)
	}
	return nil
}
