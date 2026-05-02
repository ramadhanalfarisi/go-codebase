package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/ramadhanalfarisi/go-codebase/constants"
	"github.com/ramadhanalfarisi/go-codebase/helpers"
	"github.com/ramadhanalfarisi/go-codebase/services/product/models"
	"github.com/ramadhanalfarisi/go-codebase/services/product/usecase"
)

type ProductControllerAPI struct {
	usecase usecase.ProductUsecaseInterface
}

func NewProductControllerAPI(usecase usecase.ProductUsecaseInterface) ProductControllerAPIInterface {
	return &ProductControllerAPI{usecase: usecase}
}

// CreateProduct implements [ProductControllerAPIInterface].
func (p *ProductControllerAPI) CreateProduct(c fiber.Ctx) error {
	var productInput models.CreateProductInput
	err := json.Unmarshal(c.Body(), &productInput)
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  constants.StatusError,
			Message: constants.InvalidRequestBody,
		}
		return errResponse.SendResponse(c)
	}
	msgs, isValid := helpers.Validate(productInput)
	if !isValid {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  constants.StatusError,
			Message: msgs[0],
		}
		return errResponse.SendResponse(c)
	}
	product, err := p.usecase.CreateProduct(productInput)
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusInternalServerError,
			Status:  constants.StatusError,
			Message: err.Error(),
		}
		return errResponse.SendResponse(c)
	}
	succesResponse := helpers.Response{
		Code:    fiber.StatusOK,
		Status:  constants.StatusSuccess,
		Message: "Product created successfully",
		Data:    product,
	}
	return succesResponse.SendResponse(c)
}

// DeleteProduct implements [ProductControllerAPIInterface].
func (p *ProductControllerAPI) DeleteProduct(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  constants.StatusError,
			Message: constants.IdIsRequired,
		}
		return errResponse.SendResponse(c)
	}
	idInt := helpers.StringToInt(id)
	prod, err := p.usecase.DeleteProduct(idInt)
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusInternalServerError,
			Status:  constants.StatusError,
			Message: err.Error(),
		}
		return errResponse.SendResponse(c)
	}
	succesResponse := helpers.Response{
		Code:    fiber.StatusOK,
		Status:  constants.StatusSuccess,
		Message: "Product deleted successfully",
		Data:    prod,
	}
	return succesResponse.SendResponse(c)
}

// GetProductById implements [ProductControllerAPIInterface].
func (p *ProductControllerAPI) GetProductById(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  constants.StatusError,
			Message: constants.IdIsRequired,
		}
		return errResponse.SendResponse(c)
	}
	idInt := helpers.StringToInt(id)
	prod, err := p.usecase.GetProductById(idInt)
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusInternalServerError,
			Status:  constants.StatusError,
			Message: err.Error(),
		}
		return errResponse.SendResponse(c)
	}
	succesResponse := helpers.Response{
		Code:    fiber.StatusOK,
		Status:  constants.StatusSuccess,
		Message: "Product retrieved successfully",
		Data:    prod,
	}
	return succesResponse.SendResponse(c)
}

// GetProducts implements [ProductControllerAPIInterface].
func (p *ProductControllerAPI) GetProducts(c fiber.Ctx) error {
	prods, err := p.usecase.GetProducts()
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusInternalServerError,
			Status:  constants.StatusError,
			Message: err.Error(),
		}
		return errResponse.SendResponse(c)
	}
	succesResponse := helpers.Response{
		Code:    fiber.StatusOK,
		Status:  constants.StatusSuccess,
		Message: "Products retrieved successfully",
		Data:    prods,
	}
	return succesResponse.SendResponse(c)
}

// UpdateProduct implements [ProductControllerAPIInterface].
func (p *ProductControllerAPI) UpdatePatchProduct(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  constants.StatusError,
			Message: constants.IdIsRequired,
		}
		return errResponse.SendResponse(c)
	}
	idInt := helpers.StringToInt(id)

	var productInput models.PatchProductInput
	err := json.Unmarshal(c.Body(), &productInput)
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  constants.StatusError,
			Message: constants.InvalidRequestBody,
		}
		return errResponse.SendResponse(c)
	}
	msgs, isValid := helpers.Validate(productInput)
	if !isValid {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  constants.StatusError,
			Message: msgs[0],
		}
		return errResponse.SendResponse(c)
	}
	prod, err := p.usecase.UpdateProduct(idInt, productInput)
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusInternalServerError,
			Status:  constants.StatusError,
			Message: err.Error(),
		}
		return errResponse.SendResponse(c)
	}
	succesResponse := helpers.Response{
		Code:    fiber.StatusOK,
		Status:  constants.StatusSuccess,
		Message: "Product updated successfully",
		Data:    prod,
	}
	return succesResponse.SendResponse(c)
}

// UpdatePutProduct implements [ProductControllerAPIInterface].
func (p *ProductControllerAPI) UpdatePutProduct(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  constants.StatusError,
			Message: constants.IdIsRequired,
		}
		return errResponse.SendResponse(c)
	}
	idInt := helpers.StringToInt(id)

	var productInput models.PutProductInput
	err := json.Unmarshal(c.Body(), &productInput)
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  constants.StatusError,
			Message: constants.InvalidRequestBody,
		}
		return errResponse.SendResponse(c)
	}
	msgs, isValid := helpers.Validate(productInput)
	if !isValid {
		errResponse := helpers.Response{
			Code:    fiber.StatusBadRequest,
			Status:  constants.StatusError,
			Message: msgs[0],
		}
		return errResponse.SendResponse(c)
	}
	prod, err := p.usecase.UpdatePutProduct(idInt, productInput)
	if err != nil {
		errResponse := helpers.Response{
			Code:    fiber.StatusInternalServerError,
			Status:  constants.StatusError,
			Message: err.Error(),
		}
		return errResponse.SendResponse(c)
	}
	succesResponse := helpers.Response{
		Code:    fiber.StatusOK,
		Status:  constants.StatusSuccess,
		Message: "Product updated successfully",
		Data:    prod,
	}
	return succesResponse.SendResponse(c)
}
