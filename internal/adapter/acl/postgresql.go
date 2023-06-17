package acl

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/vozalel/interview-crud-files/internal/entity/datasources"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
	"github.com/vozalel/interview-crud-files/pkg/logger"
	"github.com/vozalel/interview-crud-files/pkg/postgres"
	"net/http"
)

func New(postgres *postgres.Postgres) datasources.IManagerACL {
	return &Acl{
		Postgres: postgres,
	}
}

type Acl struct {
	*postgres.Postgres
}

func (a *Acl) GetUserSourceACL(
	ctx context.Context, user *datasources.User,
	datasource *datasources.Datasource) (datasources.UserDatasourceACL, *custom_error.CustomError) {

	acl := datasources.UserDatasourceACL{}
	sql := `SELECT "create", read, update, delete, "grant", revoke FROM acl.acl WHERE user_id = $1 AND datasource_name = $2`
	args := []interface{}{user.ID, datasource.Name}
	err := a.Pool.QueryRow(
		ctx, sql, args...).Scan(
		&acl.Create,
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
		return acl, custom_error.NewCustomError(
			err,
			http.StatusForbidden,
			"permission denied",
		)
	}
	err = fmt.Errorf("adapter - acl - GetUserSourceACL - a.Pool.QueryRow(): %w", err)
	return acl, custom_error.NewCustomError(
		err,
		http.StatusInternalServerError,
		"permission check error, please try again later",
	)

}

func (a *Acl) GrantUserSourceACL(
	ctx context.Context, user *datasources.User,
	datasource *datasources.Datasource, acl datasources.UserDatasourceACL) *custom_error.CustomError {

	sql := `INSERT INTO 
	acl.acl(user_id, datasource_name, "create", read, update, delete, "grant", revoke)
	VALUES ($1, 	 $2, 			  $3, 	    $4,   $5, 	  $6, 	  $7,      $8)
	ON CONFLICT (user_id, datasource_name) 
	    DO UPDATE SET "create" = $3, read = $4, update = $5, delete = $6, "grant" = $7, revoke = $8
	`

	args := []interface{}{user.ID, datasource.Name, acl.Create, acl.Read, acl.Update, acl.Delete, acl.Grant, acl.Revoke}

	tag, err := a.Pool.Exec(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("adapter - acl - GrantUserSourceACL - a.Pool.Exec(): %w", err)
		return custom_error.NewCustomError(
			err,
			http.StatusInternalServerError,
			"permission assignment error, please try again later",
		)
	}

	if tag.RowsAffected() == 0 {
		err = fmt.Errorf("adapter - acl - GrantUserSourceACL - tag.RowsAffected() == 0")
		return custom_error.NewCustomError(
			err,
			http.StatusInternalServerError,
			"permission assignment error, please try again later",
		)
	}
	return nil
}

func (a *Acl) RevokeUserSourceACL(
	ctx context.Context, user *datasources.User,
	datasource *datasources.Datasource) *custom_error.CustomError {
	sql := `DELETE FROM acl.acl WHERE user_id = $1 AND datasource_name = $2`
	args := []interface{}{user.ID, datasource.Name}

	tag, err := a.Pool.Exec(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("adapter - acl - RevokeUserSourceACL - a.Pool.Exec(): %w", err)
		return custom_error.NewCustomError(
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
