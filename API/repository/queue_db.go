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

func (r queueRepositoryDB) GetQueuesByCode(code string) (*model.QueueModel, error) {
	queue := model.QueueModel{}
	err := r.db.Where("Code = ?", code).First(&queue).Error
	if err != nil {
		return nil, err
	}
	return &queue, nil
}

func (r queueRepositoryDB) DeleteQueue(code string) (*model.QueueModel, error) {
	var queue model.QueueModel
	err := r.db.Where("Code = ?", code).First(&queue).Error
	if err != nil {
		return nil, err
	}
	r.db.Where("Code = ?", code).Delete(&queue)
	return &queue, nil
}

func (r queueRepositoryDB) CreateQueue(data model.QueueInput) (*model.QueueModel, error) {
	var input model.QueueInput
	newCode := r.generateCode(input.Type)
	date := time.Now()
	code := fmt.Sprintf("%v%03d", input.Type, newCode)
	Queue := model.QueueModel{
		Code: code,
		Type: input.Type,
		Date: date,
		Name: input.Name,
		Tel:  input.Tel}
	r.db.Create(&Queue)
	return &Queue, nil
}

func (r queueRepositoryDB) generateCode(genre string) (NewCode int) {
	queue := model.QueueModel{}
	r.db.Where("Type=?", genre).Limit(1).Order("Date desc").Find(&queue)
	last := queue.Date.Format("2006-02-01")
	// fmt.Println(last)
	now := time.Now().Format("2006-02-01")
	// fmt.Println(now)

	if last == now {
		strCode := strings.Trim(queue.Code, genre)
		intVar, _ := strconv.Atoi(strCode)
		NewCode := intVar + 1
		fmt.Println(NewCode)
		return NewCode
	}

	NewCode = 1
	fmt.Println(NewCode)
	return NewCode
}
