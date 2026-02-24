package request

type CreatePostRequest struct {
	Title           string   `json:"title" binding:"required,max=200" label:"标题"`
	ContentMarkdown string   `json:"content_markdown" binding:"required" label:"内容Markdown"`
	ContentHTML     string   `json:"content_html" binding:"omitempty" label:"内容HTML"`
	Summary         string   `json:"summary" label:"简介"`
	Cover           string   `json:"cover" label:"封面"`
	Tags            []string `json:"tags" binding:"omitempty,dive,required" label:"标签"`
}

type UpdatePostRequest struct {
	Title           *string  `json:"title" binding:"omitempty,max=200" label:"标题"`
	ContentMarkdown *string  `json:"content_markdown" label:"内容Markdown"`
	ContentHTML     *string  `json:"content_html" binding:"omitempty" label:"内容HTML"`
	Summary         *string  `json:"summary" label:"简介"`
	Cover           *string  `json:"cover" label:"封面"`
	Tags            *[]string `json:"tags" binding:"omitempty,dive,required" label:"标签"`
}
