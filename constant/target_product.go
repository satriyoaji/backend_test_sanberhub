package constant

import "errors"

type TargetProduct string

const (
	TargetProductAll      TargetProduct = "all"
	TargetProductTargeted TargetProduct = "targeted"
)

var TargetProducts = []TargetProduct{
	TargetProductAll,
	TargetProductTargeted,
}

func ParseTargetProduct(str string) (TargetProduct, error) {
	for _, t := range TargetProducts {
		if str == string(t) {
			return t, nil
		}
	}
	return "", errors.New(str)
}
