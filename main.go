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
var sessions *session.Session

type authError struct {
	emailError    bool
	passwordError bool
	loginFailed   bool
}

func homePage(c *fiber.Ctx) error {
	store := sessions.Get(c)
	if username := store.Get("username"); username == nil {
		return c.Redirect("/auth")
	}
	return c.Render("index", fiber.Map{}, "layout/main")
}

func authPage(c *fiber.Ctx) error {
	store := sessions.Get(c)
	if username := store.Get("username"); username != nil {
		return c.Redirect("/")
	}
	defer store.Save()
	errorStatus := store.Get("authError")
	r, ok := errorStatus.(bool)
	if !ok {
		r = false
	}
	store.Destroy()
	fmt.Println(r)
	return c.Render("auth", fiber.Map{
		"loginFailed": r,
	}, "layout/main")
}

func processAuth(c *fiber.Ctx) error {
	store := sessions.Get(c)
	defer store.Save()
	store.Set("authError", true)
	fmt.Println(c.Body())
	return c.Redirect("/auth")
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
	app.Get("/auth", authPage)
	app.Get("/logout", logout)
	app.Post("/auth", processAuth)
	log.Fatalln(app.Listen(":8080"))
}
