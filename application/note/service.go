package note

import (
	"github.com/feilongjump/jigsaw-api/application/note/dto"
	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
)

// Service Note 应用服务
type Service struct {
	noteRepo repo.NoteRepository
}

// NewNoteService 创建 Note 应用服务
func NewNoteService(noteRepo repo.NoteRepository) *Service {
	return &Service{
		noteRepo: noteRepo,
	}
}

// Create 创建 Note
func (s *Service) Create(req *dto.CreateNoteRequest) (*dto.NoteResponse, error) {
	// 转换为领域实体
	note := &entity.Note{
		Content: req.Content,
	}

	// 调用仓储进行存储
	if err := s.noteRepo.Create(note); err != nil {
		return nil, err
	}

	return &dto.NoteResponse{
		ID:        note.ID,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}, nil
}

func (s *Service) GetNote(id uint64) (*dto.NoteResponse, error) {
	// 调用仓储查询 Note
	note, err := s.noteRepo.GetNote(id)
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

func (s *Service) FindNotes(page, size int) (*dto.NotesResponse, error) {
	// 调用仓储查询 Note
	notes, total, err := s.noteRepo.FindNotes(page, size)
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

func (s *Service) Update(id uint64, req *dto.UpdateNoteRequest) error {
	note := &entity.Note{
		Content: req.Content,
	}

	if err := s.noteRepo.Update(id, note); err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(id uint64) error {
	if err := s.noteRepo.Delete(id); err != nil {
		return err
	}

	return nil
}
