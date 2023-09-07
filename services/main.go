package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
	"https://github.com/muhammadali07/rest-api-go/services/api"
)

func main() {

	REST_SERVER_ADDRESS := fmt.Sprintf("%s:%d", viper.GetString("REST_SERVER_HOST"), viper.GetInt("REST_SERVER_PORT"))

	restApp := api.NewRESTGatewaysBCASAPP(logger, REST_SERVER_ADDRESS)
	restApp.SetupRoutes()

	if err := restApp.Start(REST_SERVER_ADDRESS); err != nil {
		remark := "There was a problem, please contact customer support (IS01)"
		logger.Error(logrus.Fields{
			"error": err,
		}, "nil", remark)
	}
}
