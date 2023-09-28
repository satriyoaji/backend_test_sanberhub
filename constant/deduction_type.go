package constant

import "errors"

type DeductionType string

const (
	DeductionTypePercentage DeductionType = "percentage"
	DeductionTypeFlatPrice  DeductionType = "flat_price"
)

var DeductionTypes = []DeductionType{
	DeductionTypePercentage,
	DeductionTypeFlatPrice,
}

func ParseDeductionType(str string) (DeductionType, error) {
	for _, t := range DeductionTypes {
		if str == string(t) {
			return t, nil
		}
	}
	return "", errors.New(str)
}
