package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type GetMutationByAccountNumber struct {
	Number string `param:"no_rekening" validate:"required"`
}

type GetMutationByAccountNumberResult struct {
	CreatedAt time.Time       `json:"waktu"`
	Number    string          `json:"no_rekening"`
	Code      string          `json:"kode_transaksi"`
	Amount    decimal.Decimal `json:"nominal"`
}
