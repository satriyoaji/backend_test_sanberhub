package handler

import (
	"backend_test/model"
	"backend_test/pkg/util/responseutil"
	"backend_test/pkg/validator"
	"github.com/labstack/echo/v4"
)

func (h *Handler) AddUser(ctx echo.Context) error {
	req := model.CreateUserRequest{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	result, ce := h.userService.CreateUser(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, result, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}

func (h *Handler) GetUserBalanceByNumber(ctx echo.Context) error {
	req := model.GetUserBalanceByNumber{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	result, ce := h.userService.GetUserBalanceByNumber(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, result, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}

func (h *Handler) SaveBalanceUser(ctx echo.Context) error {
	req := model.SaveBalanceUserRequest{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	result, ce := h.userService.SaveBalanceUser(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, result, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}
