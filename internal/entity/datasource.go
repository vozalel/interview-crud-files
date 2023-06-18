package entity

import (
	"context"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
)

type Datasource struct {
	Name string `json:"name"`
	Data []byte
}

//go:generate mockgen -source=datasource.go -destination=./mock/datasource_mocks.go -package=mock

type IManagerDatasource interface {
	CreateDataSource(ctx context.Context, datasource *Datasource) *custom_error.CustomError
	ReadDataSource(ctx context.Context, datasource *Datasource) *custom_error.CustomError
	UpdateDataSource(ctx context.Context, datasource *Datasource) *custom_error.CustomError
	DeleteDataSource(ctx context.Context, datasource *Datasource) *custom_error.CustomError

	ListDataSources(ctx context.Context) ([]string, *custom_error.CustomError)
}

var (
	MaxSourcePermission = DatasourceACL{
		Read:   true,
		Update: true,
		Delete: true,
		Grant:  true,
		Revoke: true,
	}
)
