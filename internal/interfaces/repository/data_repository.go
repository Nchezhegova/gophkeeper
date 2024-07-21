package repository

import (
	"github.com/Nchezhegova/gophkeeper/internal/entities"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DataRepository struct {
	DB *gorm.DB
}

func (r *DataRepository) StoreData(userID uint32, key, dataType string, data []byte) error {
	return r.DB.Create(&entities.Data{
		UserID: userID,
		Key:    key,
		Type:   dataType,
		Data:   data,
	}).Error
}

func (r *DataRepository) GetData(userID uint32, key, dataType, identifier string) ([]entities.Data, error) {
	var data []entities.Data
	query := r.DB.Where("user_id = ?", userID)
	if key != "" {
		query = query.Where("key = ?", key)
	}
	if dataType != "" {
		query = query.Where("type = ?", dataType)
	}
	if identifier != "" {
		if dataType == "login" {
			query = query.Where("json_extract(data, '$.login') = ?", identifier)
		} else if dataType == "bank card" {
			query = query.Where("json_extract(data, '$.number') = ?", identifier)
		}
	}
	err := query.Find(&data).Error
	if len(data) == 0 {
		return nil, status.Errorf(codes.NotFound, "data not found")
	}
	return data, err
}

func (r *DataRepository) UpdateData(userID uint32, key, dataType, identifier string, newData []byte) error {
	query := r.DB.Model(&entities.Data{}).Where("user_id = ? AND key = ? AND type = ?", userID, key, dataType)
	if dataType == "login" {
		query = query.Where("json_extract(data, '$.login') = ?", identifier)
	} else if dataType == "bank card" {
		query = query.Where("json_extract(data, '$.number') = ?", identifier)
	}
	return query.Update("data", newData).Error
}

func (r *DataRepository) DeleteData(userID uint32, key, dataType, identifier string) error {
	query := r.DB.Where("user_id = ? AND key = ? AND type = ?", userID, key, dataType)
	if dataType == "login" {
		query = query.Where("json_extract(data, '$.login') = ?", identifier)
	} else if dataType == "bank card" {
		query = query.Where("json_extract(data, '$.number') = ?", identifier)
	}
	return query.Delete(&entities.Data{}).Error
}
