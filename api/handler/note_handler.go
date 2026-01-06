package handler

import (
	"github.com/feilongjump/jigsaw-api/application/note"
	"github.com/feilongjump/jigsaw-api/application/note/dto"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"github.com/feilongjump/jigsaw-api/pkg/gin_util"
	"github.com/feilongjump/jigsaw-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	noteService *note.Service
}

func NewNoteHandler(noteService *note.Service) *NoteHandler {
	return &NoteHandler{
		noteService: noteService,
	}
}

// Create 创建 Note
func (n *NoteHandler) Create(c *gin.Context) {
	var req dto.CreateNoteRequest
	if !gin_util.BindJSON(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	noteResp, err := n.noteService.Create(&req, userID)
	if err != nil {
		response.Fail(c, err_code.NoteCreateFailed)
		return
	}

	response.Success(c, noteResp)
}

// Show 查询 Note
func (n *NoteHandler) Show(c *gin.Context) {
	var req dto.NoteURIRequest
	if !gin_util.BindURI(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	noteResp, err := n.noteService.GetNote(req.ID, userID)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.Success(c, noteResp)
}

// Index 查询 Note 列表
func (n *NoteHandler) Index(c *gin.Context) {
	var req dto.IndexNoteRequest
	if !gin_util.BindQuery(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	notesResp, err := n.noteService.FindNotes(req.Page, req.Size, userID)
	if err != nil {
		response.Fail(c, err_code.NoteGetFailed)
		return
	}

	response.Success(c, notesResp)
}

// Update 更新 Note
func (n *NoteHandler) Update(c *gin.Context) {
	var reqURI dto.NoteURIRequest
	if !gin_util.BindURI(c, &reqURI) {
		return
	}

	var req dto.UpdateNoteRequest
	if !gin_util.BindJSON(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	if err := n.noteService.Update(reqURI.ID, userID, &req); err != nil {
		response.Fail(c, err_code.NoteUpdateFailed)
		return
	}

	// 返回更新后的 Note
	updatedNote, err := n.noteService.GetNote(reqURI.ID, userID)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.Success(c, updatedNote)
}

// Delete 删除 Note
func (n *NoteHandler) Delete(c *gin.Context) {
	var req dto.NoteURIRequest
	if !gin_util.BindURI(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	if err := n.noteService.Delete(req.ID, userID); err != nil {
		response.Fail(c, err_code.NoteDeleteFailed)
		return
	}

	response.Success(c, nil)
}
