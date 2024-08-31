package utils

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(filePath string) string {
	cld, err := cloudinary.NewFromURL("cloudinary://456751956474822:VZ8aBBh3vA1Mo19fI6fpeDUCA0A@derf8sbin")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := cld.Upload.Upload(context.Background(), filePath, uploader.UploadParams{})
	if err != nil {
		log.Fatal(err)
	}

	return resp.SecureURL
}
