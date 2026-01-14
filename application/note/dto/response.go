package dto

import "github.com/dromara/carbon/v2"

type NoteResponse struct {
	ID        uint64           `json:"id"`
	Content   string           `json:"content"`
	PinnedAt  *carbon.DateTime `json:"pinned_at"` // 置顶时间
	CreatedAt *carbon.DateTime `json:"created_at"`
	UpdatedAt *carbon.DateTime `json:"updated_at"`
}

type NotesResponse struct {
	Data  []*NoteResponse `json:"data"`
	Total int64           `json:"total"`
	Page  int             `json:"page"`
	Size  int             `json:"size"`
}
