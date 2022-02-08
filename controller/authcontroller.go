package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/project104/Database"
	"github.com/project104/Models"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

var Globallist Models.List

const secretkey = "secret"

func Login(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user Models.User
	Database.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "email not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})

	}

	//		to check the login page
	//		return c.JSON(user)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{

		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	})

	token, err := claims.SignedString([]byte(secretkey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "login failed"})

	}

	//	to check token
	//	return c.JSON(token)

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 40),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(Globallist)
	//	return c.JSON(fiber.Map{"message": "success"})

}

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := Models.User{

		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	Database.DB.Create(&user)

	return c.JSON(user)
}

func Show(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var todo Models.List
	return c.JSON(todo)
}

func User(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretkey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "login failed",
		})

	}

	claims := token.Claims.(*jwt.StandardClaims)
	var user Models.User
	Database.DB.Where("ID=?", claims.Issuer).First(&user)
	return c.JSON(user)

}

func Create(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	todolist := Models.List{

		Task: data["task"],
	}
	Database.DB.Create(&todolist)
	return c.JSON(todolist)
}

func Update(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	todolist := Models.List{
		Task: data["task"],
	}
	if todolist.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "ID  cannot be zero",
		})
	} else {

		// Update with map
		Database.DB.Table("lists").Where("ID = ?", todolist.ID).Updates(map[string]interface{}{"task": todolist.Task})
		// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);

	}

	//Database.DB.Create(&todolist)
	return c.JSON(todolist)
}

func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Minute * 40),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})

}
