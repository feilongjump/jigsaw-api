package dto

type IndexNoteRequest struct {
	Page int `form:"page" binding:"required,min=1" label:"页码"`
	Size int `form:"size" binding:"required,min=1,max=100" label:"每页数量"`
}

type NoteURIRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" label:"笔记 ID"`
}

type CreateNoteRequest struct {
	Content string   `json:"content" binding:"required" label:"内容"`
	FileIDs []uint64 `json:"file_ids"`
}

type UpdateNoteRequest struct {
	Content string `json:"content" binding:"required" label:"内容"`
}
