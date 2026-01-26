package ledger_record

import (
	"context"

	"github.com/feilongjump/jigsaw-api/application/ledger_record/dto"
	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
)

type Service struct {
	repo       repo.LedgerRecordRepo
	walletRepo repo.UserWalletRepo
	tm         repo.TransactionManager
}

func NewService(repo repo.LedgerRecordRepo, walletRepo repo.UserWalletRepo, tm repo.TransactionManager) *Service {
	return &Service{
		repo:       repo,
		walletRepo: walletRepo,
		tm:         tm,
	}
}

func (s *Service) Create(ctx context.Context, userID uint64, req dto.CreateLedgerRecordReq) (*dto.LedgerRecordResp, error) {
	record := &entity.LedgerRecord{
		UserID:         userID,
		Type:           req.Type,
		Amount:         req.Amount,
		SourceWalletID: req.SourceWalletID,
		TargetWalletID: req.TargetWalletID,
		CategoryID:     req.CategoryID,
		OccurredAt:     req.OccurredAt,
		Remark:         req.Remark,
		Images:         req.Images,
	}

	err := s.tm.Transaction(ctx, func(ctx context.Context) error {
		// 1. 创建记录
		if err := s.repo.Create(ctx, record); err != nil {
			return err
		}

		// 2. 更新账户余额
		switch req.Type {
		case 1: // 支出
			if req.SourceWalletID > 0 {
				if err := s.handleOutflow(ctx, req.SourceWalletID, req.Amount, userID); err != nil {
					return err
				}
			}
		case 2: // 收入
			if req.TargetWalletID > 0 {
				if err := s.handleInflow(ctx, req.TargetWalletID, req.Amount, userID); err != nil {
					return err
				}
			}
		case 3: // 转账
			if req.SourceWalletID > 0 {
				if err := s.handleOutflow(ctx, req.SourceWalletID, req.Amount, userID); err != nil {
					return err
				}
			}
			if req.TargetWalletID > 0 {
				if err := s.handleInflow(ctx, req.TargetWalletID, req.Amount, userID); err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return toLedgerRecordResp(record), nil
}

func (s *Service) Update(ctx context.Context, userID uint64, id uint64, req dto.UpdateLedgerRecordReq) (*dto.LedgerRecordResp, error) {
	var updatedRecord *entity.LedgerRecord

	err := s.tm.Transaction(ctx, func(ctx context.Context) error {
		record, err := s.repo.GetLedgerRecord(id, userID)
		if err != nil {
			return err
		}

		updated := &entity.LedgerRecord{
			UserID:         userID,
			Type:           req.Type,
			Amount:         req.Amount,
			SourceWalletID: req.SourceWalletID,
			TargetWalletID: req.TargetWalletID,
			CategoryID:     req.CategoryID,
			OccurredAt:     req.OccurredAt,
			Remark:         req.Remark,
			Images:         req.Images,
		}

		if err := s.repo.Update(ctx, id, userID, updated); err != nil {
			return err
		}

		if err := s.rollbackLedgerRecord(ctx, record, userID); err != nil {
			return err
		}

		if err := s.applyLedgerRecord(ctx, updated, userID); err != nil {
			return err
		}

		// Re-fetch within transaction to ensure we get the latest state
		updatedRecord, err = s.repo.GetLedgerRecord(id, userID)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return toLedgerRecordResp(updatedRecord), nil
}

func (s *Service) Delete(ctx context.Context, userID uint64, id uint64) error {
	return s.tm.Transaction(ctx, func(ctx context.Context) error {
		// 1. 查询记录
		record, err := s.repo.GetLedgerRecord(id, userID)
		if err != nil {
			return err_code.LedgerRecordDeleteFailed
		}

		// 2. 删除记录
		err, row := s.repo.Delete(ctx, id, userID)
		if err != nil {
			return err
		}
		if row == 0 {
			return err_code.LedgerRecordDeleteFailed
		}

		// 3. 回滚余额
		if err := s.rollbackLedgerRecord(ctx, record, userID); err != nil {
			return err
		}

		return nil
	})
}

func (s *Service) rollbackLedgerRecord(ctx context.Context, record *entity.LedgerRecord, userID uint64) error {
	amount := record.Amount
	switch record.Type {
	case 1:
		if record.SourceWalletID > 0 {
			if err := s.handleInflow(ctx, record.SourceWalletID, amount, userID); err != nil {
				return err
			}
		}
	case 2:
		if record.TargetWalletID > 0 {
			if err := s.handleOutflow(ctx, record.TargetWalletID, amount, userID); err != nil {
				return err
			}
		}
	case 3:
		if record.SourceWalletID > 0 {
			if err := s.handleInflow(ctx, record.SourceWalletID, amount, userID); err != nil {
				return err
			}
		}
		if record.TargetWalletID > 0 {
			if err := s.handleOutflow(ctx, record.TargetWalletID, amount, userID); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Service) applyLedgerRecord(ctx context.Context, record *entity.LedgerRecord, userID uint64) error {
	switch record.Type {
	case 1:
		if record.SourceWalletID > 0 {
			if err := s.handleOutflow(ctx, record.SourceWalletID, record.Amount, userID); err != nil {
				return err
			}
		}
	case 2:
		if record.TargetWalletID > 0 {
			if err := s.handleInflow(ctx, record.TargetWalletID, record.Amount, userID); err != nil {
				return err
			}
		}
	case 3:
		if record.SourceWalletID > 0 {
			if err := s.handleOutflow(ctx, record.SourceWalletID, record.Amount, userID); err != nil {
				return err
			}
		}
		if record.TargetWalletID > 0 {
			if err := s.handleInflow(ctx, record.TargetWalletID, record.Amount, userID); err != nil {
				return err
			}
		}
	}
	return nil
}

// handleOutflow 资金流出
func (s *Service) handleOutflow(ctx context.Context, walletID uint64, amount float64, userID uint64) error {
	wallet, err := s.walletRepo.GetUserWallet(walletID, userID)
	if err != nil {
		return err
	}

	// Type 5: 信用卡, Type 8: 两融账户
	if wallet.Type == entity.UserWalletTypeCreditCard || wallet.Type == entity.UserWalletTypeMargin {
		return s.walletRepo.UpdateLiability(ctx, walletID, userID, amount) // 负债增加
	}
	return s.walletRepo.UpdateBalance(ctx, walletID, userID, -amount) // 余额减少
}

// handleInflow 资金流入
func (s *Service) handleInflow(ctx context.Context, walletID uint64, amount float64, userID uint64) error {
	wallet, err := s.walletRepo.GetUserWallet(walletID, userID)
	if err != nil {
		return err
	}

	// Type 5: 信用卡, Type 8: 两融账户
	if wallet.Type == entity.UserWalletTypeCreditCard || wallet.Type == entity.UserWalletTypeMargin {
		return s.walletRepo.UpdateLiability(ctx, walletID, userID, -amount) // 负债减少
	}
	return s.walletRepo.UpdateBalance(ctx, walletID, userID, amount) // 余额增加
}

func (s *Service) FindLedgerRecords(userID uint64, req dto.ListLedgerRecordReq) (*dto.LedgerRecordListResponse, error) {
	filter := map[string]interface{}{}
	if req.Type > 0 {
		filter["type"] = req.Type
	}
	if req.WalletID > 0 {
		filter["wallet_id"] = req.WalletID
	}
	if req.CategoryID > 0 {
		filter["category_id"] = req.CategoryID
	}

	records, total, err := s.repo.FindLedgerRecords(userID, req.Page, req.PageSize, filter)
	if err != nil {
		return nil, err
	}

	var resp []*dto.LedgerRecordResp
	for _, r := range records {
		resp = append(resp, toLedgerRecordResp(r))
	}

	return &dto.LedgerRecordListResponse{
		Data:  resp,
		Total: total,
		Page:  req.Page,
		Size:  req.PageSize,
	}, nil
}

func toLedgerRecordResp(record *entity.LedgerRecord) *dto.LedgerRecordResp {
	return &dto.LedgerRecordResp{
		ID:             record.ID,
		UserID:         record.UserID,
		Type:           record.Type,
		Amount:         record.Amount,
		SourceWalletID: record.SourceWalletID,
		TargetWalletID: record.TargetWalletID,
		CategoryID:     record.CategoryID,
		OccurredAt:     record.OccurredAt,
		Remark:         record.Remark,
		Images:         record.Images,
		CreatedAt:      record.CreatedAt,
	}
}
