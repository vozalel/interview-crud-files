package acl

import (
	"context"
	"fmt"
	"github.com/vozalel/interview-crud-files/internal/entity"
	"github.com/vozalel/interview-crud-files/pkg/logger"
	"github.com/vozalel/interview-crud-files/pkg/postgres"
	"testing"
)

var code = 1 // if m.Run not started exit code set 1
var managerACL entity.IManagerACL

func getManagerACL() entity.IManagerACL {
	if managerACL == nil {
		logger.Init("test managerACL adapter", "local", "debug")

		pg, err := postgres.New(
			"postgres://user:pass@localhost:5432/postgres?pool_max_conn_idle_time=30s",
			logger.Instance,
		)
		if err != nil {
			logger.Instance.Fatal(fmt.Errorf("postgres.New error: %v", err))
		}
		managerACL = New(pg)
	}
	return managerACL
}

func TestGrantUserSourceACL(t *testing.T) {
	manager := getManagerACL()

	userID := 1
	datasourceData := "xml"

	user := entity.User{
		Name: "admin",
		ID:   &userID,
	}

	datasource := entity.Datasource{
		Name: "test source",
		Data: &datasourceData,
	}

	acl := entity.DatasourceACL{
		Read:   true,
		Update: true,
		Delete: true,

		Grant:  true,
		Revoke: false,
	}

	errCustom := manager.GrantUserSourceACL(context.Background(), &user, &datasource, acl)
	if errCustom != nil {
		t.Error("managerACL.GrantUserSourceACL() error:", errCustom)
	}
}

func TestGetUserSourceACL(t *testing.T) {
	manager := getManagerACL()

	userID := 1
	datasourceData := "xml"

	user := entity.User{
		Name: "admin",
		ID:   &userID,
	}

	datasource := entity.Datasource{
		Name: "test source",
		Data: &datasourceData,
	}

	acl, errCustom := manager.GetUserSourceACL(context.Background(), &user, &datasource)
	if errCustom != nil {
		t.Error("managerACL.GetUserSourceACL() error:", errCustom)
	}
	t.Log("acl:", acl)
}

func TestRevokeUserSourceACL(t *testing.T) {
	manager := getManagerACL()

	userID := 1
	datasourceData := "xml"

	user := entity.User{
		Name: "admin",
		ID:   &userID,
	}

	datasource := entity.Datasource{
		Name: "test source",
		Data: &datasourceData,
	}

	errCustom := manager.RevokeUserSourceACL(context.Background(), &user, &datasource)
	if errCustom != nil {
		t.Error("managerACL.RevokeUserSourceACL() error:", errCustom)
	}
}
