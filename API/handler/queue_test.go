package handler_test

import (
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
	t.Run("success", func(t *testing.T) {
		//Arrange
		responseService := []model.QueuesResponse{
			{Code: "A003", Type: "A", Date: time.Date(2020, time.April, 10, 21, 34, 01, 0, time.UTC), Name: "Golf", Tel: "1150"},
			{Code: "A004", Type: "A", Date: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC), Name: "Nop", Tel: "1112"},
			{Code: "B005", Type: "B", Date: time.Date(2022, time.July, 12, 21, 34, 05, 0, time.UTC), Name: "Steven", Tel: "191"},
		}

		// data := fmt.Sprintf(`{"data":%s}`, responseService)
		data, _ := json.Marshal(responseService)
		expected := fmt.Sprintf(`{"data":%s}`, string(data))
		// expected := `{"data:[{"Code": "A003", "Type": "A", "Date": "2020-04-10T21:34:01Z, "Name": "Golf", "Tel": "1150"},{"Code": "A004", "Type": "A", "Date": "2020-04-11T21:34:01Z, "Name": "Nop", "Tel": "1112"},{"Code": "B005", "Type": "B", "Date": "2020-04-12T21:34:05Z, "Name": "Steven", "Tel": "191"},]"}`

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

		// var queues []model.QueueResponse
		// json.Unmarshal(w.Body.Bytes(), &queues)

		//Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expected, string(responseData))

	})

	t.Run("Error", func(t *testing.T) {
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

// func TestGetQueuesType(t *testing.T) {
// 	t.Run("success", func(t *testing.T) {
// 		//Arrange
// 		responseRepo := []model.QueuesResponse{
// 			{Code: "A003", Type: "A", Date: time.Date(2020, time.April, 10, 21, 34, 01, 0, time.UTC), Name: "Golf", Tel: "1150"},
// 			{Code: "A004", Type: "A", Date: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC), Name: "Nop", Tel: "1112"},
// 		}

// 		// expected := `{"data:[{"Code": "A003", "Type": "A", "Date": "2020-04-10T21:34:01Z, "Name": "Golf", "Tel": "1150"},{"Code": "A004", "Type": "A", "Date": "2020-04-11T21:34:01Z, "Name": "Nop", "Tel": "1112"},{"Code": "B005", "Type": "B", "Date": "2020-04-12T21:34:05Z, "Name": "Steven", "Tel": "191"},]"}`
// 		// data := fmt.Sprintf(`{"data":%v}`, expected)
// 		queueService := service.NewQueueServiceMock()
// 		queueService.On("GetQueuesType", "A").Return(responseRepo, nil)

// 		queueHandler := handler.NewQueueHandler(queueService)

// 		//Act
// 		r := gin.Default()
// 		r.GET("/:Type", queueHandler.GetQueuesType)
// 		req, _ := http.NewRequest("GET", "/:Type", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		// responseData, _ := ioutil.ReadAll(w.Body)

// 		var queues []model.QueueResponse
// 		json.Unmarshal(w.Body.Bytes(), &queues)

// 		//Assert
// 		assert.Equal(t, http.StatusOK, w.Code)
// 		// assert.Equal(t, expected, queues)

// 	})

// }
