package services

import (
	"context"
	 "vtrial/pkg/Models"
)

//Create actual API definitions that are written by us
type ISearchService interface{
   SavePage(context.Context,models.Page)(error)
   ComputeResult(context.Context,models.Keywords)([]string,error)
}