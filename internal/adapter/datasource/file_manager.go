package datasource

import (
	"context"
	"github.com/vozalel/interview-crud-files/internal/entity/datasources"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
)

func NewFileManager(path string) datasources.IManagerDatasource {
	return &FileManager{Path: path}
}

type FileManager struct {
	Path string
}

func (f FileManager) CreateDataSource(ctx context.Context, datasource *datasources.Datasource) *custom_error.CustomError {
	//TODO implement me
	panic("implement me")
}

func (f FileManager) ReadDataSource(ctx context.Context, datasource *datasources.Datasource) *custom_error.CustomError {
	//TODO implement me
	panic("implement me")
}

func (f FileManager) UpdateDataSource(ctx context.Context, datasource *datasources.Datasource) *custom_error.CustomError {
	//TODO implement me
	panic("implement me")
}

func (f FileManager) DeleteDataSource(ctx context.Context, datasource *datasources.Datasource) *custom_error.CustomError {
	//TODO implement me
	panic("implement me")
}

func (f FileManager) ListDataSources(ctx context.Context) ([]datasources.DatasourceName, *custom_error.CustomError) {
	//TODO implement me
	panic("implement me")
}
