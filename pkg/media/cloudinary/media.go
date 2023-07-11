package media

import (
	"context"
	"e-course/pkg/resp"
	"e-course/pkg/utils"
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Media interface {
	Upload(file multipart.FileHeader) (*string, *resp.ErrorResp)
	Delete(fn string) (*string, *resp.ErrorResp)
}

type mediaUsecase struct {
	url string
}

// Delete implements Media.
func (u *mediaUsecase) Delete(fn string) (*string, *resp.ErrorResp) {
	cld, err := cloudinary.NewFromURL(u.url)
	if err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	var ctx = context.Background()
	filename := utils.GetFileName(fn)
	res, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: filename,
	})

	if err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return &res.Result, nil
}

// Upload implements Media.
func (u *mediaUsecase) Upload(file multipart.FileHeader) (*string, *resp.ErrorResp) {
	cld, err := cloudinary.NewFromURL(u.url)
	if err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	var ctx = context.Background()

	binaryFile, err := file.Open()
	if err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	defer func() {
		binaryFile.Close()
	}()

	if binaryFile != nil {
		pid, err := utils.GenerateRefreshToken()
		if err != nil {
			return nil, &resp.ErrorResp{
				Code: 500,
				Err:  err,
			}
		}

		uploadRes, err := cld.Upload.Upload(
			ctx,
			binaryFile,
			uploader.UploadParams{
				PublicID: pid.String(),
			},
		)
		if err != nil {
			return nil, &resp.ErrorResp{
				Code: 500,
				Err:  err,
			}
		}

		return &uploadRes.SecureURL, nil
	}
	return nil, &resp.ErrorResp{
		Code: 500,
		Err:  errors.New("cannot read file"),
	}
}

func NewMediaUsecase() Media {
	url := fmt.Sprintf(
		"cloudinary://%s:%s@%s",
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
	)
	return &mediaUsecase{
		url: url,
	}
}
