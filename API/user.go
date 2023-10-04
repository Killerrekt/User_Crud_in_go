package api

import (
	database "github.com/Killerrekt/User_Crud_in_go/Database"
	"github.com/gofiber/fiber/v2"
)

func Createuser(c *fiber.Ctx) error {

	var details = new(database.User)

	var user []database.User
	err := c.BodyParser(details)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": "Something went wrong while parsing the body",
		})
	}

	var check string
	check = details.Username
	database.DB.Find(&user, "username = ?", check)

	if len(user) > 0 {
		return c.Status(400).JSON(&fiber.Map{
			"error": "User already exists",
		})
	}

	database.DB.Create(details)
	return c.Status(200).JSON(&fiber.Map{
		"data":    details,
		"success": true,
	})
}

func Updateuser(c *fiber.Ctx) error {

	var data = new(database.User)

	c.BodyParser(data)

	var check database.User

	database.DB.Find(&check, "username = ?", data.Username)

	if (check == database.User{}) {
		return c.Status(400).JSON(&fiber.Map{
			"error": "invalid username",
		})
	}

	check.Password = data.Password
	database.DB.Save(&check)

	return c.Status(200).JSON(&fiber.Map{
		"message": "password updated successfully",
	})
}

func Deleteuser(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"Route": "Working",
	})
}

func Getuser(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"Route": "Working",
	})
}

func Checkuser(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"Route": "Working",
	})
}
