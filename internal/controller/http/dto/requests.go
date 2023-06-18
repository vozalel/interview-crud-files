package dto

type Datasource struct {
	Name string `json:"name" form:"name"`
	Data []byte `json:"data"`
}
