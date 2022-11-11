package config

import "os"

var CloudinaryCloudName = os.Getenv("CLOUDINARY_CLOUD_NAME")
var CloudinaryApiKey = os.Getenv("CLOUDINARY_API_KEY")
var CloudinaryApiSecreet = os.Getenv("CLOUDINARY_API_SECRET")
var CloudinaryUploadFolder = os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
