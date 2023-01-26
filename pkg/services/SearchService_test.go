package services

import (
	"context"
	"testing"
	"vtrial/mocks"
	models "vtrial/pkg/Models"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestSearchService_SavePage(t *testing.T) {
	/*gin.SetMode(gin.TestMode)
	mockDB:=mocks.NewIDBClient(t)
	mockService:=NewSearchService(mockDB)
	mockDB.On("InsertOnePage", mock.Anything, mock.Anything).Return(nil)
	input := &models.Page{
		Title:    "Page 10",
		Keywords: []string{"wrd1", "wrd2"},
	}
	router:= gin.Default()
	router.POST("/SavePage", mockService.SavePage(,input))
	*/

	mockService := mocks.NewIDBClient(t)
	mockService.On("InsertOnePage", mock.Anything, mock.Anything).Return(nil)
	searchservice := NewSearchService(mockService)
	input := models.Page{
		Title:    "Page 10",
		Keywords: []string{"wrd1", "wrd2"},
	}
	var ctx context.Context
	err := searchservice.SavePage(ctx, input)
	assert.Equal(t, nil, err)

}

func TestSearchService_ComputeResult(t *testing.T) {

}
