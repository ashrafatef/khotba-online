package masjeds

import (
	"khotba-online/internal/database/models"
	"khotba-online/pkg/errors"
	"khotba-online/pkg/validation"

	"github.com/gofiber/fiber/v2"
)

type CreateMasjedRequest struct {
	Name      string              `validate:"required" json:"name"`
	Languages []string            `validate:"required" json:"languages"`
	EmamId    int                 `validate:"required" json:"emamId"`
	Status    models.MasjedStatus `validate:"required" json:"status"`
	Location  interface{}         `validate:"required" json:"location"`
}

type MasjedController struct {
	repo *MasjedRepository
}

func NewMasjedController(repo *MasjedRepository) *MasjedController {
	return &MasjedController{
		repo: repo,
	}
}

func (con *MasjedController) Create(c *fiber.Ctx) error {
	masjedReq := new(CreateMasjedRequest)

	if err := c.BodyParser(masjedReq); err != nil {
		return err
	}

	if errs := validation.Validation(masjedReq); len(errs) != 0 {
		return errors.NewValidationError(errs)
	}

	createdMasjed, err := con.repo.CreateMasjed(CreateMasjedAttrs{
		Name:      masjedReq.Name,
		Languages: masjedReq.Languages,
		EmamId:    masjedReq.EmamId,
		Status:    models.MasjedStatus(masjedReq.Status),
		Location:  masjedReq.Location,
	})

	if err != nil {
		return err
	}

	return c.JSON(createdMasjed)
}
