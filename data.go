package main

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type movie struct {
	ID     int
	Name   string
	Rating float64
}

//GetMovie get all movie from database
func GetMovie(c *fiber.Ctx) error {
	queryResult, err := Db.Query("SELECT id,name,rating FROM movie")
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString("Server error")
	}
	var movieList []movie
	for queryResult.Next() {
		var temp movie
		err = queryResult.Scan(&temp.ID, &temp.Name, &temp.Rating)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString("Server error")
		}
		movieList = append(movieList, temp)
	}
	if queryResult.Err() != nil {
		log.Println(err)
		return c.Status(500).SendString("Server error")
	}
	return c.JSON(movieList)
}

type movieReview struct {
	Name   string
	Review string
	Rating int
}

// GetReview get reviews for specific movie
func GetReview(c *fiber.Ctx) error {
	// FIXME: SWITCH TO JOIN SYNTAX AND DECOUPLE DATA BEFORE RETURNING REQUEST
	// with data as (SELECT movie.name as movie,movie.rating,review.name as review_name,review.review,review.rating as movie_rating,movie.id as movie_id FROM review INNER JOIN movie ON review.movie = movie.id) SELECT * FROM data WHERE movie_id=7;
	movieID := c.Params("id")
	queryResult, err := Db.Query("SELECT movie.name as movie,movie.rating,review.name as review_name,review.review,review.rating as movie_rating,movie.id as movie_id FROM review INNER JOIN movie ON review.movie = movie.id AND movie.id = ?", movieID)
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString("Server error")
	}
	var reviewList []movieReview
	var movieName string
	var movieRating float32
	var movieIDNumber int
	for queryResult.Next() {
		var temp movieReview
		err = queryResult.Scan(&movieName, &movieRating, &temp.Name, &temp.Review, &temp.Rating, &movieIDNumber)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString("Server error")
		}
		reviewList = append(reviewList, temp)
	}
	if queryResult.Err() != nil {
		log.Println(err)
		return c.Status(500).SendString("Server error")
	}
	if movieName == "" {
		return c.SendStatus(404)
	}
	return c.JSON(fiber.Map{
		"id":     movieIDNumber,
		"name":   movieName,
		"rating": movieRating,
		"review": reviewList,
	})
}

type newMovie struct {
	Name string `json:"name"`
}

//AddMovie add movie to the database
func AddMovie(c *fiber.Ctx) error {

	data := new(newMovie)
	if err := c.BodyParser(data); err != nil {
		log.Println(err.Error())
		return c.SendStatus(500)
	}
	if len(data.Name) < 3 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Movie name should be at least 3 characters",
		})
	}
	_, err := Db.Exec("INSERT INTO movie (name) values (?)", data.Name)
	if err != nil {
		log.Panicln(err.Error())
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

type review struct {
	Rating int    `json:"rating"`
	Review string `json:"review"`
}

//AddReview add a new review to a specific movie
func AddReview(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claim := user.Claims.(jwt.MapClaims)
	firstName, ok := claim["firstName"].(string)
	if !ok {
		return c.Status(500).SendString("Server Error")
	}
	lastName, ok := claim["lastName"].(string)
	if !ok {
		return c.Status(500).SendString("Server Error")
	}
	movieID := c.Params("id")
	data := new(review)
	if err := c.BodyParser(data); err != nil {
		log.Println(err.Error())
		return c.Status(400).JSON(err.Error())
	}
	if data.Rating < 0 || data.Rating > 5 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Rating should be around 0 and 5",
		})
	}
	if x := len(data.Review); x < 1 || x > 199 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Review should be at least 1 character and maximum 199",
		})
	}
	sql, err := Db.Query("SELECT rating,voter FROM movie WHERE id=?", movieID)
	if err != nil {
		log.Panicln(err.Error())
		return c.SendStatus(500)
	}
	if !sql.Next() {
		return c.Status(404).JSON(fiber.Map{
			"error": "movie not found!",
		})
	}
	var rating float64
	var voter int
	err = sql.Scan(&rating, &voter)
	finalRating := ((rating * float64(voter)) + float64(data.Rating)) / float64(voter+1)
	_, err = Db.Exec("UPDATE movie SET rating=?,voter=? WHERE id=?", finalRating, voter+1, movieID)
	if err != nil {
		log.Println(err.Error())
		return c.SendStatus(500)
	}
	_, err = Db.Exec("INSERT INTO review (name,review,rating,movie) values (?,?,?,?)", firstName+" "+lastName, data.Review, data.Rating, movieID)
	if err != nil {
		log.Println(err.Error())
		return c.SendStatus(500)
	}
	return c.JSON(fiber.Map{
		"success": data,
	})
}

type movieData struct {
	Name   string
	Rating float64
}
