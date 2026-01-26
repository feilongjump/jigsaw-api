package user_wallet

import (
	"context"
	"encoding/json"
	"fmt"

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

func (s *Service) Create(ctx context.Context, userID uint64, req dto.CreateUserWalletReq) (*dto.UserWalletResp, error) {
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
	if err := s.repo.Create(ctx, wallet); err != nil {
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

func (s *Service) Update(ctx context.Context, userID uint64, id uint64, req dto.UpdateUserWalletReq) error {
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
		if err := s.validateExtraConfig(wallet.Type, *req.ExtraConfig); err != nil {
			return err
		}
		updated.ExtraConfig = *req.ExtraConfig
	}

	return s.repo.Update(ctx, id, userID, updated)
}

func (s *Service) validateExtraConfig(walletType uint8, config entity.JSONMap) error {
	if config == nil {
		return nil
	}

	// 序列化为 JSON 以便映射到结构体
	data, err := json.Marshal(config)
	if err != nil {
		return err_code.UserWalletConfigInvalid
	}

	switch walletType {
	case entity.UserWalletTypeCreditCard:
		var cfg entity.CreditCardConfig
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err_code.UserWalletConfigInvalid
		}
		// 校验账单日和还款日
		if cfg.BillDay < 1 || cfg.BillDay > 31 {
			return fmt.Errorf("%w: 账单日必须在 1-31 之间", err_code.UserWalletConfigInvalid)
		}
		if cfg.RepaymentDay < 1 || cfg.RepaymentDay > 31 {
			return fmt.Errorf("%w: 还款日必须在 1-31 之间", err_code.UserWalletConfigInvalid)
		}

	case entity.UserWalletTypeInvestment, entity.UserWalletTypeMargin:
		var cfg entity.InvestmentConfig
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err_code.UserWalletConfigInvalid
		}
		// 校验费率规则
		for i, rule := range cfg.Rules {
			if rule.CommissionRate < 0 {
				return fmt.Errorf("%w: 第 %d 条规则佣金费率不能为负数", err_code.UserWalletConfigInvalid, i+1)
			}
			// 可以在这里增加更多针对 rule 的校验，例如 Market 或 Type 是否必填
		}
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, userID uint64, id uint64) error {
	count, err := s.recordRepo.CountByWalletID(userID, id)
	if err != nil {
		return err
	}
	if count > 0 {
		return err_code.UserWalletDeleteForbidden
	}

	err, row := s.repo.Delete(ctx, id, userID)
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
