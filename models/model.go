package models

type Model struct {
	ProjectName string              `json:"projectName"`
	Entity      map[string][]string `json:"entity"`
	Relation    []map[string]string `json:"relation"`
	Database    map[string]string   `json:"database"`
}

type EntityValue struct {
	EntityValueName          string
	EntityDataType           string
	EntityValueNameSnakeCase string
	EntityValueNameCamelCase string
}

type CreateModelsInput struct {
	EntityName  string
	EntityValue []EntityValue
}
