package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/FinMaq/hook-reader/common/secrets"
)

var ApiHubspot *Client

func InitApiHubspot() error {
	secretStr, err := secrets.GetSecretString("HbDataWarehouse")
	if err != nil {
		return err
	}
	secretValues := struct {
		HapiKeyDataWarehouse string
	}{}
	err = json.Unmarshal([]byte(secretStr), &secretValues)
	if err != nil {
		return err
	}
	apiHubspot, err := NewApi("https://api.hubapi.com/crm/v3/", secretValues.HapiKeyDataWarehouse)
	ApiHubspot = apiHubspot
	if err != nil {
		return err
	}
	return nil
}

func GetJson(url string, target interface{}) error {

	target, err := ApiHubspot.Get(url, &target)
	if err != nil {
		log.Println(fmt.Sprintf("Ha Fallado la peticion a: %s", url), fmt.Sprintf("Error: %s", err))
		return err
	}

	//logs.Log("API", "GET Exitosa", fmt.Sprintf("Petición Exitosa: %s", url))

	return nil
}

func PostJson(url string, target interface{}, body interface{}) error {
	//fmt.Println(body)
	target, err := ApiHubspot.Post(url, &target, body)
	if err != nil {
		log.Println(fmt.Sprintf("Ha Fallado la peticion a: %s", url), fmt.Sprintf("Error: %s", err))
		return err
	}
	//logs.Log("API", "POST Exitosa", fmt.Sprintf("Petición Exitosa: %s", url))

	return nil
}

func PatchJson(url string, target interface{}, body interface{}) error {
	//fmt.Println(body)
	log.Println(url)
	log.Println(target)
	log.Println(body)

	target, err := ApiHubspot.Patch(url, &target, body)
	if err != nil {
		log.Println(fmt.Sprintf("Ha Fallado la peticion a: %s", url), fmt.Sprintf("Error: %s", err))
		return err
	}
	//logs.Log("API", "POST Exitosa", fmt.Sprintf("Petición Exitosa: %s", url))

	return nil
}
