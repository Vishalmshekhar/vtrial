package External

import (
	"context"
	"vtrial/pkg/Models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoService struct {
	coll *mongo.Collection // a struct for mongo service contains a col which is a pointer of mongo collection
}

//now create a constructor to use it

func NewMongoService(coll*mongo.Collection) *MongoService{
	return &MongoService{coll: coll}
}

func (c *MongoService) GetAllCollection() []bson.M {
cur, err := c.coll.Find(context.Background(), bson.D{})
if err != nil {
     panic(err)
}
var temp []bson.M
for cur.Next(context.Background()) {
var p bson.M
err := cur.Decode(&p)
if err != nil {
panic(err)
}
temp = append(temp, p)
}
defer cur.Close(context.Background())
return temp
}

func (c *MongoService) InsertOnePage(ctx context.Context, newPage models.Page) error {
	_, insertErr := c.coll.InsertOne(ctx, newPage)//inbuilt insertone takes in context and the new page
	return insertErr
}