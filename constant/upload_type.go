package constant

import "errors"

type UploadType string

const (
	UploadTypeMerchantLogo      UploadType = "merchant-logo"
	UploadTypeProductImage      UploadType = "product-image"
	UploadTypeNotificationImage UploadType = "notification-image"
	UploadTypePromotionImage    UploadType = "promotion-image"
)

var UploadTypes = []UploadType{
	UploadTypeMerchantLogo,
	UploadTypeProductImage,
	UploadTypeNotificationImage,
	UploadTypePromotionImage,
}

func ParseUploadType(str string) (UploadType, error) {
	for _, t := range UploadTypes {
		if str == string(t) {
			return t, nil
		}
	}
	return "", errors.New(str)
}
