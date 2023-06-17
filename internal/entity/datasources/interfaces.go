package datasources

import (
	"context"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
)

//type DatasourceName string

type Datasource struct {
	Name string
	Data *string
}

type User struct {
	Name string
	ID   *int
}

type PerformACL struct {
	Create bool
	List   bool
}

type DatasourceACL struct {
	Read   bool
	Update bool
	Delete bool

	Grant  bool
	Revoke bool
}

type UserACL struct {
	PerformACL
	Sources map[string]DatasourceACL
}

type IManagerDatasource interface {
	CreateDataSource(ctx context.Context, datasource *Datasource) *custom_error.CustomError
	ReadDataSource(ctx context.Context, datasource *Datasource) *custom_error.CustomError
	UpdateDataSource(ctx context.Context, datasource *Datasource) *custom_error.CustomError
	DeleteDataSource(ctx context.Context, datasource *Datasource) *custom_error.CustomError

	ListDataSources(ctx context.Context) ([]string, *custom_error.CustomError)
}

type IDatasourceUC interface {
	CreateDataSource(ctx context.Context, user *User, datasource *Datasource) *custom_error.CustomError
	ReadDataSource(ctx context.Context, user *User, datasource *Datasource) *custom_error.CustomError
	UpdateDataSource(ctx context.Context, user *User, datasource *Datasource) *custom_error.CustomError
	DeleteDataSource(ctx context.Context, user *User, datasource *Datasource) *custom_error.CustomError

	ListDataSources(ctx context.Context, user *User) ([]string, *custom_error.CustomError)
}

type IManagerACL interface {
	// GetUserPerformACL - return CustomError.Code = 403 if user not have access to perform any action
	GetUserPerformACL(ctx context.Context, user *User) (PerformACL, *custom_error.CustomError)

	// GrantUserPerformACL - create new ACL record if not exist or update if exist
	GrantUserPerformACL(ctx context.Context, user *User, acl PerformACL) *custom_error.CustomError

	// RevokeUserPerformACL - delete ACL record
	RevokeUserPerformACL(ctx context.Context, user *User) *custom_error.CustomError

	// GetUserSourceACL - return CustomError.Code = 403 if user not have access to datasource
	GetUserSourceACL(ctx context.Context, user *User, datasource *Datasource) (DatasourceACL, *custom_error.CustomError)

	// GrantUserSourceACL - create new ACL record if not exist or update if exist
	GrantUserSourceACL(ctx context.Context, user *User, datasource *Datasource, acl DatasourceACL) *custom_error.CustomError

	// RevokeUserSourceACL - delete ACL record
	RevokeUserSourceACL(ctx context.Context, user *User, datasource *Datasource) *custom_error.CustomError
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
