package datasources

import (
	"context"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
)

type DatasourceName string

type Datasource struct {
	Name DatasourceName
	Data *string
}

type User struct {
	Name string
	ID   *int
}

type UserDatasourceACL struct {
	Create bool
	Read   bool
	Update bool
	Delete bool

	Grant  bool
	Revoke bool
}

type IManagerDatasource interface {
	CreateDataSource(ctx context.Context, datasource *Datasource) *custom_error.CustomError
	ReadDataSource(ctx context.Context, datasource *Datasource) *custom_error.CustomError
	UpdateDataSource(ctx context.Context, datasource *Datasource) *custom_error.CustomError
	DeleteDataSource(ctx context.Context, datasource *Datasource) *custom_error.CustomError

	ListDataSources(ctx context.Context) ([]DatasourceName, *custom_error.CustomError)
}

type IManagerACL interface {
	// GetUserSourceACL - return CustomError.Code = 403 if user not have access to datasource
	GetUserSourceACL(ctx context.Context, user *User, datasource *Datasource) (UserDatasourceACL, *custom_error.CustomError)

	GrantUserSourceACL(ctx context.Context, user *User, datasource *Datasource, acl UserDatasourceACL) *custom_error.CustomError
	RevokeUserSourceACL(ctx context.Context, user *User, datasource *Datasource) *custom_error.CustomError
}
