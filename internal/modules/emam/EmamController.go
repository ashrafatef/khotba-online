package emam

import (
	"khotba-online/pkg/errors"
	"khotba-online/pkg/validation"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type EmamController struct {
	repo *EmamRepository
}

type CreateEmamRequest struct {
	Email     string `validate:"required,email" json:"email"`
	Password  string `validate:"required" json:"password"`
	FirstName string `validate:"required" json:"firstName"`
	LastName  string `validate:"required" json:"lastName"`
}

type LoginEmamRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

func NewEmamController(repo *EmamRepository) *EmamController {
	return &EmamController{
		repo: repo,
	}
}

func (con *EmamController) SignUp(c *fiber.Ctx) error {
	emam := new(CreateEmamRequest)

	if err := c.BodyParser(emam); err != nil {
		return err
	}

	if errs := validation.Validation(emam); len(errs) != 0 {
		return errors.NewValidationError(errs)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(emam.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = con.repo.CreateEmam(CreateEmamAttrs{
		Email:     emam.Email,
		Password:  string(hashedPassword),
		FirstName: emam.FirstName,
		LastName:  emam.LastName,
	})

	if err != nil {
		return err
	}

	return c.JSON("Successfully signed up")
}

func (con *EmamController) Login(c *fiber.Ctx) error {
	loginData := new(LoginEmamRequest)

	if err := c.BodyParser(loginData); err != nil {
		return err
	}

	if errs := validation.Validation(loginData); len(errs) != 0 {
		return errors.NewValidationError(errs)
	}

	emam, err := con.repo.GetEmamByEmail(loginData.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(emam.Password), []byte(loginData.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(emam.ID)),
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})

	token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		logrus.Println("Error generating token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
