package router

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/tamthanh/go-app/database"
	"github.com/tamthanh/go-app/model"
)

type Product struct {
	Id       uint   `json:"id" `
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func CreateResponseProduct(product model.Product) Product {
	return Product{
		Id:       product.Id,
		Question: product.Question,
		Answer:   product.Answer,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func Getproducts(c *fiber.Ctx) error {
	products := []model.Product{}
	database.Database.Db.Find(&products)
	responseProducts := []Product{}
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}
func findProduct(id int, product *model.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product model.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseProduct(product)

	return c.Status(200).JSON(responseUser)
}
func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product model.Product
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	err = findProduct(id, &product)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateProduct struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}

	var updateData UpdateProduct
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	product.Question = updateData.Question
	product.Answer = updateData.Answer

	database.Database.Db.Save((&product))
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}
func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product model.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findProduct(id, &product)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted User")
}
