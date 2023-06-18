package usecase

import (
	"context"
	"errors"
	"github.com/vozalel/interview-crud-files/internal/entity"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
	"github.com/vozalel/interview-crud-files/pkg/logger"
	"net/http"
)

func New(managerACL entity.IManagerACL,
	managerDatasource entity.IManagerDatasource) entity.IDatasourceUC {
	return &Usecase{
		managerACL:        managerACL,
		managerDatasource: managerDatasource,
	}
}

type Usecase struct {
	managerACL        entity.IManagerACL
	managerDatasource entity.IManagerDatasource
}

var (
	ErrorPermissionDenied = custom_error.New(
		errors.New("permission denied"),
		http.StatusForbidden,
		"permission denied",
	)
)

func (u *Usecase) CreateDataSource(
	ctx context.Context, user *entity.User,
	datasource *entity.Datasource) *custom_error.CustomError {

	aclPerform, errCustom := u.managerACL.GetUserPerformACL(ctx, user)
	if errCustom != nil {
		return errCustom.Wrap(
			"usecase - crud - CreateDataSource - u.managerACL.GetUserPerformACL()",
		)
	}
	if aclPerform.Create == false {
		return ErrorPermissionDenied
	}

	errCustom = u.managerDatasource.CreateDataSource(ctx, datasource)
	if errCustom != nil {
		return errCustom.Wrap(
			"usecase - crud - CreateDataSource - u.managerDatasource.CreateDataSource()",
		)
	}

	errCustom = u.managerACL.GrantUserSourceACL(ctx, user, datasource, entity.MaxSourcePermission)
	if errCustom != nil {
		logger.Instance.Warn("usecase - crud - CreateDataSource - u.managerDatasource.CreateDataSource()")
		// fixme transaction
		//  u.managerACL.GetUserPerformACL u.managerDatasource.CreateDataSource u.managerACL.GrantUserSourceACL
	}

	return nil
}

func (u *Usecase) ReadDataSource(ctx context.Context, user *entity.User, datasource *entity.Datasource) *custom_error.CustomError {
	aclSource, ErrCustom := u.managerACL.GetUserSourceACL(ctx, user, datasource)
	if ErrCustom != nil {
		return ErrCustom.Wrap(
			"usecase - crud - ReadDataSource - u.managerACL.GetUserSourceACL()",
		)
	}
	if aclSource.Read == false {
		return ErrorPermissionDenied
	}

	return u.managerDatasource.ReadDataSource(ctx, datasource)
}

func (u *Usecase) UpdateDataSource(ctx context.Context, user *entity.User, datasource *entity.Datasource) *custom_error.CustomError {
	aclSource, ErrCustom := u.managerACL.GetUserSourceACL(ctx, user, datasource)
	if ErrCustom != nil {
		return ErrCustom.Wrap(
			"usecase - crud - UpdateDataSource - u.managerACL.GetUserSourceACL()",
		)
	}
	if aclSource.Update == false {
		return ErrorPermissionDenied
	}

	return u.managerDatasource.UpdateDataSource(ctx, datasource)
}

func (u *Usecase) DeleteDataSource(ctx context.Context, user *entity.User, datasource *entity.Datasource) *custom_error.CustomError {
	aclSource, ErrCustom := u.managerACL.GetUserSourceACL(ctx, user, datasource)
	if ErrCustom != nil {
		return ErrCustom.Wrap(
			"usecase - crud - DeleteDataSource - u.managerACL.GetUserSourceACL()",
		)
	}
	if aclSource.Delete == false {
		return ErrorPermissionDenied
	}

	return u.managerDatasource.DeleteDataSource(ctx, datasource)
}

func (u *Usecase) ListDataSources(ctx context.Context, user *entity.User) ([]string, *custom_error.CustomError) {
	aclPerform, errCustom := u.managerACL.GetUserPerformACL(ctx, user)
	if errCustom != nil {
		return nil, errCustom.Wrap(
			"usecase - crud - ListDataSources - u.managerACL.GetUserPerformACL()",
		)
	}
	if aclPerform.List == false {
		return nil, ErrorPermissionDenied
	}
	return u.managerDatasource.ListDataSources(ctx)
}