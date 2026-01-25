package user_wallet

import (
	"github.com/feilongjump/jigsaw-api/application/user_wallet/dto"
	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
)

type Service struct {
	repo       repo.UserWalletRepo
	recordRepo repo.LedgerRecordRepo
}

func NewService(repo repo.UserWalletRepo, recordRepo repo.LedgerRecordRepo) *Service {
	return &Service{
		repo:       repo,
		recordRepo: recordRepo,
	}
}

func (s *Service) Create(userID uint64, req dto.CreateUserWalletReq) (*dto.UserWalletResp, error) {
	wallet := &entity.UserWallet{
		UserID:      userID,
		Name:        req.Name,
		Type:        req.Type,
		Balance:     req.Balance,
		Liability:   req.Liability,
		Remark:      req.Remark,
		Sort:        req.Sort,
		ExtraConfig: req.ExtraConfig,
	}
	if err := s.repo.Create(wallet); err != nil {
		return nil, err
	}
	return toUserWalletResp(wallet), nil
}

func (s *Service) FindUserWallets(userID uint64) (*dto.UserWalletListResponse, error) {
	wallets, err := s.repo.FindUserWallets(userID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.UserWalletResp
	for _, w := range wallets {
		resp = append(resp, toUserWalletResp(w))
	}

	return &dto.UserWalletListResponse{
		Data:  resp,
		Total: int64(len(resp)),
		Page:  1,
		Size:  len(resp),
	}, nil
}

func (s *Service) Update(userID uint64, id uint64, req dto.UpdateUserWalletReq) error {
	wallet, err := s.repo.GetUserWallet(id, userID)
	if err != nil {
		return err
	}

	updated := &entity.UserWallet{
		Name:        wallet.Name,
		Type:        wallet.Type,
		Balance:     wallet.Balance,
		Liability:   wallet.Liability,
		Remark:      wallet.Remark,
		Sort:        wallet.Sort,
		IsHidden:    wallet.IsHidden,
		ExtraConfig: wallet.ExtraConfig,
	}

	if req.Name != nil {
		updated.Name = *req.Name
	}
	if req.Remark != nil {
		updated.Remark = *req.Remark
	}
	if req.Sort != nil {
		updated.Sort = *req.Sort
	}
	if req.IsHidden != nil {
		updated.IsHidden = *req.IsHidden
	}
	if req.ExtraConfig != nil {
		updated.ExtraConfig = *req.ExtraConfig
	}

	return s.repo.Update(id, userID, updated)
}

func (s *Service) Delete(userID uint64, id uint64) error {
	count, err := s.recordRepo.CountByWalletID(userID, id)
	if err != nil {
		return err
	}
	if count > 0 {
		return err_code.UserWalletDeleteForbidden
	}

	err, row := s.repo.Delete(id, userID)
	if err != nil {
		return err
	}

	if row == 0 {
		// 未删除任何数据，可能是 UserWallet 不存在
		return err_code.UserWalletDeleteFailed
	}

	return nil
}

func toUserWalletResp(wallet *entity.UserWallet) *dto.UserWalletResp {
	return &dto.UserWalletResp{
		ID:          wallet.ID,
		UserID:      wallet.UserID,
		Name:        wallet.Name,
		Type:        wallet.Type,
		Balance:     wallet.Balance,
		Liability:   wallet.Liability,
		Remark:      wallet.Remark,
		Sort:        wallet.Sort,
		IsHidden:    wallet.IsHidden,
		ExtraConfig: wallet.ExtraConfig,
		CreatedAt:   wallet.CreatedAt,
	}
}
