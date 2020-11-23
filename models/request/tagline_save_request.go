package request

import "post-api/models"

type TaglineSaveRequest struct {
	UserID  string `json:"user_id" binding:"required" db:"user_id"`
	DraftID string `json:"draft_id" binding:"required" db:"draft_id"`
	Tagline string `json:"tagline" binding:"required" db:"tagline"`
}

type InterestsSaveRequest struct {
	UserID    string            `json:"user_id" binding:"required" db:"user_id" `
	DraftID   string            `json:"draft_id" binding:"required" db:"draft_id" `
	Interests models.JSONString `json:"interests" binding:"required" db:"interest" `
}

type PreviewImageSaveRequest struct {
	UserID          string `json:"user_id" binding:"required" db:"user_id"`
	DraftID         string `json:"draft_id" binding:"required" db:"draft_id"`
	PreviewImageUrl string `json:"preview_image" binding:"required" db:"preview_image"`
}

type DraftDeleteRequest struct {
	DraftID string `uri:"draft_id" binding:"required,validPostUID"`
}
