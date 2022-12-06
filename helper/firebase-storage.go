package helper

import (
	"context"
	"fmt"
	"io"
	"os"

	firebase "firebase.google.com/go"
	"github.com/thanhpk/randstr"
	"google.golang.org/api/option"
)

// UploadFileToFirebaseStorageAndGetURL uploads a file to Firebase Storage and returns its URL
func UploadFileToFirebaseStorageAndGetURL(ctx context.Context, file io.Reader) (string, error) {
	// Initialize the Firebase app
	bucket_name := "cbt-ppsm.appspot.com"
	config := &firebase.Config{
		ProjectID:     os.Getenv("FIREBASE_PROJECTID"),
		StorageBucket: os.Getenv("FIREBASE_STORAGEBUCKET"),
	}
	opt := option.WithCredentialsFile("/Users/user/project/backend-cbt/helper/cbt-ppsm-firebase-sa.json")
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return "", err
	}
	// Get a reference to the Firebase Storage bucket
	storageClient, err := app.Storage(ctx)
	if err != nil {
		return "", err
	}
	bucket, err := storageClient.Bucket(bucket_name)
	if err != nil {
		return "", err
	}
	// Upload the file to the specified path in the bucket
	file_name := randstr.Hex(16)
	obj := bucket.Object(fmt.Sprintf("%v", file_name))
	writer := obj.NewWriter(ctx)
	_, err = io.Copy(writer, file)
	if err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}
	// Get a signed URL for the uploaded file that will be valid for 1 hour
	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%v/o/%v?alt=media", string(bucket_name), string(file_name))
	return url, nil
}
