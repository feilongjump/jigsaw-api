package request

type CreateTagRequest struct {
	Name string `json:"name" binding:"required,max=50" label:"标签名"`
}

type UpdateTagRequest struct {
	Name string `json:"name" binding:"required,max=50" label:"标签名"`
}
