package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/FinMaq/hook-reader/common/DB"
	"github.com/FinMaq/hook-reader/common/api"
	outputDtos "github.com/FinMaq/hook-reader/common/dtos"
	"github.com/FinMaq/hook-reader/hubspot/dtos"
	"github.com/FinMaq/hook-reader/reader/application/services"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	err := api.InitApiHubspot()
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprint(err),
			StatusCode: 500,
		}, err
	}
	err = DB.InitDb()
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprint(err),
			StatusCode: 500,
		}, err
	}

	products, err := services.GetPendingHooks(map[string]string{
		"generar_documentos_legalizacion": "true",
	})
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprint(err),
			StatusCode: 500,
		}, err
	}

	batchUpdate := dtos.BatchUpdate{
		Inputs: []dtos.Input{},
	}

	for _, product := range products {
		batchUpdate.Inputs = append(batchUpdate.Inputs, dtos.Input{
			ID: product.ID,
			Properties: dtos.Properties{
				GenerarDocumentosLegalizacion: "false",
			},
		})
	}

	log.Println(batchUpdate)

	_, err = batchUpdate.UpdateDeals()
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprint(err),
			StatusCode: 500,
		}, err
	}

	outputs := outputDtos.ProductsToOutputs(products)

	for _, output := range outputs {
		body, err := json.Marshal(output)
		if err != nil {
			log.Println(err)
		}
		resp, err := http.Post("http://localhost:8000/", "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Println(err)
			return events.APIGatewayProxyResponse{
				Body:       fmt.Sprint(err),
				StatusCode: 500,
			}, err
		}
		log.Println(resp.StatusCode)
		if resp.StatusCode != 200 {
			log.Println(resp.Body)
			return events.APIGatewayProxyResponse{
				Body:       fmt.Sprint(err),
				StatusCode: 500,
			}, err
		}
	}
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprint("Success"),
		StatusCode: 200,
	}, err
}
