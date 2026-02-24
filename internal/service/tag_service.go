package service

import (
	"strings"

	"jigsaw-api/internal/model"
	"jigsaw-api/internal/repository"
	"jigsaw-api/pkg/database"
)

type TagService struct {
	tagRepo     *repository.TagRepository
	taggingRepo *repository.TaggingRepository
}

func NewTagService() *TagService {
	return &TagService{
		tagRepo:     repository.NewTagRepository(),
		taggingRepo: repository.NewTaggingRepository(),
	}
}

func (s *TagService) Create(name string) (*model.Tag, error) {
	return s.tagRepo.FirstOrCreateByName(name)
}

func (s *TagService) GetByID(id uint) (*model.Tag, error) {
	return s.tagRepo.GetByID(id)
}

func (s *TagService) List(page, pageSize int) ([]model.Tag, int64, error) {
	return s.tagRepo.List(page, pageSize)
}

func (s *TagService) Update(tag *model.Tag) error {
	return s.tagRepo.Update(tag)
}

func (s *TagService) Delete(id uint) error {
	return s.tagRepo.Delete(id)
}

func (s *TagService) GetTagsForResource(resourceType string, resourceID uint) ([]model.Tag, error) {
	tagIDs, err := s.taggingRepo.FindTagIDsByResource(resourceType, resourceID)
	if err != nil {
		return nil, err
	}
	return s.tagRepo.FindByIDs(tagIDs)
}

func (s *TagService) SetTagsForResource(resourceType string, resourceID uint, tagNames []string) ([]model.Tag, error) {
	normalized := normalizeTagNames(tagNames)
	tx := database.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var tags []model.Tag
	for _, name := range normalized {
		var tag model.Tag
		if err := tx.Where("name = ?", name).FirstOrCreate(&tag, model.Tag{Name: name}).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		tags = append(tags, tag)
	}

	if err := tx.Where("resource_type = ? AND resource_id = ?", resourceType, resourceID).Delete(&model.Tagging{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var taggings []model.Tagging
	for _, tag := range tags {
		taggings = append(taggings, model.Tagging{
			TagID:        tag.ID,
			ResourceType: resourceType,
			ResourceID:   resourceID,
		})
	}
	if len(taggings) > 0 {
		if err := tx.Create(&taggings).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func normalizeTagNames(names []string) []string {
	seen := make(map[string]struct{})
	var result []string
	for _, name := range names {
		trimmed := strings.TrimSpace(name)
		if trimmed == "" {
			continue
		}
		if _, ok := seen[trimmed]; ok {
			continue
		}
		seen[trimmed] = struct{}{}
		result = append(result, trimmed)
	}
	return result
}
