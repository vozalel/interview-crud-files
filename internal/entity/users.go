package entity

import (
	"context"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
)

type User struct {
	Name string
	ID   *int
}

type UserACL struct {
	PerformACL
	Sources map[string]DatasourceACL
}

//go:generate mockgen -source=users.go -destination=./mock/users_mocks_test.go -package=mock

type IDatasourceUC interface {
	CreateDataSource(ctx context.Context, user *User, datasource *Datasource) *custom_error.CustomError
	ReadDataSource(ctx context.Context, user *User, datasource *Datasource) *custom_error.CustomError
	UpdateDataSource(ctx context.Context, user *User, datasource *Datasource) *custom_error.CustomError
	DeleteDataSource(ctx context.Context, user *User, datasource *Datasource) *custom_error.CustomError

	ListDataSources(ctx context.Context, user *User) ([]string, *custom_error.CustomError)
}
