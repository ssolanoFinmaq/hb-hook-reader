package dtos

import (
	"github.com/FinMaq/hook-reader/common/api"
)

type BatchUpdate struct {
	Inputs []Input `json:"inputs"`
}

type Input struct {
	ID         string     `json:"id"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	GenerarDocumentosLegalizacion string `json:"generar_documentos_legalizacion"`
}

func (batchUpdate *BatchUpdate) UpdateDeals() (interface{}, error) {

	var response interface{}

	err := api.PostJson("objects/deals/batch/update", &response, batchUpdate)
	if err != nil {
		return response, err
	}

	return response, nil
}
