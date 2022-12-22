package productcontrollers

import (
	"project/library_Management/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllEmployee(c *fiber.Ctx) error {
	var employee []models.Employee
	models.DB.Find(&employee)

	return c.Status(fiber.StatusOK).JSON(employee)
}
func GetEmployeebyid(c *fiber.Ctx) error {
	var employee models.Employee
	id := c.Params("id")
	if err := models.DB.First(&employee, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data not found",
			})
		}
	}
	return c.JSON(employee)
}
func CreateEmployee(c *fiber.Ctx) error {
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := models.DB.Create(&employee).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(employee)

}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if models.DB.Where("id=?", id).Updates(&employee).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "data successfully update",
	})
}

func DeleteEmployee(c *fiber.Ctx) error {
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if models.DB.Delete(&employee, employee.Id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})

	}
	return c.JSON(fiber.Map{
		"message": "data successfully delete",
	})

}
