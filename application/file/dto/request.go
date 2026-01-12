package dto

type UploadFileRequest struct {
	OwnerType string `form:"owner_type" binding:"required,oneof=users notes"`
	OwnerID   uint64 `form:"owner_id"`
}

type DeleteFileRequest struct {
	Path      string `json:"path" binding:"required"`
	OwnerType string `json:"owner_type" binding:"required,oneof=users notes"`
	OwnerID   uint64 `json:"owner_id" binding:"required"`
}
