package main

import (
	"fmt"
	"vtrial/cmd/config"
	"vtrial/pkg/Controllers"
	"vtrial/pkg/DatabaseConn"
	"vtrial/pkg/External"
	"vtrial/pkg/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig() // Load Config from Config.Yaml
	StartServer()       //Start Server
}

/*
3 Layers of Abstraction
Controllers access the APIS housed in  Search Service which uses a DB client
*/
func StartServer() {
	router := gin.Default()//Set gin router
	mongoClient := External.NewMongoService(DatabaseConn.Connection()) //Use a Mongoclient constructor which takes in *mongo collection pointer
	searchService:=services.NewSearchService(mongoClient)//Use a SearchService constructor that takes in pointer to mongo collection and returns searchservice pointer that gives us access to two apis
    controller := Controllers.NewController(searchService)//Use a Controller constructor that takes in the searchService
	controller.Routes(router.Group("/" + config.Config.Server.Version))
	router.Run(":" + fmt.Sprint(config.Config.Server.Port))
}
