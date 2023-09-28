package constant

type MimeType string

const (
	MimeTypeJpg MimeType = "image/jpeg"
	MimeTypePng MimeType = "image/png"
	MimeTypeGif MimeType = "image/gif"
)

var Images = []MimeType{
	MimeTypeJpg,
	MimeTypePng,
	MimeTypeGif,
}

func IsMimeTypeImage(str string) bool {
	for _, t := range Images {
		if str == string(t) {
			return true
		}
	}
	return false
}
