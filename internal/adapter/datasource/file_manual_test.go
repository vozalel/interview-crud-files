package datasource

import (
	"context"
	"fmt"
	"github.com/vozalel/interview-crud-files/internal/entity"
	"testing"
)

var fileManager entity.IManagerDatasource

func getManagerFile() entity.IManagerDatasource {
	if fileManager == nil {
		fileManager = New("test_fm")
	}
	return fileManager
}

func TestFileManager_CreateDataSource(t *testing.T) {
	manager := getManagerFile()

	datasource := entity.Datasource{
		Name: "test_name.csv",
		Data: []byte("test_data"),
	}

	errCustom := manager.CreateDataSource(context.Background(), &datasource)
	if errCustom != nil {
		t.Error("managerACL.GrantUserSourceACL() error:", errCustom)
	}
}

func TestFileManager_ReadDataSource(t *testing.T) {
	manager := getManagerFile()

	datasource := entity.Datasource{
		Name: "test_name.csv",
	}

	errCustom := manager.ReadDataSource(context.Background(), &datasource)
	if errCustom != nil {
		t.Error("managerACL.GrantUserSourceACL() error:", errCustom)
	}
}

func TestFileManager_UpdateDataSource(t *testing.T) {
	manager := getManagerFile()

	datasource := entity.Datasource{
		Name: "test_name.csv",
		Data: []byte("test_data"),
	}

	errCustom := manager.UpdateDataSource(context.Background(), &datasource)
	if errCustom != nil {
		t.Error("managerACL.GrantUserSourceACL() error:", errCustom)
	}
}

func TestFileManager_DeleteDataSource(t *testing.T) {
	manager := getManagerFile()

	datasource := entity.Datasource{
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
