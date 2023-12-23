package handlers

// upload_handler.go
import (
	"mime/multipart"
	"net/http"
	"share-files/logger"
	"share-files/pkg/storage"
)

// HandleUpload handles POST requests to upload files.
// @Summary Uploads a file
// @Description Uploads a file to S3
// @ID upload-file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {string} string "File uploaded successfully"
// @Router /upload [post]
func HandleUpload(w http.ResponseWriter, r *http.Request) {
	sess, err := storage.NewAWSSession()
	if err != nil {
		logger.LogError(err, "Failed to create AWS session")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	file, header, err := storage.GetFileFromRequest(r)
	if err != nil {
		logger.LogError(err, "Failed to retrieve file")
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			logger.LogError(err, "Failed to close file")
			http.Error(w, "Internal Server Error", http.StatusBadRequest)
		}
	}(file)

	if err := storage.UploadFile(sess, file, header); err != nil {
		logger.LogError(err, "Failed to upload file to S3")
		http.Error(w, "Failed to upload file to S3", http.StatusInternalServerError)
		return
	}

	logger.LogInfo("Successfully uploaded to S3 bucket")
	w.WriteHeader(http.StatusOK)
}
