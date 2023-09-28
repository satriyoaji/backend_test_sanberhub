package handler

import (
	"backend_test/model"
	"backend_test/pkg/util/responseutil"
	"backend_test/pkg/validator"
	"github.com/labstack/echo/v4"
	"net/http"
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

func (h *Handler) WithdrawalBalanceUser(ctx echo.Context) error {
	req := model.WithdrawalBalanceUserRequest{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	result, ce := h.userService.WithdrawalBalanceUser(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, result, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}

func (h *Handler) GetMutationUserByNumber(ctx echo.Context) error {
	req := model.GetMutationByAccountNumber{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	result, ce := h.userService.GetUserMutationsByNumber(ctx, req)
	if ce.IsNoError() {
		response := model.ResponseBodyMutation{
			Status: "SUCCESS",
			Code:   "0000",
			Data:   result,
		}
		return ctx.JSON(http.StatusOK, response)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}
