package helper

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func ImageUploader(ctx context.Context, cld *cloudinary.Cloudinary, file any, folder string) (string, error) {
	res, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: fmt.Sprintf("cbt/%s", folder),
	})
	if err != nil {
		return "", err
	}
	return res.SecureURL, nil
}
