package main

import(
	"github.com/m0thm4n/UploadService/domain/photo"
	"github.com/m0thm4n/UploadService/mongo"
	"github.com/m0thm4n/UploadService/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	conn := os.Getenv("MONGO_ADDRESS")

	mongoClient, err:= mongo.NewClient(conn)
	if err!=nil{
		panic("Connection could not be established")
	}
	repo:= photo.NewRepository(mongoClient)

	engine := gin.New()
	engine.Use(gin.Recovery())

	routes.InitializeRoutes(engine,repo)

	engine.Run(":30")
}