package service

import (
	"jigsaw-api/internal/model"
	"jigsaw-api/internal/repository"
	"jigsaw-api/pkg/database"
)

const PostResourceType = "post"

type PostService struct {
	postRepo   *repository.PostRepository
	tagService *TagService
}

func NewPostService() *PostService {
	return &PostService{
		postRepo:   repository.NewPostRepository(),
		tagService: NewTagService(),
	}
}

func (s *PostService) Create(post *model.Post, tagNames []string) (*model.Post, error) {
	if err := s.postRepo.Create(post); err != nil {
		return nil, err
	}
	if len(tagNames) > 0 {
		tags, err := s.tagService.SetTagsForResource(PostResourceType, post.ID, tagNames)
		if err != nil {
			return nil, err
		}
		post.Tags = tags
	}
	return post, nil
}

func (s *PostService) GetByID(id uint) (*model.Post, error) {
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	tags, err := s.tagService.GetTagsForResource(PostResourceType, post.ID)
	if err != nil {
		return nil, err
	}
	post.Tags = tags
	return post, nil
}

func (s *PostService) FindByUser(userID uint, page, pageSize int) ([]model.Post, int64, error) {
	posts, total, err := s.postRepo.FindByUser(userID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	for i := range posts {
		tags, err := s.tagService.GetTagsForResource(PostResourceType, posts[i].ID)
		if err != nil {
			return nil, 0, err
		}
		posts[i].Tags = tags
	}
	return posts, total, nil
}

func (s *PostService) Update(post *model.Post, tagNames *[]string) (*model.Post, error) {
	if err := s.postRepo.Update(post); err != nil {
		return nil, err
	}
	if tagNames != nil {
		tags, err := s.tagService.SetTagsForResource(PostResourceType, post.ID, *tagNames)
		if err != nil {
			return nil, err
		}
		post.Tags = tags
	} else {
		tags, err := s.tagService.GetTagsForResource(PostResourceType, post.ID)
		if err != nil {
			return nil, err
		}
		post.Tags = tags
	}
	return post, nil
}

func (s *PostService) Delete(id uint) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Delete(&model.Post{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("resource_type = ? AND resource_id = ?", PostResourceType, id).Delete(&model.Tagging{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
