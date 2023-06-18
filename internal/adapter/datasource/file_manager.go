package datasource

import (
	"context"
	"fmt"
	"github.com/vozalel/interview-crud-files/internal/entity"
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
	"io/fs"
	"net/http"
	"os"
	"path"
)

func New(path string) entity.IManagerDatasource {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		panic(fmt.Errorf("adapter - datasource - New: %w", err))
	}
	return &FileManager{Path: path}
}

type FileManager struct {
	Path string
}

func (f *FileManager) CreateDataSource(
	ctx context.Context,
	datasource *entity.Datasource) *custom_error.CustomError {

	stat, err := os.Stat(path.Join(f.Path, datasource.Name))
	if err != nil {
		if !os.IsNotExist(err) {
			return custom_error.New(
				fmt.Errorf("adapter - FileManager - CreateDataSource - os.Stat(): %w", err),
				http.StatusInternalServerError,
				"file system error",
			)
		}
	}

	if stat != nil {
		return custom_error.New(
			fmt.Errorf("adapter - FileManager - CreateDataSource - os.Stat(): %w", err),
			http.StatusConflict,
			"file already exist",
		)
	}

	err = os.WriteFile(path.Join(f.Path, datasource.Name), datasource.Data, fs.FileMode(0777))
	if err != nil {
		return custom_error.New(
			fmt.Errorf("adapter - FileManager - CreateDataSource - os.WriteFile(): %w", err),
			http.StatusInternalServerError,
			"file manager error, please try again later",
		)
	}
	return nil
}

func (f *FileManager) ReadDataSource(
	ctx context.Context,
	datasource *entity.Datasource) *custom_error.CustomError {

	data, err := os.ReadFile(path.Join(f.Path, datasource.Name))
	if err != nil {
		// fs.PathError
		return custom_error.New(
			fmt.Errorf("adapter - FileManager - CreateDataSource - os.ReadFile(): %w", err),
			http.StatusInternalServerError,
			"file manager error, please try again later",
		)
	}

	datasource.Data = data
	return nil
}

func (f *FileManager) UpdateDataSource(
	ctx context.Context,
	datasource *entity.Datasource) *custom_error.CustomError {

	_, err := os.Stat(path.Join(f.Path, datasource.Name))
	if err != nil {
		if os.IsNotExist(err) {
			return custom_error.New(
				fmt.Errorf("adapter - FileManager - UpdateDataSource - os.Stat(): %w", err),
				http.StatusNotFound,
				"file not exist",
			)
		}
	}

	err = os.WriteFile(path.Join(f.Path, datasource.Name), datasource.Data, fs.FileMode(0777))
	if err != nil {
		return custom_error.New(
			fmt.Errorf("adapter - FileManager - UpdateDataSource - os.WriteFile(): %w", err),
			http.StatusInternalServerError,
			"file manager error, please try again later",
		)
	}
	return nil
}

func (f *FileManager) DeleteDataSource(ctx context.Context, datasource *entity.Datasource) *custom_error.CustomError {
	err := os.Remove(path.Join(f.Path, datasource.Name))
	if err != nil {
		return custom_error.New(
			fmt.Errorf("adapter - FileManager - DeleteDataSource - os.Remove(): %w", err),
			http.StatusInternalServerError,
			"file manager error, please try again later",
		)
	}
	return nil
}

func (f *FileManager) ListDataSources(ctx context.Context) ([]string, *custom_error.CustomError) {

	fileNames := make([]string, 0)

	dirEntry, err := os.ReadDir(f.Path)
	if err != nil {
		return nil, custom_error.New(
			fmt.Errorf("adapter - FileManager - ListDataSources - os.Remove(): %w", err),
			http.StatusInternalServerError,
			"file manager error, please try again later",
		)
	}

	for _, entry := range dirEntry {
		// todo need recursion for folders
		fileNames = append(fileNames, entry.Name())
	}

	return fileNames, nil
}
