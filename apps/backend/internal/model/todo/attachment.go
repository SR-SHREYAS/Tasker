package todo

import (
	"github.com/SR-SHREYAS/Tasker/internal/model"
	"github.com/google/uuid"
)

type TodoAttachment struct {
	model.Base
	TodoID      uuid.UUID `json:"todoid" db:"todo_id"`
	Name        string    `json:"name" db:"name"`
	UploadedBy  string    `json:"uploadedby" db:"uploaded_by"`
	DownloadKey string    `json:"downloadkey" db:"download_key"`
	FileSize    *int64    `json:"filesize" db:"file_size"`
	MimeType    *string   `json:"mimetype" db:"mime_type"`
}
