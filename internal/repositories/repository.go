package repositories

import (
	"dedobot/internal/models"
	"gorm.io/gorm"
)

type SkufRepo struct {
	db *gorm.DB
}

func NewSkufRepo(db *gorm.DB) *SkufRepo {
	return &SkufRepo{db}
}

func (r *SkufRepo) GetByUserID(userID int64) (*models.Skuf, error) {
	var skuf models.Skuf
	err := r.db.Where("user_id = ?", userID).First(&skuf).Error
	return &skuf, err
}

func (r *SkufRepo) Create(skuf *models.Skuf) error {
	return r.db.Create(skuf).Error
}

func (r *SkufRepo) Update(skuf *models.Skuf) error {
	return r.db.Save(skuf).Error
}

func (r *SkufRepo) GetAll() ([]models.Skuf, error) {
	var skufs []models.Skuf
	err := r.db.Find(&skufs).Error
	return skufs, err
}
