package repository

import (
	"jigsaw-api/internal/model"
	"jigsaw-api/pkg/database"

	"gorm.io/gorm"
)

type TaggingRepository struct {
	db *gorm.DB
}

func NewTaggingRepository() *TaggingRepository {
	return &TaggingRepository{
		db: database.DB,
	}
}

func (r *TaggingRepository) DeleteByResource(resourceType string, resourceID uint) error {
	return r.db.Where("resource_type = ? AND resource_id = ?", resourceType, resourceID).Delete(&model.Tagging{}).Error
}

func (r *TaggingRepository) CreateBulk(taggings []model.Tagging) error {
	if len(taggings) == 0 {
		return nil
	}
	return r.db.Create(&taggings).Error
}

func (r *TaggingRepository) FindTagIDsByResource(resourceType string, resourceID uint) ([]uint, error) {
	var tagIDs []uint
	err := r.db.Model(&model.Tagging{}).
		Where("resource_type = ? AND resource_id = ?", resourceType, resourceID).
		Pluck("tag_id", &tagIDs).Error
	if err != nil {
		return nil, err
	}
	return tagIDs, nil
}
