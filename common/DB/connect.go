package DB

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/FinMaq/hook-reader/common/secrets"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Instance *gorm.DB

func InitDb() error {
	var err error
	sectretStr, err := secrets.GetSecretString("dataWarehouse")
	if err != nil {
		return err
	}
	secretValues := struct {
		Username string
		Password string
		Engine   string
		Host     string
		Port     int
	}{}
	err = json.Unmarshal([]byte(sectretStr), &secretValues)
	if err != nil {
		return err
	}

	dbURL := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		secretValues.Engine,
		secretValues.Username,
		secretValues.Password,
		secretValues.Host,
		secretValues.Port,
		secretValues.Username)

	Instance, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		//Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}
	log.Println("Connected to Database...")
	return nil
}
