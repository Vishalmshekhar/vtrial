package External

import (
	"context"
	models "vtrial/pkg/Models"


)

type IDBClient interface {
	GetAllCollection() []models.Page //User defined function that returns bson.M array of Collections
	InsertOnePage(context.Context, models.Page) error //Our database uses inbuilt InsertOne which takes a context and the page and returns an error
}