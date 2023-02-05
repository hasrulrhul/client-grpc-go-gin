package main

import (
	"log"
	"os"

	"github.com/hasrulrhul/client-grpc-go-gin/model"
	userHandler "github.com/hasrulrhul/client-grpc-go-gin/user/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	port := os.Getenv("APP_PORT")
	targetPort := os.Getenv("APP_PORT_SERVER")
	conn, err := grpc.Dial(":"+targetPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cloud not conect to %v %v", targetPort, err)
	}
	user := model.NewUsersClient(conn)
	router := gin.Default()

	userHandler.CreateUserHandler(router, user)

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
