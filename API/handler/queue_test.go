package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"q/handler"
	"q/model"
	"q/service"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetQueues(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	t.Run("success", func(t *testing.T) {
		//Arrange
		responseService := []model.QueuesResponse{
			{Code: "A003", Type: "A", Date: time.Date(2020, time.April, 10, 21, 34, 01, 0, time.UTC), Name: "Golf", Tel: "1150"},
			{Code: "A004", Type: "A", Date: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC), Name: "Nop", Tel: "1112"},
			{Code: "B005", Type: "B", Date: time.Date(2022, time.July, 12, 21, 34, 05, 0, time.UTC), Name: "Steven", Tel: "191"},
		}

		data, _ := json.Marshal(responseService)
		expected := fmt.Sprintf(`{"data":%s}`, string(data))
		queueService := service.NewQueueServiceMock()
		queueService.On("GetQueues").Return(responseService, nil)
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.GET("/", queueHandler.GetQueues)
		req, _ := http.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := ioutil.ReadAll(w.Body)

		//Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expected, string(responseData))

	})

	t.Run("Error Service", func(t *testing.T) {
		//Arrange
		responseService := []model.QueuesResponse{}
		queueService := service.NewQueueServiceMock()
		queueService.On("GetQueues").Return(responseService, errors.New("Error something"))
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.GET("/", queueHandler.GetQueues)
		req, _ := http.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

}

func TestGetQueuesType(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	t.Run("success", func(t *testing.T) {
		//Arrange
		types := "A"
		responseService := []model.QueuesResponse{
			{Code: "A003", Type: "A", Date: time.Date(2020, time.April, 10, 21, 34, 01, 0, time.UTC), Name: "Golf", Tel: "1150"},
			{Code: "A004", Type: "A", Date: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC), Name: "Nop", Tel: "1112"},
		}
		queueService := service.NewQueueServiceMock()
		queueService.On("GetQueuesType", types).Return(responseService, nil)

		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.GET("/:Type", queueHandler.GetQueuesType)
		req, _ := http.NewRequest("GET", "/"+types, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Error Service", func(t *testing.T) {
		//Arrange
		types := "A"
		responseService := []model.QueuesResponse{}
		queueService := service.NewQueueServiceMock()
		queueService.On("GetQueuesType", types).Return(responseService, errors.New("Error something"))
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.GET("/:Type", queueHandler.GetQueuesType)
		req, _ := http.NewRequest("GET", "/"+types, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("Invalid type", func(t *testing.T) {
		//Arrange
		types := "F"
		responseService := []model.QueuesResponse{}
		queueService := service.NewQueueServiceMock()
		queueService.On("GetQueuesType", types).Return(responseService, nil)
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.GET("/:Type", queueHandler.GetQueuesType)
		req, _ := http.NewRequest("GET", "/"+types, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusNotAcceptable, w.Code)

	})

}

func TestSearchQueue(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	t.Run("success", func(t *testing.T) {
		//Arrange
		name := "Nop"
		types := "A"
		responseService := model.QueueResponse{Code: "A004", Date: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC), Name: "Nop", Tel: "1112"}
		queueService := service.NewQueueServiceMock()
		queueService.On("SearchQueue", name, types).Return(&responseService, nil)

		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.GET("/search", queueHandler.SearchQueue)
		url := fmt.Sprintf("/search?name=%v&types=%v", name, types)
		req, _ := http.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Error Service", func(t *testing.T) {
		//Arrange
		name := "Nop"
		types := "A"
		responseRepo := model.QueueResponse{}
		queueService := service.NewQueueServiceMock()
		queueService.On("SearchQueue", name, types).Return(&responseRepo, errors.New("Error something"))
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.GET("/search", queueHandler.SearchQueue)
		url := fmt.Sprintf("/search?name=%v&types=%v", name, types)
		req, _ := http.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("Invalid type", func(t *testing.T) {
		//Arrange
		name := "Nop"
		types := "s"
		responseRepo := model.QueueResponse{}
		queueService := service.NewQueueServiceMock()
		queueService.On("SearchQueue", name, types).Return(&responseRepo, nil)
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.GET("/search", queueHandler.SearchQueue)
		url := fmt.Sprintf("/search?name=%v&types=%v", name, types)
		req, _ := http.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusNotAcceptable, w.Code)

	})

}

func TestGetQueue(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	t.Run("success", func(t *testing.T) {
		//Arrange
		code := "A004"
		responseService := model.QueueResponse{Code: "A004", Date: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC), Name: "Nop", Tel: "1112"}
		queueService := service.NewQueueServiceMock()
		queueService.On("GetQueue", code).Return(&responseService, nil)
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.GET("/code/:Code", queueHandler.GetQueue)
		req, _ := http.NewRequest("GET", "/code/"+code, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Error Service", func(t *testing.T) {
		//Arrange
		code := "A004"
		responseService := model.QueueResponse{}
		queueService := service.NewQueueServiceMock()
		queueService.On("GetQueue", code).Return(&responseService, errors.New("Error something"))
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.GET("/code/:Code", queueHandler.GetQueue)
		req, _ := http.NewRequest("GET", "/code/"+code, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

}

func TestAddQueue(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	t.Run("Create", func(t *testing.T) {
		//Arrange
		data := model.QueueInput{Type: "A", Name: "Steven", Tel: "0856565565"}
		jsonValue, _ := json.Marshal(data)
		responseService := model.QueueResponse{Code: "A005", Date: time.Now(), Name: "Steven", Tel: "0856565565"}
		queueService := service.NewQueueServiceMock()
		queueService.On("AddQueue", data).Return(&responseService, nil)
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.POST("/", queueHandler.AddQueue)
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusCreated, w.Code)

	})

	t.Run("Error Service", func(t *testing.T) {
		//Arrange
		data := model.QueueInput{Type: "A", Name: "Steven", Tel: "0856565565"}
		jsonValue, _ := json.Marshal(data)
		responseService := model.QueueResponse{}
		queueService := service.NewQueueServiceMock()
		queueService.On("AddQueue", data).Return(&responseService, errors.New("Error something"))
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.POST("/", queueHandler.AddQueue)
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("Error Conflict", func(t *testing.T) {
		//Arrange
		data := `{Type: "A", Name: "Steven", //Tel: "0856565565"}`
		jsonValue, _ := json.Marshal(data)
		responseService := model.QueueResponse{Code: "A005", Date: time.Now(), Name: "Steven", Tel: "0856565565"}
		queueService := service.NewQueueServiceMock()
		queueService.On("AddQueue", jsonValue).Return(&responseService, nil)
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.POST("/", queueHandler.AddQueue)
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusConflict, w.Code)

	})

}

func TestDeQueue(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	t.Run("Delete", func(t *testing.T) {
		//Arrange
		code := "A004"
		responseService := model.QueueResponse{Code: "A004", Date: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC), Name: "Nop", Tel: "1112"}
		queueService := service.NewQueueServiceMock()
		queueService.On("DeQueue", code).Return(&responseService, nil)
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.DELETE("/:Code", queueHandler.DeQueue)
		req, _ := http.NewRequest("DELETE", "/"+code, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Error Service", func(t *testing.T) {
		//Arrange
		code := "A004"
		responseService := model.QueueResponse{}
		queueService := service.NewQueueServiceMock()
		queueService.On("DeQueue", code).Return(&responseService, errors.New("Error something"))
		queueHandler := handler.NewQueueHandler(queueService)

		//Act
		r := gin.Default()
		r.DELETE("/:Code", queueHandler.DeQueue)
		req, _ := http.NewRequest("DELETE", "/"+code, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

}
