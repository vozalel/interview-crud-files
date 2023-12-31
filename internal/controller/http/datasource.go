package http

import (
	"fmt"
	"github.com/vozalel/interview-crud-files/internal/controller/http/dto"
	"github.com/vozalel/interview-crud-files/internal/entity"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type datasourceRoutes struct {
	datasourceUC entity.IDatasourceUC
}

func newDatasourceRoutes(handler *gin.RouterGroup, datasourceUC entity.IDatasourceUC) {
	r := &datasourceRoutes{datasourceUC}

	{
		handler.POST("/", r.createDatasource)
		handler.GET("/", r.readDatasource)
		handler.PATCH("/", r.updateDatasource)
		handler.DELETE("/", r.deleteDatasource)
	}
}

// createDatasource godoc
// @Summary     create datasource
// @Description create datasource by datasource name
// @ID          createDatasource
// @Tags  	    datasource
// @Accept      multipart/form-data
// @Produce     json
// @Param 		files formData []file true "files to download"
// @Success     200 {string} string
// @Failure     400 {object} properErrorResponse "Incorrect request"
// @Failure     403 {object} properErrorResponse "Permission deny"
// @Failure     409 {object} properErrorResponse "File already exist"
// @Failure		500 {object} properErrorResponse "Internal error"
// @Router      /source [post]
func (datasourceRoutes *datasourceRoutes) createDatasource(ctx *gin.Context) {
	ctxNew := ctx.Request.Context()

	user, ok := ctx.Value(dto.ContextKeyUser).(entity.User)
	if !ok {
		respondWithCustomError(ctx,
			custom_error.New(
				dto.ErrorUserNotFoundInContext,
				http.StatusBadRequest,
				dto.MessageIncorrectRequest,
			),
		)
		return
	}

	if err := ctx.Request.ParseMultipartForm(32 << 20); err != nil {
		respondWithCustomError(
			ctx,
			custom_error.New(
				fmt.Errorf("c.Request.ParseMultipartForm() error:%w", err),
				http.StatusBadRequest,
				dto.MessageIncorrectRequest,
			),
		)
		return
	}

	datasource, err := dto.ParseRequestBody(ctx.Request.MultipartForm.File)
	if err != nil {
		respondWithCustomError(ctx,
			custom_error.New(
				fmt.Errorf("http - routes - createDatasource - dto.ParseRequestBody():%w", err),
				http.StatusBadRequest,
				dto.MessageIncorrectRequest,
			),
		)
		return
	}

	errCustom := datasourceRoutes.datasourceUC.CreateDataSource(ctxNew, &user, &datasource)
	if errCustom != nil {
		respondWithCustomError(
			ctx,
			errCustom.Wrap(
				"http - routes - createDatasource datasourceRoutes.datasourceUC.CreateDataSource()",
			),
		)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageOk)
}

// @Summary     read datasource
// @Description read datasource
// @ID          readDatasource
// @Tags  	    datasource
// @Produce     json
// @Param       name query string true "name"
// @Success     200 {object} ResponseDatasource
// @Failure     400 {object} properErrorResponse "Incorrect request"
// @Failure     403 {object} properErrorResponse "Permission deny"
// @Failure     404 {object} properErrorResponse "Not found"
// @Failure		500 {object} properErrorResponse "Internal error"
// @Router      /source [get]
func (datasourceRoutes *datasourceRoutes) readDatasource(ctx *gin.Context) {
	var (
		datasource entity.Datasource
	)

	ctxNew := ctx.Request.Context()

	user, ok := ctx.Value(dto.ContextKeyUser).(entity.User)
	if !ok {
		respondWithCustomError(ctx,
			custom_error.New(
				dto.ErrorUserNotFoundInContext,
				http.StatusBadRequest,
				dto.MessageIncorrectRequest,
			),
		)
		return
	}

	datasourceDTO := dto.Datasource{}
	if err := ctx.ShouldBindQuery(&datasourceDTO); err != nil {
		respondWithCustomError(
			ctx,
			custom_error.New(
				fmt.Errorf("http - routes - readDatasource - dto.ParseRequestBody():%w", err),
				http.StatusBadRequest,
				dto.MessageIncorrectRequest,
			),
		)
		return
	}

	datasource = entity.Datasource{
		Name: datasourceDTO.Name,
	}

	errCustom := datasourceRoutes.datasourceUC.ReadDataSource(ctxNew, &user, &datasource)
	if errCustom != nil {
		respondWithCustomError(ctx,
			errCustom.Wrap("http - routes - readDatasource - datasourceRoutes.datasourceUC.ReadDataSource()"),
		)
		return
	}

	// TODO: send form-data request
	ctx.JSON(http.StatusOK, datasource)
}

// @Summary     update datasource
// @Description update datasource by name
// @ID          datasource_update
// @Tags  	    datasource
// @Accept      multipart/form-data
// @Produce     json
// @Param 		files formData []file true "files to update"
// @Success     200 {string} string "ok"
// @Failure     400 {object} properErrorResponse "Incorrect request"
// @Failure     403 {object} properErrorResponse "Permission deny"
// @Failure     404 {object} properErrorResponse "File not exist"
// @Failure		500 {object} properErrorResponse "Internal error"
// @Router      /source [patch]
func (datasourceRoutes *datasourceRoutes) updateDatasource(ctx *gin.Context) {
	var (
		datasource entity.Datasource
	)

	ctxNew := ctx.Request.Context()

	user, ok := ctx.Value(dto.ContextKeyUser).(entity.User)
	if !ok {
		respondWithCustomError(ctx,
			custom_error.New(
				dto.ErrorUserNotFoundInContext,
				http.StatusBadRequest,
				dto.MessageIncorrectRequest,
			),
		)
		return
	}

	if err := ctx.Request.ParseMultipartForm(32 << 20); err != nil {
		respondWithCustomError(
			ctx,
			custom_error.New(
				fmt.Errorf("c.Request.ParseMultipartForm() error:%w", err),
				http.StatusBadRequest,
				dto.MessageIncorrectRequest,
			),
		)
		return
	}

	datasource, err := dto.ParseRequestBody(ctx.Request.MultipartForm.File)
	if err != nil {
		respondWithCustomError(ctx,
			custom_error.New(
				fmt.Errorf("http - routes - updateDatasource - dto.ParseRequestBody():%w", err),
				http.StatusBadRequest,
				dto.MessageIncorrectRequest,
			),
		)
		return
	}

	errCustom := datasourceRoutes.datasourceUC.UpdateDataSource(ctxNew, &user, &datasource)
	if errCustom != nil {
		respondWithCustomError(
			ctx,
			errCustom.Wrap(
				"http - routes - updateDatasource datasourceRoutes.datasourceUC.CreateDataSource()",
			),
		)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageOk)
}

// @Summary     delete datasource
// @Description delete datasource by name
// @ID          datasource_delete
// @Tags  	    datasource
// @Produce     json
// @Param       name query string true "delete by name"
// @Success     200 {string} string "ok"
// @Failure     403 {object} properErrorResponse "No data or no access to it"
// @Failure		500 {object} properErrorResponse "Internal error"
// @Router      /source [delete]
func (datasourceRoutes *datasourceRoutes) deleteDatasource(ctx *gin.Context) {
	var (
		datasourceDTO dto.Datasource
		datasource    entity.Datasource
	)

	ctxNew := ctx.Request.Context()

	user, ok := ctx.Value(dto.ContextKeyUser).(entity.User)
	if !ok {
		respondWithCustomError(ctx,
			custom_error.New(
				dto.ErrorUserNotFoundInContext,
				http.StatusBadRequest,
				dto.MessageIncorrectRequest,
			),
		)
		return
	}

	if err := ctx.ShouldBindQuery(&datasourceDTO); err != nil {
		respondWithCustomError(
			ctx,
			custom_error.New(
				fmt.Errorf("http - routes - deleteDatasource - dto.ParseRequestBody():%w", err),
				http.StatusBadRequest,
				dto.MessageIncorrectRequest,
			),
		)
		return
	}

	datasource = entity.Datasource{
		Name: datasourceDTO.Name,
		Data: datasourceDTO.Data,
	}

	errCustom := datasourceRoutes.datasourceUC.DeleteDataSource(ctxNew, &user, &datasource)
	if errCustom != nil {
		respondWithCustomError(
			ctx,
			errCustom.Wrap(
				"http - routes - deleteDatasource datasourceRoutes.datasourceUC.CreateDataSource()",
			),
		)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageOk)
}
