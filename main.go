package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"github.com/gofiber/template/handlebars"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var sessions *session.Session

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func compareHashAndPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type movieData struct {
	Name   string
	Rating float32
	Voter  int
}

func unknownRoute(c *fiber.Ctx) error {
	return c.Redirect("/")
}
func homePage(c *fiber.Ctx) error {
	store := sessions.Get(c)
	username := store.Get("username")
	if username == nil {
		return c.Redirect("/auth")
	}
	if user, err := username.(string); !err {
		if user == "" {
			return c.Redirect("/auth")
		}
	}
	var data []movieData
	queryResult, err := db.Query("SELECT name,rating,voter FROM movie")
	if err != nil {
		log.Panicln(err.Error())
	}
	for queryResult.Next() {
		var temp movieData
		err = queryResult.Scan(&temp.Name, &temp.Rating, &temp.Voter)
		if err != nil {
			log.Println(err.Error())
		}
		data = append(data, temp)
	}
	return c.Render("index", fiber.Map{
		"username":  username,
		"film":      data,
		"totalFilm": len(data),
	}, "layout/main")
}

func authPage(c *fiber.Ctx) error {
	store := sessions.Get(c)
	username := store.Get("username")
	if username != nil {
		if user := username.(string); user != "" {
			return c.Redirect("/")
		}
	}
	defer store.Save()
	errorStatus := store.Get("authError")
	var authError string
	if errorStatus != nil {
		r, ok := errorStatus.(string)
		if ok {
			authError = r
		}
	}
	store.Delete("authError")
	return c.Render("auth", fiber.Map{
		"loginFailed": authError,
	}, "layout/main")
}

func processAuth(c *fiber.Ctx) error {
	store := sessions.Get(c)
	defer store.Save()
	email := c.FormValue("email")
	password := c.FormValue("password")
	if c.FormValue("authAction") == "register" {
		rows, err := db.Query("SELECT email from users where email = ?", email)
		if err != nil {
			log.Println(err.Error())
			store.Set("authError", "Unknown server error!")
			return c.Redirect("/auth")
		}
		found := false
		for rows.Next() {
			var emailStored string
			var err = rows.Scan(&emailStored)
			if err != nil {
				log.Println(err.Error())
			}
			if emailStored == email {
				found = true
			}
		}
		if err = rows.Err(); err != nil {
			log.Println(err)
			return c.Redirect("/auth")
		}
		if !found {
			hashedPassword, err := hashPassword(password)
			if err != nil {
				log.Println(err.Error())
			}
			_, sqlError := db.Exec("INSERT INTO users (email,password) values (?,?)", email, hashedPassword)
			if sqlError != nil {
				store.Set("authError", "Unknown server error")
				return c.Redirect("/auth")
			}
			store.Set("username", email)
			return c.Redirect("/")

		}
		store.Set("authError", "User with this email already exists!")
		return c.Redirect("/auth")

	}
	if c.FormValue("authAction") == "login" {
		rows, err := db.Query("SELECT password from users WHERE email = ?", email)
		if err != nil {
			log.Println(err.Error())
			store.Set("authError", "Unknown server error")
			return c.Redirect("/auth")
		}
		found := false
		for rows.Next() {
			var hash string
			err = rows.Scan(&hash)
			if err != nil {
				log.Println(err.Error())
				store.Set("authError", "Unknown server error")
				return c.Redirect("/auth")
			}
			found = compareHashAndPassword(password, hash)
		}
		if !found {
			store.Set("authError", "Invalid email or password!")
			return c.Redirect("/auth")
		}
		if found {
			store.Set("username", email)
			return c.Redirect("/")
		}
	}
	return c.Redirect("/auth")
}

func logout(c *fiber.Ctx) error {
	store := sessions.Get(c)
	defer store.Save()
	store.Delete("username")
	store.Destroy()
	store.Regenerate()
	return c.Redirect("/")
}

func main() {
	// Initialize database
	databaseUser, databasePassword, databaseName := "hung", "RavelTan@123", "movie"
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
	app.Get("*", unknownRoute)
	log.Fatalln(app.Listen(":8080"))
}
