package service

import (
	"backend_test/entity"
	"backend_test/model"
	"backend_test/repository"
	"errors"

	pkgerror "backend_test/pkg/error"
	"backend_test/pkg/util/copyutil"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PersonService interface {
	GetPersons(ctx echo.Context, filter model.GetPersonsFilter) (*[]model.GetPersonsResult, int, pkgerror.CustomError)
	CreatePerson(ctx echo.Context, req model.CreatePersonRequest) (*model.CreatePersonResult, pkgerror.CustomError)
	GetPersonByID(ctx echo.Context, req model.GetPersonByIDRequest) (*model.GetPersonByIDResult, pkgerror.CustomError)
	GetPersonCountryByName(ctx echo.Context, req model.GetPersonCountryByNameRequest) (*string, pkgerror.CustomError)
	EditPerson(ctx echo.Context, req model.EditPersonRequest) (*model.EditPersonResult, pkgerror.CustomError)
}

type PersonServiceImpl struct {
	repo repository.Repository
}

func NewPersonService(
	repo repository.Repository) *PersonServiceImpl {
	return &PersonServiceImpl{
		repo: repo,
	}
}

func (s PersonServiceImpl) GetPersons(ctx echo.Context, filter model.GetPersonsFilter) (*[]model.GetPersonsResult, int, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	results := []model.GetPersonsResult{}
	count, err := s.repo.CountPersons(rctx, filter)
	if err != nil {
		log.Error("Count person error: ", err)
		return nil, count, pkgerror.ErrSystemError
	}
	if count == 0 {
		return &results, count, pkgerror.NoError
	}
	persons, err := s.repo.FindPersons(rctx, filter)
	if err != nil {
		log.Error("Find persons error: ", err)
		return nil, count, pkgerror.ErrSystemError
	}
	copyutil.Copy(&persons, &results)
	return &results, count, pkgerror.NoError
}

func (s *PersonServiceImpl) GetPersonByID(ctx echo.Context, req model.GetPersonByIDRequest) (*model.GetPersonByIDResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	person, err := s.repo.FindPersonByID(rctx, uint(req.PersonID))
	if err != nil {
		log.Error("Find person by ID error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrPersonNotFound.WithError(err)
		}
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	result := model.GetPersonByIDResult{}
	copyutil.Copy(&person, &result)
	return &result, pkgerror.NoError
}

func (s *PersonServiceImpl) GetPersonCountryByName(ctx echo.Context, req model.GetPersonCountryByNameRequest) (*string, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	person, err := s.repo.FindPersonByName(rctx, req.Name)
	result := ""
	if err != nil {
		log.Error("Find person by ID error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return &result, pkgerror.ErrPersonNotFound.WithError(err)
		}
		return &result, pkgerror.ErrSystemError.WithError(err)
	}

	return &person.Country, pkgerror.NoError
}

func (s *PersonServiceImpl) CreatePerson(ctx echo.Context, req model.CreatePersonRequest) (*model.CreatePersonResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	txSuccess := false
	err := s.repo.TxBegin()
	if err != nil {
		log.Error("Start db transaction error: ", err)
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	defer func() {
		if r := recover(); r != nil || !txSuccess {
			err = s.repo.TxRollback()
			if err != nil {
				log.Error("Rollback db transaction error: ", err)
			}
		}
	}()
	var person entity.Person
	copyutil.Copy(&req, &person)
	err = s.repo.CreatePerson(rctx, &person)
	if err != nil {
		log.Error("Create person error: ", err)
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	err = s.repo.TxCommit()
	if err != nil {
		log.Error("Commit db transaction error: ", err)
	}
	var result model.CreatePersonResult
	copyutil.Copy(&person, &result)
	txSuccess = true
	return &result, pkgerror.NoError
}

func (s *PersonServiceImpl) EditPerson(ctx echo.Context, req model.EditPersonRequest) (*model.EditPersonResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	person, err := s.repo.FindPersonByID(rctx, uint(req.PersonID))
	if err != nil {
		log.Error("Find person by ID error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrPersonNotFound.WithError(err)
		}
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	// Start transaction
	txSuccess := false
	err = s.repo.TxBegin()
	if err != nil {
		log.Error("Start db transaction error: ", err)
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	defer func() {
		if r := recover(); r != nil || !txSuccess {
			err = s.repo.TxRollback()
			if err != nil {
				log.Error("Rollback db transaction error: ", err)
			}
		}
	}()
	err = s.repo.UpdatePerson(rctx, &person)
	if err != nil {
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	// Commit transaction
	err = s.repo.TxCommit()
	if err != nil {
		log.Error("Commit db transaction error: ", err)
	}
	result := model.EditPersonResult{}
	copyutil.Copy(&person, &result)
	txSuccess = true
	return &result, pkgerror.NoError
}
