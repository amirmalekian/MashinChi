package handler

import (
	"net/http"
	"strconv"

	"github.com/amirmalekian/Mashinchi/src/api/dto"
	"github.com/amirmalekian/Mashinchi/src/api/helper"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/amirmalekian/Mashinchi/src/services"
	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	return &CountryHandler{
		service: services.NewCountryService(cfg),
	}
}

// CreateCountry godoc
// @Summary Create a country
// @Description Create a country
// @Tags Countries
// @Accept json
// @produces json
// @Param Request body dto.CreateUpdateCountryRequest true "Create a country"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/ [post]
// @Security AuthBearer
func (h *CountryHandler) Create(ctx *gin.Context) {
	req := dto.CreateUpdateCountryRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, 121, err))
		return
	}

	res, err := h.service.Create(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

// UpdateCountry godoc
// @Summary Update a country
// @Description Update a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdateCountryRequest true "Update a country"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [put]
// @Security AuthBearer
func (h *CountryHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	req := dto.CreateUpdateCountryRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, 121, err))
		return
	}

	res, err := h.service.Update(ctx, id, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

// DeleteCountry godoc
// @Summary Delete a country
// @Description Delete a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [delete]
// @Security AuthBearer
func (h *CountryHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, 121))
		return
	}

	err := h.service.Delete(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, 0))
}

// GetCountry godoc
// @Summary Get a country
// @Description Get a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [get]
// @Security AuthBearer
func (h *CountryHandler) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, 121))
		return
	}

	res, err := h.service.GetById(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

// GetCountries godoc
// @Summary Get Countries
// @Description Get Countries
// @Tags Countries
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CountryResponse]} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/get-by-filter [post]
// @Security AuthBearer
func (h *CountryHandler) GetByFilter(ctx *gin.Context) {
req := dto.PaginationInputWithFilter{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, 121, err))
		return
	}

	res, err := h.service.GetByFilter(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}
