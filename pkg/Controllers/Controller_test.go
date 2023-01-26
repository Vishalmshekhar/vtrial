package Controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"testing"

	"vtrial/mocks"
	"vtrial/pkg/Models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*func TestController_Routes(t *testing.T) {

}*/

func TestController_healthCheck(t *testing.T) {
	mockService := mocks.NewISearchService(t)
	controller := NewController(mockService)
	mockResponse := "\"Alive\""
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/", controller.healthCheck)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestController_UseSavePage(t *testing.T) {
	mockService := mocks.NewISearchService(t)
	mockService.On("SavePage", mock.Anything, mock.Anything).Return(nil)
	controller := NewController(mockService)
	router := gin.Default()
	router.POST("/UseSavePage", controller.UseSavePage)
	input := &models.Page{
		Title:    "Page 10",
		Keywords: []string{"wrd1", "wrd2"},
	}
	jsonInput, _ := json.Marshal(input)
	req := httptest.NewRequest("POST", "/UseSavePage", bytes.NewBuffer(jsonInput))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code)
}
func TestController_UseSavepageError1(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockSearchService := mocks.NewISearchService(t)
	mockController := NewController(mockSearchService)
	test := `"dkf ldjf ldlfjldj ljf"`
	router.POST("/UseSavePage", mockController.UseSavePage)
	req := httptest.NewRequest("POST", "/UseSavePage", bytes.NewBuffer([]byte(test)))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.NotEmpty(t, resp.Body)
}
func TestController_UseSavepageError2(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockSearchService := mocks.NewISearchService(t)
	mockController := NewController(mockSearchService)
	testPage := models.Page{
		Title:    "test",
		Keywords: []string{"hello", "world"},
	}
	jsonBytes, _ := json.Marshal(testPage)
	mockSearchService.On("SavePage", mock.Anything, mock.Anything).Return(errors.New("something wrong"))
	router.POST("/UseSavePage", mockController.UseSavePage)
	req := httptest.NewRequest("POST", "/UseSavePage", bytes.NewBuffer([]byte(jsonBytes)))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.NotEmpty(t, resp.Body)

}

func TestController_UseComputeResult(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSearchService := mocks.NewISearchService(t)
	mockController := NewController(mockSearchService)
	router := gin.Default()
	
	input := &models.Keywords{
		ArrayOfString: []string{"wrd1", "wrd2"},
	}
	mockSearchService.On("ComputeResult", mock.Anything, mock.Anything).Return([]string{}, nil)
	router.GET("/ComputeResult", mockController.UseComputeResult)
	jsonInput, _ := json.Marshal(input)
	req := httptest.NewRequest("GET", "/ComputeResult", bytes.NewBuffer(jsonInput))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestController_UseComputeResultError1(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockSearchService := mocks.NewISearchService(t)
	mockController := NewController(mockSearchService)
	test := `"dkf ldjf ldlfjldj ljf"`
	jsonBytes, _ := json.Marshal(test)
	router.GET("/compute", mockController.UseComputeResult)
	req := httptest.NewRequest("GET", "/compute", bytes.NewBuffer([]byte(jsonBytes)))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.NotEmpty(t, resp.Body)
}

func TestController_UseComputeResultError2(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockSearchService := mocks.NewISearchService(t)
	mockController := NewController(mockSearchService)
	testPage := models.Keywords{ArrayOfString: []string{"x", "y"}}
	jsonBytes, _ := json.Marshal(testPage)
	mockSearchService.On("ComputeResult", mock.Anything, mock.Anything).Return([]string{}, errors.New("error"))
	router.GET("/compute", mockController.UseComputeResult)
	req := httptest.NewRequest("GET", "/compute", bytes.NewBuffer([]byte(jsonBytes)))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.NotEmpty(t, resp.Body)
}
