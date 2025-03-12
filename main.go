package main

import (
	"example/controllers"
	"example/dao"
	"fmt"

	"gofr.dev/pkg/gofr/datasource/mongo"

	"gofr.dev/pkg/gofr"
)

func main() {
	// initialise gofr object
	app := gofr.New()

	db := mongo.New(mongo.Config{URI: "mongodb://root:pass@localhost:27017/", Database: "test"})
	userDAO := dao.NewUserDAO(db)
	userController := controllers.NewUserDAO(userDAO)
	// inject the mongo into gofr to use mongoDB across the application
	// using gofr context
	app.AddMongo(db)
	fmt.Println("Running on port : 9000")
	// register route greet
	app.GET("/greet", func(ctx *gofr.Context) (any, error) {
		return "Hello World!", nil
	})

	app.POST("/sign-up", userController.SignUp())

	// Runs the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Run()
}
