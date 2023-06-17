package datasource

import (
	"context"
	"fmt"
	"github.com/vozalel/interview-crud-files/internal/entity/datasources"
	"testing"
)

var fileManager datasources.IManagerDatasource

func getManagerFile() datasources.IManagerDatasource {
	if fileManager == nil {
		fileManager = New("test_fm")
	}
	return fileManager
}

func TestFileManager_CreateDataSource(t *testing.T) {
	manager := getManagerFile()
	data := "test_data"

	datasource := datasources.Datasource{
		"test_name.csv",
		&data,
	}

	errCustom := manager.CreateDataSource(context.Background(), &datasource)
	if errCustom != nil {
		t.Error("managerACL.GrantUserSourceACL() error:", errCustom)
	}
}

func TestFileManager_ReadDataSource(t *testing.T) {
	manager := getManagerFile()

	datasource := datasources.Datasource{
		Name: "test_name.csv",
	}

	errCustom := manager.ReadDataSource(context.Background(), &datasource)
	if errCustom != nil {
		t.Error("managerACL.GrantUserSourceACL() error:", errCustom)
	}
}

func TestFileManager_UpdateDataSource(t *testing.T) {
	manager := getManagerFile()
	data := "test_data_update"

	datasource := datasources.Datasource{
		"test_name.csv",
		&data,
	}

	errCustom := manager.UpdateDataSource(context.Background(), &datasource)
	if errCustom != nil {
		t.Error("managerACL.GrantUserSourceACL() error:", errCustom)
	}
}

func TestFileManager_DeleteDataSource(t *testing.T) {
	manager := getManagerFile()

	datasource := datasources.Datasource{
		Name: "test_name.csv",
	}

	errCustom := manager.DeleteDataSource(context.Background(), &datasource)
	if errCustom != nil {
		t.Error("managerACL.GrantUserSourceACL() error:", errCustom)
	}
}

func TestFileManager_ListDataSource(t *testing.T) {
	manager := getManagerFile()

	list, errCustom := manager.ListDataSources(context.Background())
	if errCustom != nil {
		t.Error("managerACL.GrantUserSourceACL() error:", errCustom)
	}
	fmt.Println(list)
}
