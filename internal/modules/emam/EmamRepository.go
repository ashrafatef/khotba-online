package emam

import (
	"khotba-online/internal/database/models"

	"gorm.io/gorm"
)

type EmamRepository struct {
	db *gorm.DB
}

func NewEmamRepository(db *gorm.DB) *EmamRepository {
	return &EmamRepository{
		db: db,
	}
}

type CreateEmamAttrs struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (r *EmamRepository) CreateEmam(emam CreateEmamAttrs) (models.Emam, error) {
	createdEmam := &models.Emam{
		Email:     emam.Email,
		Password:  emam.Password,
		FirstName: emam.FirstName,
		LastName:  emam.LastName,
	}
	results := r.db.Create(&createdEmam).Error
	if results != nil {
		return models.Emam{}, results
	}
	return *createdEmam, nil
}


func (r *EmamRepository) GetEmamByEmail(email string) (models.Emam, error) {
	var emam models.Emam
	results := r.db.Where("email = ?", email).First(&emam).Error
	if results != nil {
		return models.Emam{}, results
	}
	return emam, nil
}
