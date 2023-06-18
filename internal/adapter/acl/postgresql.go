package acl

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/vozalel/interview-crud-files/internal/entity"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
	"github.com/vozalel/interview-crud-files/pkg/logger"
	"github.com/vozalel/interview-crud-files/pkg/postgres"
	"net/http"
)

func New(postgres *postgres.Postgres) entity.IManagerACL {
	return &Acl{
		Postgres: postgres,
	}
}

const (
	tableACLUserDatasource = "acl.user_datasource"
	tableACLUserPerform    = "acl.user_perform"
)

type Acl struct {
	*postgres.Postgres
}

func (a *Acl) GetUserPerformACL(
	ctx context.Context,
	user *entity.User) (entity.PerformACL, *custom_error.CustomError) {

	acl := entity.PerformACL{}
	sql := `SELECT "create", list FROM ` + tableACLUserPerform + ` WHERE user_id = $1`
	args := []interface{}{*user.ID}
	err := a.Pool.QueryRow(
		ctx, sql, args...).Scan(
		&acl.Create,
		&acl.List,
	)

	if err == nil {
		return acl, nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		err = fmt.Errorf("adapter - acl - GetUserPerformACL - errors.Is(err, pgx.ErrNoRows)")
		return acl, custom_error.New(
			err,
			http.StatusForbidden,
			"permission denied",
		)
	}
	err = fmt.Errorf("adapter - acl - GetUserPerformACL - a.Pool.QueryRow(): %w", err)
	return acl, custom_error.New(
		err,
		http.StatusInternalServerError,
		"internal server error",
	)
}

func (a *Acl) GrantUserPerformACL(
	ctx context.Context, user *entity.User,
	acl entity.PerformACL) *custom_error.CustomError {

	sql := `INSERT INTO` + tableACLUserPerform +
		`(user_id, "create", list)
	VALUES				  
		 ($1, 	   $2, 	     $3)
	ON CONFLICT (user_id, datasource_name) 
	    DO UPDATE SET "create" = $2, list = $3
	`

	args := []interface{}{
		user.ID,
		acl.Create,
		acl.List,
	}

	tag, err := a.Pool.Exec(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("adapter - acl - GrantUserSourceACL - a.Pool.Exec(): %w", err)
		return custom_error.New(
			err,
			http.StatusInternalServerError,
			"permission assignment error, please try again later",
		)
	}

	if tag.RowsAffected() == 0 {
		err = fmt.Errorf("adapter - acl - GrantUserSourceACL - tag.RowsAffected() == 0")
		return custom_error.New(
			err,
			http.StatusInternalServerError,
			"permission assignment error, please try again later",
		)
	}
	return nil
}

func (a *Acl) RevokeUserPerformACL(
	ctx context.Context, user *entity.User) *custom_error.CustomError {

	sql := `DELETE FROM ` + tableACLUserPerform + `WHERE user_id = $1`
	args := []interface{}{user.ID}
	tag, err := a.Pool.Exec(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("adapter - acl - RevokeUserPerformACL - a.Pool.Exec(): %w", err)
		return custom_error.New(
			err,
			http.StatusInternalServerError,
			"permission revocation error, please try again later",
		)
	}

	if tag.RowsAffected() == 0 {
		err = fmt.Errorf("adapter - acl - RevokeUserPerformACL - tag.RowsAffected() == 0")
		return custom_error.New(
			err,
			http.StatusInternalServerError,
			"permission revocation error, please try again later",
		)
	}
	return nil
}

func (a *Acl) GetUserSourceACL(
	ctx context.Context, user *entity.User,
	datasource *entity.Datasource) (entity.DatasourceACL, *custom_error.CustomError) {

	acl := entity.DatasourceACL{}
	sql := `SELECT read, update, delete, "grant", revoke FROM ` + tableACLUserDatasource + ` WHERE user_id = $1 AND datasource_name = $2`
	args := []interface{}{user.ID, datasource.Name}
	err := a.Pool.QueryRow(
		ctx, sql, args...).Scan(
		&acl.Read,
		&acl.Update,
		&acl.Delete,
		&acl.Grant,
		&acl.Revoke,
	)

	if err == nil {
		return acl, nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		err = fmt.Errorf("adapter - acl - GetUserSourceACL - errors.Is(err, pgx.ErrNoRows)")
		return acl, custom_error.New(
			err,
			http.StatusForbidden,
			"permission denied",
		)
	}
	err = fmt.Errorf("adapter - acl - GetUserSourceACL - a.Pool.QueryRow(): %w", err)
	return acl, custom_error.New(
		err,
		http.StatusInternalServerError,
		"permission check error, please try again later",
	)

}

func (a *Acl) GrantUserSourceACL(
	ctx context.Context, user *entity.User,
	datasource *entity.Datasource,
	acl entity.DatasourceACL) *custom_error.CustomError {

	sql := `INSERT INTO ` + tableACLUserDatasource +
		`(user_id, datasource_name, read, update, delete, "grant", revoke)
	VALUES ($1, 	 $2, 			  $3,   $4,     $5,     $6, 	 $7)
	ON CONFLICT (user_id, datasource_name) 
	    DO UPDATE SET read = $3, update = $4, delete = $5, "grant" = $6, revoke = $7
	`

	args := []interface{}{
		user.ID,
		datasource.Name,
		acl.Read,
		acl.Update,
		acl.Delete,
		acl.Grant,
		acl.Revoke,
	}

	tag, err := a.Pool.Exec(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("adapter - acl - GrantUserSourceACL - a.Pool.Exec(): %w", err)
		return custom_error.New(
			err,
			http.StatusInternalServerError,
			"permission assignment error, please try again later",
		)
	}

	if tag.RowsAffected() == 0 {
		err = fmt.Errorf("adapter - acl - GrantUserSourceACL - tag.RowsAffected() == 0")
		return custom_error.New(
			err,
			http.StatusInternalServerError,
			"permission assignment error, please try again later",
		)
	}
	return nil
}

func (a *Acl) RevokeUserSourceACL(
	ctx context.Context, user *entity.User,
	datasource *entity.Datasource) *custom_error.CustomError {

	sql := `DELETE FROM ` + tableACLUserDatasource + ` WHERE user_id = $1 AND datasource_name = $2`
	args := []interface{}{user.ID, datasource.Name}

	tag, err := a.Pool.Exec(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("adapter - acl - RevokeUserSourceACL - a.Pool.Exec(): %w", err)
		return custom_error.New(
			err,
			http.StatusInternalServerError,
			"permission assignment error, please try again later",
		)
	}

	if tag.RowsAffected() == 0 {
		logger.Instance.Warn(
			"adapter - acl - RevokeUserSourceACL - tag.RowsAffected() == 0",
			"user_id", user.ID,
			"datasource_name", datasource.Name,
		)
	}

	return nil
}
