package authcontrollers

import (
	"log"
	"project/library_Management/models"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Message": "Status Ok",
	})
}
func createJWTToken(user models.User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = exp
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}
func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. %+\-]+@[a-z0-9.%+\-]+\.[a-z0-9.%+\-]`)
	return Re.MatchString(email)
}
func GetAllDataUser(c *fiber.Ctx) error {
	var user models.User
	if err := models.DB.Find(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(&user)
}
func SignUp(c *fiber.Ctx) error {
	// var data map[string]interface{}
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		log.Println("unable parse body")
	}

	// check password should >= 6 character
	if len(string(user.Password)) < 6 {
		log.Println(len(user.Password))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password must be greater than 6 character",
		})
	}
	if !ValidateEmail(string(user.Email)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Email Address",
		})
	}
	user.SetPassword(string(user.Password))
	models.DB.Where("email=?", (string(user.Email))).First(&user)
	if user.Id != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email already exist",
		})
	}
	if err := models.DB.Create(&user); err != nil {
		log.Println(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":    user,
		"message": "Account created successfully",
	})
}

type LoginRequest struct {
	Email    string
	Password string
}

func Login(c *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if req.Email == "" || req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
	}

	user := new(models.User)
	models.DB.Where("email = ?", req.Email).First(&user)

	if user.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Login Credential",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, exp, err := createJWTToken(*user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
}
