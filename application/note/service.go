package note

import (
	"github.com/feilongjump/jigsaw-api/application/note/dto"
	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
)

// Service Note 应用服务
type Service struct {
	noteRepo repo.NoteRepository
	fileRepo repo.FileRepository
}

// NewNoteService 创建 Note 应用服务
func NewNoteService(noteRepo repo.NoteRepository, fileRepo repo.FileRepository) *Service {
	return &Service{
		noteRepo: noteRepo,
		fileRepo: fileRepo,
	}
}

// Create 创建 Note
func (s *Service) Create(req *dto.CreateNoteRequest, userID uint64) (*dto.NoteResponse, error) {
	// 转换为领域实体
	note := &entity.Note{
		UserID:  userID,
		Content: req.Content,
	}

	// 调用仓储进行存储
	if err := s.noteRepo.Create(note); err != nil {
		return nil, err
	}

	// 绑定文件 (如果有)
	if len(req.FileIDs) > 0 {
		if err := s.fileRepo.BindFiles(req.FileIDs, userID, "notes", note.ID); err != nil {
			// 记录日志，但不中断流程? 或者返回错误?
			// 这里如果绑定失败，Note 已经创建成功了。
			// 最好是事务一致性。但 Repository 模式通常不跨 Repo 事务。
			// 简单起见，返回错误，用户可以重试或者看到 Note 没有附件。
			return nil, err
		}
	}

	return &dto.NoteResponse{
		ID:        note.ID,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}, nil
}

func (s *Service) GetNote(id, userID uint64) (*dto.NoteResponse, error) {
	// 调用仓储查询 Note
	note, err := s.noteRepo.GetNote(id, userID)
	if err != nil {
		return nil, err
	}

	return &dto.NoteResponse{
		ID:        note.ID,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}, nil
}

func (s *Service) FindNotes(page, size int, userID uint64) (*dto.NotesResponse, error) {
	// 调用仓储查询 Note
	notes, total, err := s.noteRepo.FindNotes(page, size, userID)
	if err != nil {
		return nil, err
	}

	data := make([]*dto.NoteResponse, 0, len(notes))
	for _, note := range notes {
		data = append(data, &dto.NoteResponse{
			ID:        note.ID,
			Content:   note.Content,
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
		})
	}

	return &dto.NotesResponse{
		Data:  data,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (s *Service) Update(id, userID uint64, req *dto.UpdateNoteRequest) error {
	note := &entity.Note{
		Content: req.Content,
	}

	if err := s.noteRepo.Update(id, userID, note); err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(id, userID uint64) error {
	err, row := s.noteRepo.Delete(id, userID)
	if err != nil {
		return err
	}

	if row == 0 {
		// 未删除任何数据，可能是 Note 不存在
		return err_code.NoteDeleteFailed
	}

	return nil
}
