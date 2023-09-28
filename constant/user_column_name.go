package constant

import "errors"

type UserColumn string

const (
	UserColumnName    UserColumn = "name"
	UserColumnNumber  UserColumn = "number"
	UserColumnNIK     UserColumn = "nik"
	UserColumnPhone   UserColumn = "phone"
	UserColumnBalance UserColumn = "balance"
)

var UserColumns = []UserColumn{
	UserColumnName,
	UserColumnNumber,
	UserColumnNIK,
	UserColumnPhone,
	UserColumnBalance,
}

func ParseUserColumnName(str string) (UserColumn, error) {
	for _, t := range UserColumns {
		if str == string(t) {
			return t, nil
		}
	}
	return "", errors.New(str)
}
