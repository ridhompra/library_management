package productcontrollers

import (
	"project/library_Management/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllBook(c *fiber.Ctx) error {
	var books []models.Book
	models.DB.Find(&books)

	return c.Status(fiber.StatusOK).JSON(books)
}
func GetBookbyid(c *fiber.Ctx) error {
	var book models.Book
	id := c.Params("id")
	if err := models.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data not found",
			})
		}
	}
	return c.JSON(book)
}
func CreateBook(c *fiber.Ctx) error {
	var books models.Book
	if err := c.BodyParser(&books); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := models.DB.Create(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(books)

}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if models.DB.Where("id=?", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "data successfully update",
	})
}

func DeleteBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if models.DB.Delete(&book, book.Id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})

	}
	return c.JSON(fiber.Map{
		"message": "data successfully delete",
	})

}
