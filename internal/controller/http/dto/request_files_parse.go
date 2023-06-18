package dto

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/vozalel/interview-crud-files/internal/entity"
	"io"
	"mime/multipart"
)

func ParseRequestBody(fileHeaders map[string][]*multipart.FileHeader) (datasource entity.Datasource, err error) {
	var body bytes.Buffer

	writer := multipart.NewWriter(&body)
	defer writer.Close()
	fileNames := fileHeaders["files"]

	var files []multipart.File
	defer func() {
		for _, file := range files {
			_ = file.Close()
		}
	}()

	switch len(fileNames) {
	case 0:
		return datasource, errors.New("multipart form parsing error")
	case 1:
		//ok
	default:
		return datasource, errors.New("uploading more than one file is not provided")
	}

	file := fileNames[0]
	f, err := file.Open()
	defer f.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, f); err != nil {
		return datasource, fmt.Errorf("dto - requestParseFiles - io.Copy: %w", err)
	}

	datasource.Name = file.Filename
	datasource.Data = buf.Bytes()

	return datasource, nil
}
