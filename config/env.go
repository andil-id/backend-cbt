package config

import "os"

var GinMode = os.Getenv("GIN_MODE")
var PathLog = os.Getenv("PATH_LOG")
var CloudinaryCloudName = os.Getenv("CLOUDINARY_CLOUD_NAME")
var CloudinaryApiKey = os.Getenv("CLOUDINARY_API_KEY")
var CloudinaryApiSecreet = os.Getenv("CLOUDINARY_API_SECRET")
var CloudinaryUploadFolder = os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
