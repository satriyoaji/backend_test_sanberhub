package constant

import "errors"

type DiscountType string

const (
	DiscountTypeProductPrice DiscountType = "product_price"
	DiscountTypeTotalPrice   DiscountType = "total_price"
	DiscountTypeDeliveryFee  DiscountType = "delivery_fee"
)

var DiscountTypes = []DiscountType{
	DiscountTypeProductPrice,
	DiscountTypeTotalPrice,
	DiscountTypeDeliveryFee,
}

func ParseDiscountType(str string) (DiscountType, error) {
	for _, t := range DiscountTypes {
		if str == string(t) {
			return t, nil
		}
	}
	return "", errors.New(str)
}
