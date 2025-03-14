package repository

import (
	"butler/application/domains/services/outbound_order/models"
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

func (r *repo) GetById(ctx context.Context, id int64) (*models.OutboundOrder, error) {
	record := &models.OutboundOrder{}
	result := r.dbWithContext(ctx).Limit(1).Find(&record, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if record.OutboundOrderId == 0 {
		return nil, nil
	}
	return record, nil
}

func (r *repo) GetOne(ctx context.Context, params *models.GetRequest) (*models.OutboundOrder, error) {
	record := &models.OutboundOrder{}
	query := r.dbWithContext(ctx).Model(record)
	query = r.filter(query, params)
	query = r.sort(query, params)
	result := query.Limit(1).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return record, nil
}

func (r *repo) GetList(ctx context.Context, params *models.GetRequest) ([]*models.OutboundOrder, error) {
	records := []*models.OutboundOrder{}
	query := r.dbWithContext(ctx).Model(&models.OutboundOrder{})
	query = r.filter(query, params)
	query = r.sort(query, params)
	if err := query.Scan(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r *repo) Update(ctx context.Context, obj *models.OutboundOrder) (*models.OutboundOrder, error) {
	if obj.OutboundOrderId == 0 {
		return nil, fmt.Errorf("id is required")
	}
	result := r.dbWithContext(ctx).Updates(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) UpdateMany(ctx context.Context, objs []*models.OutboundOrder) error {
	tx := r.dbWithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, obj := range objs {
		if err := tx.Updates(obj).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) filter(query *gorm.DB, params *models.GetRequest) *gorm.DB {
	if params.SalesOrderNumber != "" {
		query = query.Where("sales_order_number = ?", params.SalesOrderNumber)
	}
	if params.WarehouseId != 0 {
		query = query.Where("warehouse_id = ?", params.WarehouseId)
	}
	if params.StatusId != 0 {
		query = query.Where("status_id = ?", params.StatusId)
	}
	if len(params.StatusIds) > 0 {
		query = query.Where("status_id in ?", params.StatusIds)
	}
	return query
}

func (r *repo) sort(query *gorm.DB, params *models.GetRequest) *gorm.DB {
	if params.OrderType != "asc" {
		params.OrderType = "desc"
	}
	if params.SortField != "" {
		query = query.Order(params.SortField + " " + params.OrderType)
	} else {
		query = query.Order("outbound_order_id desc")
	}
	return query
}

func (r *repo) GetListOutboundItems(ctx context.Context, outboundOrderID int64) ([]*models.OutboundOrderItem, error) {
	items := []*models.OutboundOrderItem{}
	if outboundOrderID == 0 {
		return nil, nil
	}

	result := r.dbWithContext(ctx).Where("outbound_order_id = ?", outboundOrderID).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}
