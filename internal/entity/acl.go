package entity

import (
	"context"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
)

type DatasourceACL struct {
	Read   bool
	Update bool
	Delete bool

	Grant  bool
	Revoke bool
}

//go:generate mockgen -source=acl.go -destination=./mock/acl_mocks.go -package=mock

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