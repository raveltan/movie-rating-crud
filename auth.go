package main

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type loginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registrationData struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func compareHashAndPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Refresh the current refresh token to get a new access token.
func Refresh(c *fiber.Ctx) error {
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
	email, ok := claim["email"].(string)
	if !ok {
		return c.Status(500).SendString("Server Error")
	}
	// Create refreshToken and token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["firstName"] = firstName
	claims["lastName"] = lastName
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaim := refreshToken.Claims.(jwt.MapClaims)
	refreshClaim["firstName"] = firstName
	refreshClaim["lastName"] = lastName
	refreshClaim["email"] = email
	refreshClaim["exp"] = time.Now().Add(time.Hour * 730 * 12).Unix()

	t, err := token.SignedString([]byte("一给我里GIAO GIAO"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	r, err := refreshToken.SignedString([]byte("GIAO GIAO"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(map[string]string{
		"refresh": r,
		"token":   t,
	})

}

// Login to the application with email and password
func Login(c *fiber.Ctx) error {
	loginCreds := new(loginData)
	if err := c.BodyParser(loginCreds); err != nil {
		log.Println("LOGIN:", err.Error())
		return c.Status(400).JSON(map[string]string{
			"error": "request should be in JSON and should contain email and password",
		})
	}
	if len(loginCreds.Email) == 0 || len(loginCreds.Password) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email should not be empty and password should be at least 8 characters",
		})
	}
	rows, err := Db.Query("SELECT email,password,firstName,lastName from users where email = ?", loginCreds.Email)
	if err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(map[string]string{
			"error": "Internal server error",
		})
	}
	found := false
	var firstName string
	var lastName string
	for rows.Next() {
		var emailStored string
		var hashedPassword string
		var err = rows.Scan(&emailStored, &hashedPassword, &firstName, &lastName)
		if err != nil {
			log.Println(err.Error())
		}
		if emailStored == loginCreds.Email && compareHashAndPassword(loginCreds.Password, hashedPassword) {
			found = true
			break
		}
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(map[string]string{
			"error": "Internal server error",
		})
	}
	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	// Create refreshToken and token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["firstName"] = firstName
	claims["lastName"] = lastName
	claims["email"] = loginCreds.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaim := refreshToken.Claims.(jwt.MapClaims)
	refreshClaim["firstName"] = firstName
	refreshClaim["lastName"] = lastName
	refreshClaim["email"] = loginCreds.Email
	refreshClaim["exp"] = time.Now().Add(time.Hour * 730 * 12).Unix()

	t, err := token.SignedString([]byte("一给我里GIAO GIAO"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	r, err := refreshToken.SignedString([]byte("GIAO GIAO"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(map[string]string{
		"refresh": r,
		"token":   t,
	})

}

// Register to application with email and password
func Register(c *fiber.Ctx) error {
	registerCreds := new(registrationData)
	if err := c.BodyParser(registerCreds); err != nil {
		log.Println("LOGIN:", err.Error())
		return c.Status(400).JSON(map[string]string{
			"error": "request should be in JSON and should contains (email,password,firstName,lastName",
		})
	}
	// Checks if email is empty
	if registerCreds.Email == "" {
		return c.Status(400).JSON(map[string]string{
			"error": "Email should not be empty",
		})
	}
	//check if password > 8 characters
	if len(registerCreds.Password) < 8 {
		return c.Status(400).JSON(map[string]string{
			"error": "Password should be more that 8 characters",
		})
	}
	//Check if lastname and firstname is filled
	if len(registerCreds.FirstName) < 3 || len(registerCreds.LastName) < 3 {
		return c.Status(400).JSON(map[string]string{
			"error": "Firstname and Lastname should be more than 3 characters",
		})
	}
	rows, err := Db.Query("SELECT email from users where email = ?", registerCreds.Email)
	if err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(map[string]string{
			"error": "Internal server error",
		})
	}
	found := false
	for rows.Next() {
		var emailStored string
		var err = rows.Scan(&emailStored)
		if err != nil {
			log.Println(err.Error())
		}
		if emailStored == registerCreds.Email {
			found = true
			break
		}
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(map[string]string{
			"error": "Internal server error",
		})
	}
	if !found {
		hashedPassword, err := hashPassword(registerCreds.Password)
		if err != nil {
			log.Println(err.Error())
			return c.Status(500).JSON(map[string]string{
				"error": "Internal server error",
			})
		}
		_, sqlError := Db.Exec("INSERT INTO users (email,password,firstName,lastName) values (?,?,?,?)", registerCreds.Email, hashedPassword, registerCreds.FirstName, registerCreds.LastName)
		if sqlError != nil {
			log.Println(sqlError.Error())
			return c.Status(500).JSON(map[string]string{
				"error": "Internal server error",
			})
		}
		// Create refreshToken and token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["firstName"] = registerCreds.FirstName
		claims["lastName"] = registerCreds.LastName
		claims["email"] = registerCreds.Email
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		refreshToken := jwt.New(jwt.SigningMethodHS256)
		refreshClaim := refreshToken.Claims.(jwt.MapClaims)
		refreshClaim["firstName"] = registerCreds.FirstName
		refreshClaim["lastName"] = registerCreds.LastName
		refreshClaim["email"] = registerCreds.Email
		refreshClaim["exp"] = time.Now().Add(time.Hour * 730 * 12).Unix()

		t, err := token.SignedString([]byte("一给我里GIAO GIAO"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		r, err := refreshToken.SignedString([]byte("GIAO GIAO"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(map[string]string{
			"refresh": r,
			"token":   t,
		})
	}
	return c.Status(409).JSON(map[string]string{
		"error": "User exists",
	})
}
