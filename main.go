package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
)

// Db is the instance of Mysql database
var Db *sql.DB

func main() {
	// Initialize database
	// databaseUser, databasePassword, databaseName := "sql12366524", "7fESNz9TQR", "sql12366524"
	// var err error
	// db, err = sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp(sql12.freemysqlhosting.net:3306)/"+databaseName)

	//Local server
	databaseUser, databasePassword, databaseName := "hung", "RavelTan@123", "movie"
	var err error
	Db, err = sql.Open("mysql", databaseUser+":"+databasePassword+"@/"+databaseName)

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

	//Refresh route
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("GIAO GIAO"),
	}))

	app.Get("/api/refresh", Refresh)

	//Restricted route
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("一给我里GIAO GIAO"),
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})

	// app.Get("/logout", processLogout)
	// app.Post("/auth", processAuth)
	// app.Post("/add-movie", processAddMovie)
	// app.Get("/review/:id", reviewPage)
	// app.Post("/add-review", proccessAndReview)
	// app.Get("*", unknownRoute)

	//port := os.Getenv("PORT")

	//local
	port := "3000"
	log.Fatalln(app.Listen(":" + port))
}

// type movieData struct {
// 	Id     int
// 	Name   string
// 	Rating string
// 	Voter  int
// }

// func unknownRoute(c *fiber.Ctx) error {
// 	return c.Redirect("/")
// }
// func homePage(c *fiber.Ctx) error {
// 	if username == nil {
// 		return c.Redirect("/auth")
// 	}
// 	if user, err := username.(string); !err {
// 		if user == "" {
// 			return c.Redirect("/auth")
// 		}
// 	}
// 	var data []movieData
// 	queryResult, err := db.Query("SELECT id,name,rating,voter FROM movie")
// 	if err != nil {
// 		log.Panicln(err.Error())
// 	}
// 	for queryResult.Next() {
// 		var temp movieData
// 		var tempRate float64
// 		err = queryResult.Scan(&temp.Id, &temp.Name, &tempRate, &temp.Voter)
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		temp.Rating = fmt.Sprintf("%.2f", tempRate)
// 		data = append(data, temp)
// 	}
// 	return c.Render("index", fiber.Map{
// 		"username":  username,
// 		"film":      data,
// 		"totalFilm": len(data),
// 	}, "layout/main")
// }

// func authPage(c *fiber.Ctx) error {
// 	store := sessions.Get(c)
// 	username := store.Get("username")
// 	if username != nil {
// 		if user := username.(string); user != "" {
// 			return c.Redirect("/")
// 		}
// 	}
// 	defer store.Save()
// 	errorStatus := store.Get("authError")
// 	var authError string
// 	if errorStatus != nil {
// 		r, ok := errorStatus.(string)
// 		if ok {
// 			authError = r
// 		}
// 	}
// 	store.Delete("authError")
// 	return c.Render("auth", fiber.Map{
// 		"loginFailed": authError,
// 	}, "layout/main")
// }

// func processAddMovie(c *fiber.Ctx) error {
// 	store := sessions.Get(c)
// 	username := store.Get("username")
// 	movie := c.FormValue("name")
// 	if movie != "" && username != nil {
// 		_, err := db.Exec("INSERT INTO movie (name) values (?)", movie)
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 	}
// 	return c.Redirect("/")
// }

// func processAuth(c *fiber.Ctx) error {
// 	store := sessions.Get(c)
// 	defer store.Save()
// 	email := c.FormValue("email")
// 	password := c.FormValue("password")
// 	if c.FormValue("authAction") == "register" {
// 		rows, err := db.Query("SELECT email from users where email = ?", email)
// 		if err != nil {
// 			log.Println(err.Error())
// 			store.Set("authError", "Unknown server error!")
// 			return c.Redirect("/auth")
// 		}
// 		found := false
// 		for rows.Next() {
// 			var emailStored string
// 			var err = rows.Scan(&emailStored)
// 			if err != nil {
// 				log.Println(err.Error())
// 			}
// 			if emailStored == email {
// 				found = true
// 			}
// 		}
// 		if err = rows.Err(); err != nil {
// 			log.Println(err)
// 			return c.Redirect("/auth")
// 		}
// 		if !found {
// 			hashedPassword, err := hashPassword(password)
// 			if err != nil {
// 				log.Println(err.Error())
// 			}
// 			_, sqlError := db.Exec("INSERT INTO users (email,password) values (?,?)", email, hashedPassword)
// 			if sqlError != nil {
// 				store.Set("authError", "Unknown server error")
// 				return c.Redirect("/auth")
// 			}
// 			store.Set("username", email)
// 			return c.Redirect("/")

