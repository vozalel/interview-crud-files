package acl

import (
	"github.com/vozalel/interview-crud-files/internal/entity/datasources"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
	"github.com/vozalel/interview-crud-files/pkg/postgres"
)

func New(postgres *postgres.Postgres) datasources.IManagerACL {
	return &Acl{
		Postgres: postgres,
	}
}

type Acl struct {
	*postgres.Postgres
}

func (a Acl) GetUserSourceACL(user *datasources.User, datasource *datasources.Datasource) (datasources.UserDatasourceACL, *custom_error.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (a Acl) GrantUserSourceACL(user *datasources.User, datasource *datasources.Datasource, acl datasources.UserDatasourceACL) *custom_error.CustomError {
	//TODO implement me
	panic("implement me")
}

func (a Acl) RevokeUserSourceACL(user *datasources.User, datasource *datasources.Datasource, acl datasources.UserDatasourceACL) *custom_error.CustomError {
	//TODO implement me
	panic("implement me")
}
