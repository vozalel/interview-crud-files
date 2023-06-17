package datasource

import (
	"github.com/vozalel/interview-crud-files/internal/entity/datasources"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
)

func NewFileManager(path string) datasources.IManagerDatasource {
	return &FileManager{Path: path}
}

type FileManager struct {
	Path string
}

func (f *FileManager) CreateDataSource(datasource *datasources.Datasource) *custom_error.CustomError {
	//TODO implement me
	panic("implement me")
}

func (f *FileManager) ReadDataSource(datasource *datasources.Datasource) *custom_error.CustomError {
	//TODO implement me
	panic("implement me")
}

func (f *FileManager) UpdateDataSource(datasource *datasources.Datasource) *custom_error.CustomError {
	//TODO implement me
	panic("implement me")
}

func (f *FileManager) DeleteDataSource(datasource *datasources.Datasource) *custom_error.CustomError {
	//TODO implement me
	panic("implement me")
}

func (f *FileManager) ListDataSources() ([]datasources.DatasourceName, *custom_error.CustomError) {
	//TODO implement me
	panic("implement me")
}