// 		}
// 		store.Set("authError", "User with this email already exists!")
// 		return c.Redirect("/auth")

// 	}
// 	if c.FormValue("authAction") == "login" {
// 		rows, err := db.Query("SELECT password from users WHERE email = ?", email)
// 		if err != nil {
// 			log.Println(err.Error())
// 			store.Set("authError", "Unknown server error")
// 			return c.Redirect("/auth")
// 		}
// 		found := false
// 		for rows.Next() {
// 			var hash string
// 			err = rows.Scan(&hash)
// 			if err != nil {
// 				log.Println(err.Error())
// 				store.Set("authError", "Unknown server error")
// 				return c.Redirect("/auth")
// 			}
// 			found = compareHashAndPassword(password, hash)
// 		}
// 		if !found {
// 			store.Set("authError", "Invalid email or password!")
// 			return c.Redirect("/auth")
// 		}
// 		if found {
// 			store.Set("username", email)
// 			return c.Redirect("/")
// 		}
// 	}
// 	return c.Redirect("/auth")
// }

// type review struct {
// 	Email  string
// 	Review string
// 	Rating int
// }

// func reviewPage(c *fiber.Ctx) error {
// 	store := sessions.Get(c)
// 	username := store.Get("username")
// 	movieID := c.Params("id")
// 	if username == nil {
// 		return c.Redirect("/")
// 	}
// 	//Get movie data
// 	var rating string
// 	var movieName string
// 	var totalReview int
// 	sqlResult, err := db.Query("SELECT name,rating,voter from movie WHERE id=?", movieID)
// 	if err != nil {
// 		return c.Redirect("/")
// 	}
// 	if t, err := sqlResult.ColumnTypes(); len(t) == 0 || err != nil {
// 		return c.Redirect("/")
// 	}
// 	for sqlResult.Next() {
// 		var tempRate float64
// 		err = sqlResult.Scan(&movieName, &tempRate, &totalReview)
// 		rating = fmt.Sprintf("%.2f", tempRate)
// 		if err != nil {
// 			return c.Redirect("/")
// 		}
// 	}
// 	//Get reviews
// 	var reviews []review
// 	sqlResult, err = db.Query("SELECT email,review,rating FROM review WHERE movie=?", movieID)
// 	if err != nil {
// 		return c.Redirect("/")
// 	}
// 	for sqlResult.Next() {
// 		var temp review
// 		err = sqlResult.Scan(&temp.Email, &temp.Review, &temp.Rating)
// 		if err != nil {
// 			return c.Redirect("/")
// 		}
// 		reviews = append(reviews, temp)
// 	}
// 	return c.Render("review", fiber.Map{
// 		"totalReview": totalReview,
// 		"username":    username,
// 		"rating":      rating,
// 		"id":          movieID,
// 		"movieName":   movieName,
// 		"reviews":     reviews,
// 	}, "layout/main")
// }

// func proccessAndReview(c *fiber.Ctx) error {
// 	store := sessions.Get(c)
// 	username := store.Get("username")
// 	if username == nil {
// 		c.Redirect("/")
// 	}
// 	var prevRating float64
// 	var prevReview int
// 	var newRating int
// 	var err error
// 	reviewText := c.FormValue("name")
// 	id := c.FormValue("movie")
// 	if reviewText == "" {
// 		return c.Redirect("/review/" + id)
// 	}
// 	prevRating, err = strconv.ParseFloat(c.FormValue("rating"), 32)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	prevReview, err = strconv.Atoi(c.FormValue("review"))
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	newRating, err = strconv.Atoi(c.FormValue("nrating"))
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	totalStar := prevRating*float64(prevReview) + float64(newRating)
// 	finalReviewer := prevReview + 1
// 	finalRate := totalStar / float64(finalReviewer)
// 	_, err = db.Exec("UPDATE movie SET rating=?,voter=? WHERE id=?", finalRate, finalReviewer, id)
// 	if err != nil {
// 		log.Println(err.Error())
// 		c.Redirect("/review/" + id)
// 	}
// 	_, err = db.Exec("INSERT INTO review (email,review,rating,movie) values (?,?,?,?)", username, reviewText, newRating, id)
// 	if err != nil {
// 		log.Println(err.Error())
// 		c.Redirect("/review/" + id)
// 	}
// 	return c.Redirect("/review/" + id)
// }

// func processLogout(c *fiber.Ctx) error {
// 	store := sessions.Get(c)
// 	defer store.Save()
// 	store.Delete("username")
// 	store.Destroy()
// 	store.Regenerate()
// 	return c.Redirect("/")
// }
