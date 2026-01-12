package dto

import "github.com/dromara/carbon/v2"

type FileResponse struct {
	Name      string           `json:"name"`
	Url       string           `json:"url"`
	Size      int64            `json:"size"`
	MimeType  string           `json:"mime_type"`
	OwnerType string           `json:"owner_type"`
	OwnerID   uint64           `json:"owner_id"`
	CreatedAt *carbon.DateTime `json:"created_at"`
}
