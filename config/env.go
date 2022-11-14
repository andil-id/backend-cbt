package config

import "os"

func PathLog() string {
	return os.Getenv("PATH_LOG")
}

func GinMode() string {
	return os.Getenv("GIN_MODE")
}

func CloudinaryCloudName() string {
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func CloudinaryApiKey() string {
	return os.Getenv("CLOUDINARY_API_KEY")
}

func CloudinaryApiSecret() string {
	return os.Getenv("CLOUDINARY_API_SECRET")
}

func CloudinaryUploadFolder() string {
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}
