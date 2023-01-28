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
	mockService := mocks.NewIDBClient(t)
	mockService.On("InsertOnePage", mock.Anything, mock.Anything).Return(nil)
	searchservice := NewSearchService(mockService)

	input := models.Page{
		Title:    "Page 10",
		Keywords: []string{"wrd1", "wrd2"},
	}

	err := searchservice.SavePage(context.TODO(), input)

	assert.Equal(t, nil, err)
}

func TestSearchService_GetResult(t *testing.T) {
	out := []models.Page{
		{
			Title:    "Page 10",
			Keywords: []string{"wrd1", "wrd2"},
		},
	}

	mockDB := mocks.NewIDBClient(t)
	mockDB.On("GetAllCollection").Return(out)
	searchservice := NewSearchService(mockDB)

	res := []string{
		"Page 10",
	}

	input := models.Keywords{
		ArrayOfString: []string{"wrd1", "wrd2"},
	}

	exp, err := searchservice.ComputeResult(context.TODO(), input)
	assert.Equal(t, exp, res)
	assert.Equal(t, nil, err)
}
