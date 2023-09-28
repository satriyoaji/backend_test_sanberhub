package constant

import "errors"

type MutationCode string

const (
	MutationCodeSave       MutationCode = "C"
	MutationCodeWithdrawal MutationCode = "D"
)

var MutationCodes = []MutationCode{
	MutationCodeSave,
	MutationCodeWithdrawal,
}

func ParseMutationCodeName(str string) (MutationCode, error) {
	for _, t := range MutationCodes {
		if str == string(t) {
			return t, nil
		}
	}
	return "", errors.New(str)
}
