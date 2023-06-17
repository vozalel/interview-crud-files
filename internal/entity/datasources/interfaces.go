package datasources

import "github.com/vozalel/interview-crud-files/internal/entity/custom_error"

type Datasource struct {
	Name   string
	Source *string
}

type User struct {
	Name string
	ID   *uint64
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
	CreateDataSource(datasource *Datasource) *custom_error.CustomError
	ReadDataSource(datasource *Datasource) *custom_error.CustomError
	UpdateDataSource(datasource *Datasource) *custom_error.CustomError
	DeleteDataSource(datasource *Datasource) *custom_error.CustomError
}

type IManagerACL interface {
	// GetUserSourceACL - return CustomError.Code = 403 if user not have access to datasource
	GetUserSourceACL(user *User, datasource *Datasource) (UserDatasourceACL, *custom_error.CustomError)

	GrantUserSourceACL(user *User, datasource *Datasource, acl UserDatasourceACL) *custom_error.CustomError
	RevokeUserSourceACL(user *User, datasource *Datasource, acl UserDatasourceACL) *custom_error.CustomError
}
