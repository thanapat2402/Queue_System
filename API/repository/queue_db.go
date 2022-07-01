package repository

import (
	"fmt"
	"q/model"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type queueRepositoryDB struct {
	db *gorm.DB
}

func NewQueueRepositoryDB(db *gorm.DB) QueueRepository {
	return queueRepositoryDB{db: db}
}

func (r queueRepositoryDB) GetAllQueues() ([]model.QueueModel, error) {
	queues := []model.QueueModel{}
	err := r.db.Order("Date").Find(&queues).Error
	if err != nil {
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

func (r queueRepositoryDB) GetQueuesByCode(strcode string) (*model.QueueModel, error) {
	queue := model.QueueModel{}
	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	Type := strings.Trim(strcode, num)
	err := r.db.Where("Code = ? AND Type = ?", code, Type).First(&queue).Error
	if err != nil {
		return nil, err
	}
	return &queue, nil
}

func (r queueRepositoryDB) DeleteQueue(strcode string) (*model.QueueModel, error) {
	var queue model.QueueModel
	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	Type := strings.Trim(strcode, num)
	// code, _ := strconv.Atoi(strings.TrimLeft(strcode, "ABCD"))
	err := r.db.Where("Code = ? AND Type = ?", code, Type).First(&queue).Error
	if err != nil {
		return nil, err
	}
	r.db.Where("Code = ?", code).Delete(&queue)
	return &queue, nil
}

func (r queueRepositoryDB) CreateQueue(data model.QueueInput) (*model.QueueModel, error) {
	newCode := r.generateCode(data.Type)
	date := time.Now()
	Queue := model.QueueModel{
		Code: newCode,
		Type: data.Type,
		Date: date,
		Name: data.Name,
		Tel:  data.Tel}
	r.db.Create(&Queue)
	return &Queue, nil
}

func (r queueRepositoryDB) generateCode(genre string) (NewCode int) {
	queue := model.QueueModel{}
	r.db.Where("Type=?", genre).Limit(1).Order("Date desc").Find(&queue)
	last := queue.Date.Format("2006-02-01")
	now := time.Now().Format("2006-02-01")
	if last == now {
		NewCode := queue.Code + 1
		fmt.Println(NewCode)
		return NewCode
	}
	NewCode = 1
	fmt.Println(NewCode)
	return NewCode
}
