package repository

import (
	"errors"
	"q/model"
	// "gorm.io/gorm"
)

func (r queueRepositoryDB) GetCurrentQueue(types string) (*model.QueueModel, error) {
	currentqueue := model.QueueModel{}
	result := r.db.Order("Date").Where("Type = ?", types).Find(&currentqueue)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("current Code not found")
	}
	return &currentqueue, nil
}

func (r queueRepositoryDB) DeleteQueuebyUID(UserID string) (*model.QueueModel, error) {
	queue := model.QueueModel{}
	result := r.db.Order("Date").Where("user_id = ?", UserID).Find(&queue)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user Code not found")
	}
	r.db.Where("user_id", UserID).Delete(&queue)
	return &queue, nil
}

func (r queueRepositoryDB) GetQueueByUserID(UserID string) (*model.QueueModel, error) {
	queue := model.QueueModel{}
	result := r.db.Where("user_id = ?", UserID).Find(&queue)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	return &queue, nil
}
