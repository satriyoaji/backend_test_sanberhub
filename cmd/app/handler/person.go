package handler

import (
	"backend_test/model"
	"backend_test/pkg/util/responseutil"
	"backend_test/pkg/validator"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetPersons(ctx echo.Context) error {
	var filter model.GetPersonsFilter
	if err := validator.BindAndValidate(ctx, &filter); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	defaultPageRequest(&filter.PageRequest)
	results, total, err := h.personService.GetPersons(ctx, filter)
	if err.IsNoError() {
		pagination := model.Pagination{}
		pagination.PageNum = &filter.PageRequest.PageNum
		pagination.PageSize = &filter.PageRequest.PageSize
		pagination.TotalData = &total
		if results != nil {
			countShops := len(*results)
			pagination.PageSize = &countShops
		}
		return responseutil.SendSuccessReponse(ctx, results, &pagination)
	}
	return responseutil.SendErrorResponse(ctx, err)
}

func (h *Handler) AddPerson(ctx echo.Context) error {
	req := model.CreatePersonRequest{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	result, ce := h.personService.CreatePerson(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, result, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}

func (h *Handler) GetPersonByID(ctx echo.Context) error {
	req := model.GetPersonByIDRequest{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	result, ce := h.personService.GetPersonByID(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, result, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}

func (h *Handler) GetPersonCountryByName(ctx echo.Context) error {
	req := model.GetPersonCountryByNameRequest{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	result, ce := h.personService.GetPersonCountryByName(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, result, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}

func (h *Handler) EditPerson(ctx echo.Context) error {
	req := model.EditPersonRequest{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	result, ce := h.personService.EditPerson(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, result, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}
