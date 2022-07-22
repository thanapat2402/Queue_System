package repository

import (
	"errors"
	"log"
	"q/model"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

//Adapter private
type queueRepositoryDB struct {
	db *gorm.DB
}

//Constructor Public เพื่อ new instance
func NewQueueRepositoryDB(db *gorm.DB) QueueRepository {
	return queueRepositoryDB{db: db}
}

//buld all receiver function for interface
func (r queueRepositoryDB) GetAllQueues() ([]model.QueueModel, error) {
	queues := []model.QueueModel{}
	err := r.db.Order("Date").Find(&queues).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return queues, nil
}

func (r queueRepositoryDB) GetQueuesByType(types string) ([]model.QueueModel, error) {
	queues := []model.QueueModel{}
	err := r.db.Where("Type = ?", types).Order("Date").Find(&queues).Error
	if err != nil {
		return nil, err
	}
	return queues, nil
}

func (r queueRepositoryDB) GetQueuesByNameTypes(name string, types string) (*model.QueueModel, error) {
	queue := model.QueueModel{}
	err := r.db.Where("Name = ? AND Type = ?", name, types).First(&queue).Error
	if err != nil {
		return nil, err
	}
	return &queue, nil
}

func (r queueRepositoryDB) GetQueuesByCode(strcode string) (*model.QueueModel, error) {
	queue := model.QueueModel{}
	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	Type := strings.Trim(strcode, num)
	result := r.db.Where("Code = ? AND Type = ?", code, Type).Find(&queue)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user Code not found")
	}
	return &queue, nil
}

func (r queueRepositoryDB) DeleteQueue(strcode string) (*model.QueueModel, error) {
	var queue model.QueueModel
	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	types := strings.Trim(strcode, num)
	result := r.db.Where("Code = ? AND Type = ?", code, types).Find(&queue)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user Code not found")
	}
	r.db.Where("Code = ? AND Type = ?", code, types).Delete(&queue)
	return &queue, nil
}

func (r queueRepositoryDB) CreateQueue(data model.QueueInput) (*model.QueueModel, error) {
	queue := model.QueueModel{}
	result := r.db.Where("user_id = ?", data.UserID).Find(&queue)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		newCode := r.generateCode(data.Type)
		date := time.Now()
		Queue := model.QueueModel{
			Code:   newCode,
			Type:   data.Type,
			Date:   date,
			Name:   data.Name,
			Tel:    data.Tel,
			UserID: data.UserID}
		r.db.Create(&Queue)
		return &Queue, nil
	}
	return nil, errors.New("queue already exists")
}

func (r queueRepositoryDB) generateCode(genre string) (NewCode int) {
	queue := model.QueueModel{}
	r.db.Where("Type=?", genre).Limit(1).Order("Date desc").Find(&queue)
	last := queue.Date.Format("2006-02-01")
	now := time.Now().Format("2006-02-01")
	if last == now {
		NewCode := queue.Code + 1
		return NewCode
	}
	NewCode = 1
	return NewCode
}
