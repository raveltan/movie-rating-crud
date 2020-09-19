package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"github.com/gofiber/template/handlebars"
)

var db *sql.DB

func main() {
	// Initialize database
	databaseUser, databasePassword, databaseName := "hung", "RavelTan@123", "blog"
	var err error
	db, err = sql.Open("mysql", databaseUser+":"+databasePassword+"@/"+databaseName)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	// Start fiber web server
	sessions := session.New()
	engine := handlebars.New("./views", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	fmt.Println(sessions)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello world",
		}, "layout/main")
	})
	log.Fatalln(app.Listen(":8080"))
}
