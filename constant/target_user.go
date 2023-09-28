package constant

import "errors"

type TargetUser string

const (
	TargetUserAll      TargetUser = "all"
	TargetUserTargeted TargetUser = "targeted"
)

var TargetUsers = []TargetUser{
	TargetUserAll,
	TargetUserTargeted,
}

func ParseTargetUser(str string) (TargetUser, error) {
	for _, t := range TargetUsers {
		if str == string(t) {
			return t, nil
		}
	}
	return "", errors.New(str)
}
