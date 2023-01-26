package DatabaseConn

import (
	"context"
	"fmt"
	"vtrial/cmd/config"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() (*mongo.Collection) {
	URI := config.Config.Database.Protocol + "://" + config.Config.Database.Host + ":" + fmt.Sprint(config.Config.Database.Port)//eg mongodb://mongo-container:27017
	clientOptions := options.Client().ApplyURI(URI)//set new client instance and apply the URI we generated through config
	client, err := mongo.Connect(context.TODO(), clientOptions)//Initialize client using client options context.Todo() used when we dont know what context to define
	if err != nil {
		panic(err)
	}
	return client.Database(config.Config.Database.DBName).Collection(config.Config.Database.Collection)//return pointer of this DB and this collections
}