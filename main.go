package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"github.com/gofiber/template/handlebars"
)

var db *sql.DB
var sessions *session.Session

func homePage(c *fiber.Ctx) error {
	store := sessions.Get(c)
	if username := store.Get("username"); username == nil {
		return c.Redirect("/login")
	}
	return c.SendString("hello")
}

func loginPage(c *fiber.Ctx) error {
	store := sessions.Get(c)
	defer store.Save()
	store.Set("username", "ravel")
	return c.Render("index", fiber.Map{}, "layout/main")
}

func logout(c *fiber.Ctx) error {
	store := sessions.Get(c)
	defer store.Save()
	store.Destroy()
	return c.Redirect("/")
}

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
	sessions = session.New()
	engine := handlebars.New("./views", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", homePage)
	app.Get("/login", loginPage)
	app.Get("/logout", logout)
	log.Fatalln(app.Listen(":8080"))
}
