package ledger_category

import (
	"fmt"
	"sort"
	"strings"

	"github.com/feilongjump/jigsaw-api/application/ledger_category/dto"
	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
)

type Service struct {
	repo repo.LedgerCategoryRepo
}

func NewService(repo repo.LedgerCategoryRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(userID uint64, req dto.CreateLedgerCategoryReq) error {
	category := &entity.LedgerCategory{
		UserID:   userID,
		ParentID: req.ParentID,
		Type:     req.Type,
		Name:     req.Name,
		Icon:     req.Icon,
		Sort:     req.Sort,
	}

	// 如果有父级，计算 Path
	if req.ParentID > 0 {
		parent, err := s.repo.GetLedgerCategory(req.ParentID, userID)
		if err != nil {
			return err // 父级不存在
		}
		// Path 格式: 父级Path + 父级ID + "-"
		// 假设父级 Path 是 "0-1-", 父级ID是 5, 则当前 Path 为 "0-1-5-"
		if parent.Path == "" {
			category.Path = fmt.Sprintf("0-%d-", parent.ID)
		} else {
			category.Path = fmt.Sprintf("%s%d-", parent.Path, parent.ID)
		}

		// 强制子类类型与父类一致
		category.Type = parent.Type
	} else {
		category.Path = "0-"
	}

	return s.repo.Create(category)
}

func (s *Service) FindLedgerCategories(userID uint64) ([]*dto.LedgerCategoryResp, error) {
	categories, err := s.repo.FindLedgerCategories(userID)
	if err != nil {
		return nil, err
	}

	return buildTree(categories, 0), nil
}

func (s *Service) GetLedgerCategory(userID uint64, id uint64) (*entity.LedgerCategory, error) {
	category, err := s.repo.GetLedgerCategory(id, userID)
	if err != nil {
		return nil, err
	}
	if category.UserID != userID {
		return nil, err_code.LedgerCategoryNotFound
	}
	return category, nil
}

// buildTree 递归构建树
func buildTree(all []*entity.LedgerCategory, parentID uint64) []*dto.LedgerCategoryResp {
	var tree []*dto.LedgerCategoryResp
	for _, c := range all {
		if c.ParentID == parentID {
			node := &dto.LedgerCategoryResp{
				ID:        c.ID,
				UserID:    c.UserID,
				ParentID:  c.ParentID,
				Type:      c.Type,
				Name:      c.Name,
				Icon:      c.Icon,
				Sort:      c.Sort,
				Path:      c.Path,
				CreatedAt: c.CreatedAt,
			}
			node.Children = buildTree(all, c.ID)
			tree = append(tree, node)
		}
	}
	// 排序: Sort 倒序
	sort.Slice(tree, func(i, j int) bool {
		return tree[i].Sort > tree[j].Sort
	})
	return tree
}

func (s *Service) Update(userID uint64, id uint64, req dto.UpdateLedgerCategoryReq) error {
	category, err := s.repo.GetLedgerCategory(id, userID)
	if err != nil {
		return err
	}
	if category.UserID != userID {
		return err_code.LedgerCategoryNotFound
	}

	updated := &entity.LedgerCategory{
		Name:     category.Name,
		Icon:     category.Icon,
		Sort:     category.Sort,
		ParentID: category.ParentID,
		Type:     category.Type,
		Path:     category.Path,
	}

	if req.Name != nil {
		updated.Name = *req.Name
	}
	if req.Icon != nil {
		updated.Icon = *req.Icon
	}
	if req.Sort != nil {
		updated.Sort = *req.Sort
	}

	parentChanged := false
	if req.ParentID != nil && *req.ParentID != category.ParentID {
		newParentID := *req.ParentID
		if newParentID == category.ID {
			return err_code.LedgerCategoryUpdateFailed
		}

		var newPath string
		if newParentID > 0 {
			parent, err := s.repo.GetLedgerCategory(newParentID, userID)
			if err != nil {
				return err
			}
			oldPrefix := fmt.Sprintf("%s%d-", category.Path, category.ID)
			if parent.ID == category.ID || strings.HasPrefix(parent.Path, oldPrefix) {
				return err_code.LedgerCategoryUpdateFailed
			}
			if parent.Path == "" {
				newPath = fmt.Sprintf("0-%d-", parent.ID)
			} else {
				newPath = fmt.Sprintf("%s%d-", parent.Path, parent.ID)
			}
			updated.Type = parent.Type
		} else {
			newPath = "0-"
		}

		updated.ParentID = newParentID
		updated.Path = newPath
		parentChanged = true
	}

	if err := s.repo.Update(id, userID, updated); err != nil {
		return err
	}

	if parentChanged {
		oldPrefix := fmt.Sprintf("%s%d-", category.Path, category.ID)
		newPrefix := fmt.Sprintf("%s%d-", updated.Path, category.ID)
		if oldPrefix != newPrefix {
			categories, err := s.repo.FindLedgerCategories(userID)
			if err != nil {
				return err
			}
			for _, c := range categories {
				if c.UserID != userID {
					continue
				}
				if strings.HasPrefix(c.Path, oldPrefix) {
					newChildPath := strings.Replace(c.Path, oldPrefix, newPrefix, 1)
					descendant := &entity.LedgerCategory{
						Name:     c.Name,
						Icon:     c.Icon,
						Sort:     c.Sort,
						ParentID: c.ParentID,
						Type:     c.Type,
						Path:     newChildPath,
					}
					if err := s.repo.Update(c.ID, userID, descendant); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (s *Service) Delete(userID uint64, id uint64) error {
	// TODO: 检查是否有子分类？检查是否被使用？
	// 这里简化处理，直接删除
	err, row := s.repo.Delete(id, userID)
	if err != nil {
		return err
	}
	if row == 0 {
		return err_code.LedgerCategoryDeleteFailed
	}
	return nil
}
