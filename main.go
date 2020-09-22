package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
)

// Db is the instance of Mysql database
var Db *sql.DB

func main() {
	//Initialize database
	databaseUser, databasePassword, databaseName := "sql12366524", "7fESNz9TQR", "sql12366524"
	var err error
	Db, err = sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp(sql12.freemysqlhosting.net:3306)/"+databaseName)

	//Local server
	// databaseUser, databasePassword, databaseName := "hung", "RavelTan@123", "movie"
	// var err error
	// Db, err = sql.Open("mysql", databaseUser+":"+databasePassword+"@/"+databaseName)

	if err != nil {
		panic(err.Error())
	}
	err = Db.Ping()
	if err != nil {
		panic(err.Error())
	}
	// Start fiber web server
	app := fiber.New()
	app.Use(cors.New())

	//Unrestricted route
	app.Post("/api/login", Login)
	app.Post("/api/register", Register)
	app.Static("/", "./frontend/dist")
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
	app.Get("/api/movie/:id", getMovieData)
	app.Get("/api/review/:id", GetReview)
	app.Post("/api/movie/add", AddMovie)
	app.Post("/api/review/:id/add", AddReview)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Redirect("/")
	})
	port := os.Getenv("PORT")

	// local
	// port := "3000"
	log.Fatalln(app.Listen(":" + port))
}
