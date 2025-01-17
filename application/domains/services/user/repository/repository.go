// Kiểm tra validate cho hàm GET ONE, UPDATE ở tầng repository.
// Nên sử dụng subquery trong repository.
// Lưu ý ở tầng repository:
// 	+ Không 'join' vào những bảng khác một cách tùy tiện.
// 	+ Không tự ý dùng 'Group by' trong 'join'.
// 	+ Hạn chế sử dụng search query 'OR' để filter. Nên format lại data để search.

package repository

import (
	"butler/application/domains/services/user/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func InitRepo(db *gorm.DB) IRepository {
	return &repo{
		DB: db,
	}
}
func (r *repo) dbWithContext(ctx context.Context) *gorm.DB {
	return r.DB.WithContext(ctx)
}
func (r *repo) GetById(ctx context.Context, id int64) (*models.User, error) {
	record := &models.User{}
	result := r.dbWithContext(ctx).Limit(1).Find(&record, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if record.Id == 0 {
		return nil, nil
	}
	return record, nil
}

func (r *repo) GetOne(ctx context.Context, params *models.GetRequest) (*models.User, error) {
	record := &models.User{}
	query := r.dbWithContext(ctx).Model(record)
	query = r.filter(query, params)
	result := query.Limit(1).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return record, nil
}

func (r *repo) GetList(ctx context.Context, params *models.GetRequest) ([]*models.User, error) {
	records := []*models.User{}
	query := r.dbWithContext(ctx).Model(&models.User{})
	query = r.filter(query, params)

	if err := query.Scan(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r *repo) Update(ctx context.Context, obj *models.User) (*models.User, error) {
	if obj.Id == 0 {
		return nil, fmt.Errorf("id is required")
	}
	result := r.dbWithContext(ctx).Updates(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) filter(query *gorm.DB, params *models.GetRequest) *gorm.DB {
	if params.UserId != 0 {
		query = query.Where("id = ?", params.UserId)
	}

	if params.Email != "" {
		query = query.Where("email = ?", params.Email)
	}

	return query
}
