package services

import (
	"context"
	"sort"
	"strings"
	"vtrial/pkg/External"
	"vtrial/pkg/Models"

	
)

/*Notice how we always declare APIS in services this provides an additional layer of abstraction
controllers themselves don't possess these API but they use services.
But the APIs itself interact with the DB so we need a constructor to that which inturn helps to use the functions we
defined in the interface IDBClient
*/

type SearchService struct{
	DBClient External.IDBClient
}

func NewSearchService(dbParam External.IDBClient) *SearchService{
	return &SearchService{DBClient: dbParam}
}

func (s SearchService) SavePage(ctx context.Context, page models.Page) (error) {
	return s.DBClient.InsertOnePage(ctx, page)
}

func (s SearchService) ComputeResult(ctx context.Context, words models.Keywords) ([]string, error) {
pages := s.DBClient.GetAllCollection()

res := []models.Result{}
for i := 0; i < len(pages); i++ {
	var tempRs models.Result
    tempScore := 0
	for j := 0; j < len(pages[i].Keywords); j++ {
      for k := 0; k < len(words.ArrayOfString); k++ {
		if strings.EqualFold(pages[i].Keywords[j], words.ArrayOfString[k]) {
			tempScore += (10 - k) * (10 - j)
          }
		}
	}
if tempScore > 0 {
	tempRs.Title = pages[i].Title
	tempRs.Score = tempScore
}
	res = append(res, tempRs)
}

sort.Stable(models.PagesByScore(res))
ans := []string{}
for i := 0; i < len(res); i++ {
	ans = append(ans, res[i].Title)
}
return ans, nil
}