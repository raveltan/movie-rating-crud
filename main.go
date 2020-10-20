package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
)

// Db is the instance of Mysql database
var Db *sql.DB

type config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	} `yaml:"database"`
}

func main() {
	fmt.Println("---------------------")
	fmt.Println("Movie CRUD Webserver")
	fmt.Println("---------------------")

	// Local server
	databaseUser, databasePassword, databaseName := "john", "John@1234", "movie"
	var err error
	Db, err = sql.Open("mysql", databaseUser+":"+databasePassword+"@/"+databaseName)

	if err != nil {
		panic(err.Error())
	}
	err = Db.Ping()
	if err != nil {
		log.Println("Please make sure sql server is running and mysql setup is correct, refer to README.md for instruction")
		panic(err.Error())
	}
	// Start fiber web server
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(cors.New())

	//Unrestricted route
	app.Post("/api/login", Login)
	app.Post("/api/register", Register)
	app.Static("/", "./public")
	// Others routes

	//Refresh route
	app.Use("/api/refresh", jwtware.New(jwtware.Config{
		SigningKey: []byte("GIAO GIAO"),
	}))

	app.Post("/api/refresh", Refresh)

	//Restricted route
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("一给我里GIAO GIAO"),
	}))

	app.Get("/api/movies", GetMovie)
	app.Get("/api/review/:id", GetReview)
	app.Post("/api/movie/add", AddMovie)
	app.Post("/api/review/:id/add", AddReview)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Redirect("/")
	})

	var port string

	// local
	port = "80"

	fmt.Println("Connect to server at http://localhost:" + port)
	log.Fatalln(app.Listen(":" + port))
}
