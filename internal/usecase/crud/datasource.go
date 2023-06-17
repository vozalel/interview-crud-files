package crud

import (
	"context"
	"github.com/vozalel/interview-crud-files/internal/entity/datasources"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
)

func New(managerACL datasources.IManagerACL,
	managerDatasource datasources.IManagerDatasource) datasources.IManagerDatasource {
	return &Usecase{
		managerACL:        managerACL,
		managerDatasource: managerDatasource,
	}
}

type Usecase struct {
	managerACL        datasources.IManagerACL
	managerDatasource datasources.IManagerDatasource
}

func (u *Usecase) CreateDataSource(ctx context.Context, datasource *datasources.Datasource) *custom_error.CustomError {
	// FIXME: check ACL
	return u.managerDatasource.CreateDataSource(ctx, datasource)
}

func (u *Usecase) ReadDataSource(ctx context.Context, datasource *datasources.Datasource) *custom_error.CustomError {
	// FIXME: check ACL
	return u.managerDatasource.ReadDataSource(ctx, datasource)
}

func (u *Usecase) UpdateDataSource(ctx context.Context, datasource *datasources.Datasource) *custom_error.CustomError {
	// FIXME: check ACL
	return u.managerDatasource.UpdateDataSource(ctx, datasource)
}

func (u *Usecase) DeleteDataSource(ctx context.Context, datasource *datasources.Datasource) *custom_error.CustomError {
	// FIXME: check ACL
	return u.managerDatasource.DeleteDataSource(ctx, datasource)
}

func (u *Usecase) ListDataSources(ctx context.Context) ([]datasources.DatasourceName, *custom_error.CustomError) {
	// FIXME: check ACL
	return u.managerDatasource.ListDataSources(ctx)
}
