package types

import "fmt"

type EnumPhotosType string

const (
	EnumPhotosTypeImageGif   EnumPhotosType = "image/gif"
	EnumPhotosTypeImageXIcon EnumPhotosType = "image/x-icon"
	EnumPhotosTypeImageJpeg  EnumPhotosType = "image/jpeg"
	EnumPhotosTypeImagePng   EnumPhotosType = "image/png"
	EnumPhotosTypeImageTiff  EnumPhotosType = "image/tiff"
)

func (e *EnumPhotosType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumPhotosType(s)
	case string:
		*e = EnumPhotosType(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumPhotosType: %T", src)
	}
	return nil
}
