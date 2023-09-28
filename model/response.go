package model

type ResponseBody struct {
	Status       string      `json:"status"`
	Code         string      `json:"code"`
	Data         interface{} `json:"data"`
	Pagination   *Pagination `json:"pagination"`
	ErrorMessage *string     `json:"error_message"`
	ErrorDebug   *string     `json:"error_debug,omitempty"`
}
