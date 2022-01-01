package main

import (
	"food-app-go/infrastructure/persistence"
	"food-app-go/interfaces"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	_ "log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	services, err := persistence.NewRepositories(dbdriver, user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	defer services.Close()

	services.Automigrate()

	users := interfaces.NewUsers(services.User)

	r := gin.Default()
	/*r.Use(middleware.CORSMiddleware()) //For CORS*/

	//user routes
	r.POST("/users", users.SaveUser)
	r.GET("/users/:user_id", users.GetUser)

	//Starting the application
	app_port := os.Getenv("PORT") //using heroku host
	if app_port == "" {
		app_port = "8888" //localhost
	}
	log.Fatal(r.Run(":" + app_port))

}
