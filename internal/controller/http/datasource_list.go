package http

import (
	"github.com/gin-gonic/gin"
	"github.com/vozalel/interview-crud-files/internal/entity"
)

type datasourceListRoutes struct {
	datasourceUC entity.IDatasourceUC
}

func newDatasourceListRoutes(handler *gin.RouterGroup, datasourceUC entity.IDatasourceUC) {
	r := &datasourceRoutes{datasourceUC}

	{
		handler.GET("/", r.readDatasourceList)
	}

}

// @Summary     Read datasource list
// @Description Read datasource list
// @ID          readDatasourceList
// @Tags  	    datasource list
// @Produce     json
// @Success     200 {object} ResponseDatasourceList
// @Failure     400 {object} properErrorResponse "Incorrect request"
// @Failure     403 {object} properErrorResponse "Permission deny"
// @Failure     404 {object} properErrorResponse "Not found"
// @Failure		500 {object} properErrorResponse "Internal error"
// @Router      /list [get]
func (datasourceRoutes *datasourceRoutes) readDatasourceList(ctx *gin.Context) {

}
