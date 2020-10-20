package main

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

// Unexport all functions

type movie struct {
	ID      string
	Name    string
	Rating  float64
	AddedBy string
	Voters  int
	AddedOn string
}

func getMovieList(c *fiber.Ctx) error {
	queryResult, err := Db.Query("SELECT movie_id,name,added_on,added_by,rating,voters FROM movies order by added_on desc")
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString("Server error")
	}
	var movieList = []movie{}
	for queryResult.Next() {
		var temp movie
		err = queryResult.Scan(&temp.ID, &temp.Name, &temp.AddedOn, &temp.AddedBy, &temp.Rating, &temp.Voters)
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

func getMovie(c *fiber.Ctx) error {
	movieID := c.Params("id")
	queryResult, err := Db.Query("SELECT movie_id,name,added_on,added_by,rating,voters FROM movies where movie_id = ?", movieID)
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString("Server error")
	}
	var temp movie
	found := false
	for queryResult.Next() {
		err = queryResult.Scan(&temp.ID, &temp.Name, &temp.AddedOn, &temp.AddedBy, &temp.Rating, &temp.Voters)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString("Server error")
		}
		found = true
	}
	if queryResult.Err() != nil {
		log.Println(err)
		return c.Status(500).SendString("Server error")
	}
	if !found {
		return c.Status(404).JSON(
			fiber.Map{
				"error": "Movie not found",
			},
		)
	}
	return c.JSON(temp)
}

func getMovieAddedBy(c *fiber.Ctx) error {
	userID := c.Params("id")
	queryResult, err := Db.Query("SELECT movie_id,name,added_on,added_by,rating,voters FROM movies where added_by = ? order by added_on desc", userID)
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString("Server error")
	}
	var movieList = []movie{}
	for queryResult.Next() {
		var temp movie
		err = queryResult.Scan(&temp.ID, &temp.Name, &temp.AddedOn, &temp.AddedBy, &temp.Rating, &temp.Voters)
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

type newMovie struct {
	Name string `json:"name"`
}

func addMovie(c *fiber.Ctx) error {
	data := new(newMovie)
	if err := c.BodyParser(data); err != nil {
		log.Println(err.Error())
		return c.SendStatus(500)
	}
	if len(data.Name) < 2 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Movie name should be at least 2 characters",
		})
	}
	user := c.Locals("user").(*jwt.Token)
	claim := user.Claims.(jwt.MapClaims)
	email, ok := claim["email"].(string)
	if !ok {
		return c.Status(500).SendString("Server Error")
	}
	u1 := uuid.NewV4().String()
	_, err := Db.Exec("INSERT INTO movies (movie_id,name,added_by) values (?,?,?)", u1, data.Name, email)
	if err != nil {
		log.Panicln(err.Error())
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
