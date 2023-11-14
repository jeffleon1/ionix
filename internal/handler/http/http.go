package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jeffleon1/ionix/internal/controller/drug"
	"github.com/jeffleon1/ionix/internal/controller/user"
	"github.com/jeffleon1/ionix/internal/controller/vaccination"
	"github.com/jeffleon1/ionix/pkg/entities"
	"github.com/jeffleon1/ionix/pkg/models"
	"github.com/jeffleon1/ionix/pkg/utils"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Msg    string      `json:"msg"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Err    interface{} `json:"error"`
}

type controllers struct {
	drug        drug.Controller
	vaccination vaccination.Controller
	user        user.Controller
}

type Handler struct {
	ctrls controllers
}

func New(drug drug.Controller, vaccination vaccination.Controller, user user.Controller) *Handler {
	return &Handler{
		ctrls: controllers{
			drug,
			vaccination,
			user,
		},
	}
}

// SignUp
//
//		@Tags			User
//		@Summary		User Signup
//		@Description	User Signup
//		@Accept			json
//		@Produce		json
//	    @Param request body entities.User true "query params"
//		@Success		200		{object}	models.BaseResponseModel
//		@Failure		400		{object}	models.BaseResponseModel
//		@Failure		422		{object}	models.BaseResponseModel
//	@Router			/signup [post]
func (h *Handler) SignUp(ctx *gin.Context) {
	var requestBody entities.User
	if err := ctx.BindJSON(&requestBody); err != nil {
		logrus.Errorf("Signup error making binding error: %s", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, utils.StatusUnprocessableEntity(err.Error()))
		return
	}
	requestBody.Email = strings.ToLower(requestBody.Email)

	validate := validator.New()

	// Validate the User struct
	if err := validate.Struct(requestBody); err != nil {
		errors := err.(validator.ValidationErrors)
		logrus.Errorf("Signup error: %s", errors.Error())
		ctx.JSON(http.StatusBadRequest, utils.StatusFail(errors.Error()))
		return
	}

	token, err := h.ctrls.user.Signup(ctx.Request.Context(), requestBody)
	if err != nil {
		logrus.Errorf("Signup error: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	logrus.Infof("User Created Successfully: %d", requestBody.ID)
	ctx.JSON(http.StatusOK, utils.StatusOK(token))
}

// SignIn
//
//		@Tags			User
//		@Summary		User Signin
//		@Description	User Signin
//		@Accept			json
//	 @Param request body models.LoginModel true "query params"
//		@Success		200		{object}	models.BaseResponseModel
//		@Failure		400		{object}	models.BaseResponseModel
//		@Failure		422		{object}	models.BaseResponseModel
//		@Router			/signin [post]
func (h *Handler) SignIn(ctx *gin.Context) {
	var requestBody models.LoginModel
	if err := ctx.BindJSON(&requestBody); err != nil {
		logrus.Errorf("Signin error making binding error: %s", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, utils.StatusUnprocessableEntity(err.Error()))
		return
	}
	requestBody.Email = strings.ToLower(requestBody.Email)
	token, err := h.ctrls.user.Login(ctx.Request.Context(), requestBody.Email, requestBody.Password)
	if err != nil {
		logrus.Errorf("Signin error: %s", err.Error())
		ctx.JSON(http.StatusUnauthorized, utils.StatusUnauthorized(err.Error()))
		return
	}

	logrus.Infof("User Login Successfully: %s", requestBody.Email)
	ctx.JSON(http.StatusOK, utils.StatusOK(token))
}

// GetDrug
//
//		@Tags			Drugs
//		@Summary		Get drug by id
//		@Description	Get drug by id
//		@Accept			json
//		@Produce		json
//	    @Param          id   path      string  true  "drug ID"
//		@Param			Authorization	header		string	true	"access token"	extensions(x-example=token)
//		@Success		200				{object}	models.BaseResponseModel
//		@Failure		400				{object}	models.BaseResponseModel
//		@Router			/drugs/{id} [get]
func (h *Handler) GetDrug(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.Errorf("Error parsing param id in Get drug: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	uintID := uint(u64)

	entity, err := h.ctrls.drug.Get(ctx, uintID)
	if err != nil {
		logrus.Errorf("Error getting drug with id %d: %s", uintID, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.StatusOK(entity))
}

// GetVaccination
//
//	@Tags			Vaccination
//	@Summary		Get vaccination by id
//	@Description	Get vaccination by id
//	@Accept			json
//	@Produce		json
//	@Param          id   path      string  true  "vaccination ID"
//	@Param			Authorization	header		string	true	"access token"	extensions(x-example=token)
//	@Success		200				{object}	models.BaseResponseModel
//	@Failure		400				{object}	models.BaseResponseModel
//	@Router			/vaccination/{id} [get]
func (h *Handler) GetVaccination(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.Errorf("Error getting param id in method get vaccination: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	uintID := uint(u64)

	entity, err := h.ctrls.vaccination.Get(ctx, uintID)
	if err != nil {
		logrus.Errorf("Error getting vaccination with id %d: %s", uintID, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.StatusOK(entity))
}

// GetDrugs
//
//	@Tags			Drugs
//	@Summary		Get drugs
//	@Description	Get drugs
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"access token"	extensions(x-example=token)
//	@Success		200				{object}	models.BaseResponseModel
//	@Failure		400				{object}	models.BaseResponseModel
//	@Router			/drugs [get]
func (h *Handler) GetDrugs(c *gin.Context) {
	ctx := c.Request.Context()
	entities, err := h.ctrls.drug.GetAll(ctx)
	if err != nil {
		logrus.Errorf("Error getting drugs: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.StatusOK(entities))
}

// GetVaccinations
//
//	@Tags			Vaccination
//	@Summary		GetVaccinations
//	@Description	GetVaccinations
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"access token"	extensions(x-example=token)
//	@Success		200				{object}	models.BaseResponseModel
//	@Failure		400				{object}	models.BaseResponseModel
//	@Router			/vaccination [get]
func (h *Handler) GetVaccinations(c *gin.Context) {
	ctx := c.Request.Context()
	entities, err := h.ctrls.vaccination.GetAll(ctx)
	if err != nil {
		logrus.Errorf("Error getting vaccinations: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.StatusOK(entities))
}

// CreateVaccination
//
//		@Tags			Vaccination
//		@Summary		create vaccination
//		@Description	create vaccination
//		@Accept			json
//	@Param			Authorization	header		string	true	"access token"	extensions(x-example=token)
//	 @Param request body entities.Vaccination true "query params"
//		@Success		200		{object}	models.BaseResponseModel
//		@Failure		400		{object}	models.BaseResponseModel
//		@Failure		422		{object}	models.BaseResponseModel
//		@Router			/vaccination [post]
func (h *Handler) CreateVaccination(c *gin.Context) {
	ctx := c.Request.Context()
	var requestBody entities.Vaccination

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logrus.Errorf("Error making binding in method creating vaccination: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, utils.StatusUnprocessableEntity(err.Error()))
		return
	}

	validate := validator.New()

	// Validate the vaccination struct
	if err := validate.Struct(requestBody); err != nil {
		errors := err.(validator.ValidationErrors)
		logrus.Errorf("Error validating vaccination: %s", errors.Error())
		c.JSON(http.StatusBadRequest, utils.StatusFail(errors.Error()))
		return
	}

	err := h.ctrls.vaccination.Create(ctx, &requestBody)
	if err != nil {
		logrus.Errorf("Error creating vaccination: %s", err.Error())
		c.JSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.StatusOK(nil))
}

// CreateDrug
//
//		@Tags			Drugs
//		@Summary		create drug
//		@Description	create drug
//		@Accept			json
//	@Param			Authorization	header		string	true	"access token"	extensions(x-example=token)
//	 @Param request body entities.Drug true "query params"
//		@Success		200		{object}	models.BaseResponseModel
//		@Failure		400		{object}	models.BaseResponseModel
//		@Failure		422		{object}	models.BaseResponseModel
//		@Router			/drugs [post]
func (h *Handler) CreateDrug(c *gin.Context) {
	ctx := c.Request.Context()
	var requestBody entities.Drug

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logrus.Errorf("Error making binding in method creating drug: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, utils.StatusUnprocessableEntity(err.Error()))
		return
	}

	validate := validator.New()

	// Validate the drug struct
	if err := validate.Struct(requestBody); err != nil {
		errors := err.(validator.ValidationErrors)
		logrus.Errorf("Error creating vaccination: %s", errors.Error())
		c.JSON(http.StatusBadRequest, utils.StatusFail(errors.Error()))
		return
	}

	err := h.ctrls.drug.Create(ctx, &requestBody)
	if err != nil {
		logrus.Errorf("Error creating drug: %s", err.Error())
		c.JSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.StatusOK(nil))
}

// UpdateVaccination
//
//		@Tags			Vaccination
//		@Summary		update vaccination
//		@Description	update vaccination
//		@Accept			json
//	@Param			Authorization	header		string	true	"access token"	extensions(x-example=token)
//	 @Param request body entities.Vaccination true "query params"
//		@Success		200		{object}	models.BaseResponseModel
//		@Failure		400		{object}	models.BaseResponseModel
//		@Failure		422		{object}	models.BaseResponseModel
//		@Router			/vaccination [put]
func (h *Handler) UpdateVaccination(c *gin.Context) {
	ctx := c.Request.Context()
	var requestBody entities.Vaccination

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logrus.Errorf("Error making binding JSON in method updating vaccination: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	id := c.Param("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.Errorf("Error parsing param id in method update vaccination: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	uintID := uint(u64)

	if err = h.ctrls.vaccination.Update(ctx, &requestBody, uintID); err != nil {
		logrus.Errorf("Error updating vaccination: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.StatusOK(nil))
}

// UpdateDrug
//
//		@Tags			Drugs
//		@Summary		update drug
//		@Description	update drug
//		@Accept			json
//	    @Param			Authorization	header		string	true	"access token"	extensions(x-example=token)
//	    @Param request body entities.Drug true "query params"
//		@Success		200		{object}	models.BaseResponseModel
//		@Failure		400		{object}	models.BaseResponseModel
//		@Failure		422		{object}	models.BaseResponseModel
//		@Router			/drugs [put]
func (h *Handler) UpdateDrug(c *gin.Context) {
	ctx := c.Request.Context()
	var requestBody entities.Drug

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logrus.Errorf("Error making binding JSON in method updating drug: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	id := c.Param("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.Errorf("Error parsing param id in method update drug: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	uintID := uint(u64)

	if err = h.ctrls.drug.Update(ctx, &requestBody, uintID); err != nil {
		logrus.Errorf("Error updating drug: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.StatusOK(nil))
}

//	 DeleteDrug
//
//		@Tags			Drugs
//		@Summary		Delete drug by id
//		@Description	Delete drug by id
//		@Accept			json
//		@Produce		json
//		@Param          id   path      string  true  "drugs ID"
//		@Param			Authorization	header		string	true	"access token"	extensions(x-example=token)
//		@Success		200				{object}	models.BaseResponseModel
//		@Failure		400				{object}	models.BaseResponseModel
//		@Router			/drugs/{id} [delete]
func (h *Handler) DeleteDrug(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.Errorf("Error parsing param id in method delete drug: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	uintID := uint(u64)

	err = h.ctrls.drug.Delete(ctx, uintID)
	if err != nil {
		logrus.Errorf("Error deleting drug with id %d: %s", uintID, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.StatusOK(nil))
}

//	 DeleteVaccination
//
//		@Tags			Vaccination
//		@Summary		Delete vaccination by id
//		@Description	Delete vaccination by id
//		@Accept			json
//		@Produce		json
//		@Param          id   path      string  true  "vaccination ID"
//		@Param			Authorization	header		string	true	"access token"	extensions(x-example=token)
//		@Success		200				{object}	models.BaseResponseModel
//		@Failure		400				{object}	models.BaseResponseModel
//		@Router			/vaccination/{id} [delete]
func (h *Handler) DeleteVaccination(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.Errorf("Error getting param id in method delete vaccination: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	uintID := uint(u64)
	err = h.ctrls.vaccination.Delete(ctx, uintID)
	if err != nil {
		logrus.Errorf("Error deleting drug with id %d: %s", uintID, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.StatusFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.StatusOK(nil))
}
