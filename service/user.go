package service

import (
	"backend_test/constant"
	"backend_test/entity"
	"backend_test/model"
	"backend_test/pkg/util"
	"backend_test/repository"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"

	pkgerror "backend_test/pkg/error"
	"backend_test/pkg/util/copyutil"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserService interface {
	GetUsers(ctx echo.Context, filter model.GetUsersFilter) (*[]model.GetUsersResult, int, pkgerror.CustomError)
	CreateUser(ctx echo.Context, req model.CreateUserRequest) (*model.CreateUserResult, pkgerror.CustomError)
	GetUserByID(ctx echo.Context, req model.GetUserByIDRequest) (*model.GetUserByIDResult, pkgerror.CustomError)
	GetUserBalanceByNumber(ctx echo.Context, req model.GetUserBalanceByNumber) (*model.GetUserBalanceResult, pkgerror.CustomError)
	EditUser(ctx echo.Context, req model.EditUserRequest) (*model.EditUserResult, pkgerror.CustomError)
	SaveBalanceUser(ctx echo.Context, req model.SaveBalanceUserRequest) (*model.GetUserBalanceResult, pkgerror.CustomError)
	WithdrawalBalanceUser(ctx echo.Context, req model.WithdrawalBalanceUserRequest) (*model.GetUserBalanceResult, pkgerror.CustomError)
	GetUserMutationsByNumber(ctx echo.Context, req model.GetMutationByAccountNumber) ([]model.GetMutationByAccountNumberResult, pkgerror.CustomError)
}

type UserServiceImpl struct {
	repo repository.Repository
}

func NewUserService(
	repo repository.Repository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s UserServiceImpl) GetUsers(ctx echo.Context, filter model.GetUsersFilter) (*[]model.GetUsersResult, int, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	results := []model.GetUsersResult{}
	count, err := s.repo.CountUsers(rctx, filter)
	if err != nil {
		log.Error("Count user error: ", err)
		return nil, count, pkgerror.ErrSystemError
	}
	if count == 0 {
		return &results, count, pkgerror.NoError
	}
	users, err := s.repo.FindUsers(rctx, filter)
	if err != nil {
		log.Error("Find users error: ", err)
		return nil, count, pkgerror.ErrSystemError
	}
	copyutil.Copy(&users, &results)
	return &results, count, pkgerror.NoError
}

func (s *UserServiceImpl) GetUserByID(ctx echo.Context, req model.GetUserByIDRequest) (*model.GetUserByIDResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	user, err := s.repo.FindUserByID(rctx, uint(req.UserID))
	if err != nil {
		log.Error("Find user by ID error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrUserNotFound.WithError(err)
		}
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	result := model.GetUserByIDResult{}
	copyutil.Copy(&user, &result)
	return &result, pkgerror.NoError
}

func (s *UserServiceImpl) GetUserBalanceByNumber(ctx echo.Context, req model.GetUserBalanceByNumber) (*model.GetUserBalanceResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	userFound, err := s.repo.FindUserByColumnValue(rctx, string(constant.UserColumnNumber), req.Number)
	if err != nil {
		log.Error("Find user by number error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrUserNotFound.WithError(errors.New("Nasabah dengan `no_rekening` tersebut tidak dikenali"))
		}
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	result := model.GetUserBalanceResult{}
	copyutil.Copy(&userFound, &result)
	return &result, pkgerror.NoError
}

func (s *UserServiceImpl) CreateUser(ctx echo.Context, req model.CreateUserRequest) (*model.CreateUserResult, pkgerror.CustomError) {
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

	userFound, err := s.repo.FindUserByColumnValue(rctx, string(constant.UserColumnNIK), req.NIK)
	if err != nil {
		log.Error("Find user by NIK error: ", err)
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrSystemError.WithError(err)
		}
	}
	if userFound.Phone != "" {
		return nil, pkgerror.ErrUserRequestIsExist.WithError(errors.New("Data Nasabah sudah ada dengan `nik` tersebut."))
	}

	userFound, err = s.repo.FindUserByColumnValue(rctx, string(constant.UserColumnPhone), req.Phone)
	if err != nil {
		log.Error("Find user by Phone error: ", err)
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrSystemError.WithError(err)
		}
	}
	if userFound.Phone != "" {
		return nil, pkgerror.ErrUserRequestIsExist.WithError(errors.New("Data Nasabah sudah ada dengan `no_hp` tersebut."))
	}

	var user entity.User
	copyutil.Copy(&req, &user)
	user.Balance = decimal.NewFromInt32(int32(0))
	// generate unique account number
	user.Number = util.GenerateAccountNumber(9999999999, 1000000000)

	fmt.Println("User create: ", user)
	err = s.repo.CreateUser(rctx, &user)
	if err != nil {
		log.Error("Create user error: ", err)
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	err = s.repo.TxCommit()
	if err != nil {
		log.Error("Commit db transaction error: ", err)
	}
	var result model.CreateUserResult
	copyutil.Copy(&user, &result)
	txSuccess = true
	return &result, pkgerror.NoError
}

func (s *UserServiceImpl) EditUser(ctx echo.Context, req model.EditUserRequest) (*model.EditUserResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	user, err := s.repo.FindUserByID(rctx, uint(req.UserID))
	if err != nil {
		log.Error("Find user by ID error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrUserNotFound.WithError(err)
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
	err = s.repo.UpdateUser(rctx, &user)
	if err != nil {
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	// Commit transaction
	err = s.repo.TxCommit()
	if err != nil {
		log.Error("Commit db transaction error: ", err)
	}
	result := model.EditUserResult{}
	copyutil.Copy(&user, &result)
	txSuccess = true
	return &result, pkgerror.NoError
}

func (s *UserServiceImpl) SaveBalanceUser(ctx echo.Context, req model.SaveBalanceUserRequest) (*model.GetUserBalanceResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	user, err := s.repo.FindUserByColumnValue(rctx, string(constant.UserColumnNumber), req.Number)
	if err != nil {
		log.Error("Find user by number error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrUserNotFound.WithError(errors.New("Nasabah dengan `no_rekening` tersebut tidak dikenali"))
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

	sumBalance := user.Balance.InexactFloat64() + float64(req.Amount)
	user.Balance = decimal.NewFromFloat(sumBalance)
	err = s.repo.UpdateUser(rctx, &user)
	if err != nil {
		return nil, pkgerror.ErrSystemError.WithError(err)
	}

	// log to mutations
	mutation := entity.Mutation{
		Number: req.Number,
		Code:   string(constant.MutationCodeSave),
		Amount: decimal.NewFromInt(int64(req.Amount)),
	}
	err = s.repo.CreateMutation(rctx, &mutation)
	if err != nil {
		return nil, pkgerror.ErrSystemError.WithError(err)
	}

	// Commit transaction
	err = s.repo.TxCommit()
	if err != nil {
		log.Error("Commit db transaction error: ", err)
	}
	result := model.GetUserBalanceResult{}
	copyutil.Copy(&user, &result)
	txSuccess = true
	return &result, pkgerror.NoError
}

func (s *UserServiceImpl) WithdrawalBalanceUser(ctx echo.Context, req model.WithdrawalBalanceUserRequest) (*model.GetUserBalanceResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	user, err := s.repo.FindUserByColumnValue(rctx, string(constant.UserColumnNumber), req.Number)
	if err != nil {
		log.Error("Find user by number error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrUserNotFound.WithError(errors.New("Nasabah dengan `no_rekening` tersebut tidak dikenali"))
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

	// if balance not enough
	if user.Balance.LessThan(decimal.NewFromInt(int64(req.Amount))) {
		return nil, pkgerror.ErrUserBalanceNotEnough.WithError(errors.New("Saldo nasabah tidak mencukupi"))
	}

	sumBalance := user.Balance.InexactFloat64() - float64(req.Amount)
	user.Balance = decimal.NewFromFloat(sumBalance)
	err = s.repo.UpdateUser(rctx, &user)
	if err != nil {
		return nil, pkgerror.ErrSystemError.WithError(err)
	}

	// log to mutations
	mutation := entity.Mutation{
		Number: req.Number,
		Code:   string(constant.MutationCodeWithdrawal),
		Amount: decimal.NewFromInt(int64(req.Amount)),
	}
	err = s.repo.CreateMutation(rctx, &mutation)
	if err != nil {
		return nil, pkgerror.ErrSystemError.WithError(err)
	}

	// Commit transaction
	err = s.repo.TxCommit()
	if err != nil {
		log.Error("Commit db transaction error: ", err)
	}
	result := model.GetUserBalanceResult{}
	copyutil.Copy(&user, &result)
	txSuccess = true
	return &result, pkgerror.NoError
}

func (s *UserServiceImpl) GetUserMutationsByNumber(ctx echo.Context, req model.GetMutationByAccountNumber) ([]model.GetMutationByAccountNumberResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	_, err := s.repo.FindUserByColumnValue(rctx, string(constant.UserColumnNumber), req.Number)
	if err != nil {
		log.Error("Find user by number error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrUserNotFound.WithError(errors.New("Nasabah dengan `no_rekening` tersebut tidak dikenali"))
		}
		return nil, pkgerror.ErrSystemError.WithError(err)
	}

	mutations, err := s.repo.FindMutationsByNumber(rctx, req.Number)
	if err != nil {
		log.Error("Find user mutations by number error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrUserNotFound.WithError(errors.New("Mutasi dengan `no_rekening` tersebut belum ada data"))
		}
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	fmt.Println("mutations: ", mutations)

	var result []model.GetMutationByAccountNumberResult
	copyutil.Copy(&mutations, &result)
	return result, pkgerror.NoError
}
