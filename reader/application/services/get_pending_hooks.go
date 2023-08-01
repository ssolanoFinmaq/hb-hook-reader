package services

import (
	"github.com/FinMaq/hook-reader/common/DB"
	"github.com/FinMaq/hook-reader/reader/domain/entities"
)

func GetPendingHooks(filters map[string]string) ([]entities.Products, error) {
	filter := map[string]string{
		"generar_documentos_legalizacion": "true",
	}
	products := []entities.Products{}
	err := DB.Instance.Joins("Client").Joins("Distribuitor").Where(filter).Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}
