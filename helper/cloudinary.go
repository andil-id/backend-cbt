package helper

import (
	"context"
	"fmt"

	"github.com/andil-id/api/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func ImageUploader(ctx context.Context, cld *cloudinary.Cloudinary, file any, folder string) (string, error) {
	baseFolder := config.CloudinaryUploadFolder()
	res, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: fmt.Sprintf("%s/%s", baseFolder, folder),
	})
	if err != nil {
		return "", err
	}
	return res.SecureURL, nil
}
