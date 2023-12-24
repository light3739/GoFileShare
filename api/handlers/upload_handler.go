package handlers

import (
	"mime/multipart"
	"net/http"
	"share-files/logger"
	"share-files/pkg/storage"
)

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

	tag := storage.GenerateTag()

	if err := storage.AddTag(sess, header.Filename, tag); err != nil {
		logger.LogError(err, "Failed to add tag to S3 object")
		http.Error(w, "Failed to add tag to S3 object", http.StatusInternalServerError)
		return
	}
	logger.LogInfo("Successfully uploaded to S3 bucket. Tag: " + tag)
	w.WriteHeader(http.StatusOK)
}
