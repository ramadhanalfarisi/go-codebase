package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/ramadhanalfarisi/go-codebase/helpers"
	"github.com/ramadhanalfarisi/go-codebase/services/user/models"
	user_usecase "github.com/ramadhanalfarisi/go-codebase/services/user/usecase"
)

type UserController struct {
	usecase user_usecase.UserUsecaseInterface
}

func NewUserController(usecase user_usecase.UserUsecaseInterface) UserControllerInterface {
	return &UserController{
		usecase: usecase,
	}
}

func (u *UserController) UserLogin(c fiber.Ctx) error {
	var userLoginInput models.UserLoginInput
	err := json.Unmarshal(c.Body(), &userLoginInput)
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: "invalid request body",
		}
		return errResponse.SendResponse(c)
	}
	msgs, isValid := helpers.Validate(userLoginInput)
	if !isValid {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: msgs[0],
		}
		return errResponse.SendResponse(c)
	}
	jwt, err := u.usecase.UserLogin(userLoginInput)
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusUnauthorized,
			Status:  "error",
			Message: err.Error(),
		}
		return errResponse.SendResponse(c)
	}
	succesResponse := helpers.ResponseData{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Login successful",
		Data:    fiber.Map{"token": jwt},
	}
	return succesResponse.SendResponse(c)
}

func (u *UserController) UserRegister(c fiber.Ctx) error {
	var userRegisterInput models.UserRegisterInput
	err := json.Unmarshal(c.Body(), &userRegisterInput)
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: "invalid request body",
		}
		return errResponse.SendResponse(c)
	}
	msgs, isValid := helpers.Validate(userRegisterInput)
	if !isValid {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: msgs[0],
		}
		return errResponse.SendResponse(c)
	}
	err = u.usecase.UserRegister(userRegisterInput)
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: err.Error(),
		}
		return errResponse.SendResponse(c)
	}
	successResponse := helpers.ResponseData{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Register successful",
		Data:    nil,
	}
	return successResponse.SendResponse(c)
}
