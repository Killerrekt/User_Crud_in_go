package api

import "github.com/gofiber/fiber/v2"

func Createuser(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"Route": "Working",
	})
}

func Updateuser(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"Route": "Working",
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
