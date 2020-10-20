package main

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

type movieReview struct {
	Author  string
	Comment string
	Rating  int
	AddedOn string
}

func getMovieReview(c *fiber.Ctx) error {
	movieID := c.Params("id")
	queryResult, err := Db.Query("select author,comment,rate,added_on from review where movie_id = ? order by added_on desc", movieID)
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString("Database error")
	}
	var data []movieReview = []movieReview{}
	for queryResult.Next() {
		var temp movieReview
		err = queryResult.Scan(&temp.Author, &temp.Comment, &temp.Rating, &temp.AddedOn)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString("Server error")
		}
		data = append(data, temp)
	}
	if queryResult.Err() != nil {
		log.Println(err)
		return c.Status(500).SendString("Server error")
	}
	return c.JSON(data)
}

type review struct {
	Rating  int    `json:"Rating"`
	Comment string `json:"Comment"`
}

func addReview(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claim := user.Claims.(jwt.MapClaims)
	email, ok := claim["email"].(string)
	if !ok {
		return c.Status(500).SendString("Server Error")
	}
	movieID := c.Params("id")
	data := new(review)
	if err := c.BodyParser(data); err != nil {
		log.Println(err.Error())
		return c.Status(400).JSON("Invalid request")
	}
	if data.Rating < 0 || data.Rating > 5 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Rating should be around 0 and 5",
		})
	}
	if x := len(data.Comment); x < 1 || x > 199 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Review should be at least 1 character and maximum 199",
		})
	}
	sql, err := Db.Query("SELECT rating,voters FROM movies WHERE movie_id=?", movieID)
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
	_, err = Db.Exec("UPDATE movies SET rating=?,voters=? WHERE movie_id=?", finalRating, voter+1, movieID)
	if err != nil {
		log.Println(err.Error())
		return c.SendStatus(500)
	}
	u1 := uuid.NewV4().String()
	_, err = Db.Exec("INSERT INTO review (review_id,author,comment,rate,movie_id) values (?,?,?,?,?)", u1, email, data.Comment, data.Rating, movieID)
	if err != nil {
		log.Println(err.Error())
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

func getReviewAddedBy(c *fiber.Ctx) error {
	author := c.Params("id")
	queryResult, err := Db.Query("select author,comment,rate,added_on from review where author = ? order by added_on desc", author)
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString("Database error")
	}
	var data []movieReview = []movieReview{}
	for queryResult.Next() {
		var temp movieReview
		err = queryResult.Scan(&temp.Author, &temp.Comment, &temp.Rating, &temp.AddedOn)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString("Server error")
		}
		data = append(data, temp)
	}
	if queryResult.Err() != nil {
		log.Println(err)
		return c.Status(500).SendString("Server error")
	}
	return c.JSON(data)

}
