package masjeds

import (
	"khotba-online/internal/database/models"

	"gorm.io/gorm"
)

type MasjedRepository struct {
	db *gorm.DB
}

func NewMasjedRepository(db *gorm.DB) *MasjedRepository {
	return &MasjedRepository{
		db: db,
	}
}

type CreateMasjedAttrs struct {
	Name      string              `json:"name"`
	Languages []string            `json:"languages"`
	EmamId    int                 `json:"emam_id"`
	Status    models.MasjedStatus `json:"status"`
	Location  interface{}         `json:"location"`
}

func (repo *MasjedRepository) CreateMasjed(attrs CreateMasjedAttrs) (models.Masjed, error) {

	createdMasjed := &models.Masjed{
		Name: attrs.Name,
		// Languages: attrs.Languages,
		EmamId:   attrs.EmamId,
		Status:   attrs.Status,
		Location: attrs.Location,
	}
	results := repo.db.Create(&createdMasjed).Error
	if results != nil {
		return models.Masjed{}, results
	}
	return *createdMasjed, nil
}

func (repo *MasjedRepository) GetMasjedByEmamId(emamId int) (models.Masjed, error) {
	var masjed models.Masjed

	result := repo.db.Where("emam_id = ?", emamId).First(&masjed)
	if result.Error != nil {
		return models.Masjed{}, result.Error
	}

	return masjed, nil
}
