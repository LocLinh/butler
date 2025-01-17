// Lưu ý ở tầng service:
// 	+ Hàm GetListPaging bổ sung config (0: get all, 1: chỉ get data, 2: chỉ get count).
// 	+ Tách thành hàm riêng những chức năng có khả năng tái sử dụng.
// Phải kèm theo lỗi error của hệ thống khi trả lỗi.

package service

import (
	"butler/application/domains/services/user/models"
	repo "butler/application/domains/services/user/repository"
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
func (s *service) GetById(ctx context.Context, id int64) (*models.User, error) {
	record, err := s.repo.GetById(ctx, id)
	if err != nil {
		logrus.Errorf("error when get user by id: %v", err)
		return nil, fmt.Errorf("error when get user by id: %v", err)
	}
	return record, nil
}

func (s *service) GetOne(ctx context.Context, params *models.GetRequest) (*models.User, error) {
	record, err := s.repo.GetOne(ctx, params)
	if err != nil {
		logrus.Errorf("error when get picking item: err: %v by params: %#v", err, params)
		return nil, fmt.Errorf("error when get picking item: err: %v by params: %#v", err, params)
	}
	return record, nil
}

func (s *service) GetList(ctx context.Context, params *models.GetRequest) ([]*models.User, error) {
	records, err := s.repo.GetList(ctx, params)
	if err != nil {
		logrus.Errorf("Error when get list user: %v", err)
		return nil, fmt.Errorf("error when get list picking item: %v", err)
	}

	return records, nil
}

func (s *service) Update(ctx context.Context, obj *models.User) (*models.User, error) {
	record, err := s.repo.Update(ctx, obj)
	if err != nil {
		logrus.Errorf("error when update user: %v", err)
		return nil, fmt.Errorf("error when update picking item: %v", err)
	}

	return record, nil
}
