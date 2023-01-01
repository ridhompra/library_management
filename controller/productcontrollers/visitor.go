package productcontrollers

import (
	"project/library_Management/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllVisitor(c *fiber.Ctx) error {
	var visitor []models.Visitor
	models.DB.Find(&visitor)

	return c.Status(fiber.StatusOK).JSON(visitor)
}
func GetVisitorById(c *fiber.Ctx) error {
	var visitor models.Visitor
	id := c.Params("id")
	if err := models.DB.First(&visitor, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data not found",
			})
		}
	}
	return c.JSON(visitor)
}
func CreateVisitor(c *fiber.Ctx) error {
	var visitor models.Visitor
	if err := c.BodyParser(&visitor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	visitor.CreateAt = time.Now()
	if err := models.DB.Create(&visitor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(visitor)

}

func UpdateVisitor(c *fiber.Ctx) error {
	id := c.Params("id")
	var visitor models.Visitor
	if err := c.BodyParser(&visitor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if models.DB.Where("id=?", id).Updates(&visitor).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "data successfully update",
	})
}

func DeleteVisitor(c *fiber.Ctx) error {
	var visitor models.Visitor
	if err := c.BodyParser(&visitor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if models.DB.Delete(&visitor, visitor.Id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})

	}
	return c.JSON(fiber.Map{
		"message": "data successfully delete ",
	})

}
