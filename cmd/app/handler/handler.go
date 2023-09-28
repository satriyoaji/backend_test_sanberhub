package handler

import (
	"backend_test/model"
	"backend_test/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	personService service.PersonService
	userService   service.UserService
}

func NewHandler(
	personService service.PersonService,
	userService service.UserService,
) *Handler {
	return &Handler{
		userService:   userService,
		personService: personService,
	}
}

func defaultPageRequest(pr *model.PageRequest) {
	if pr.PageNum == 0 {
		pr.PageNum = 1
	}
	if pr.PageSize == 0 {
		pr.PageSize = 10
	}
}

func RegisterHandlers(e *echo.Echo, h *Handler) {
	e.POST("/daftar", h.AddUser)
	e.GET("/saldo/:no_rekening", h.GetUserBalanceByNumber)
	e.POST("/tabung", h.SaveBalanceUser)
	e.POST("/tarik", h.WithdrawalBalanceUser)
	e.GET("/mutasi/:no_rekening", h.GetMutationUserByNumber)

}
