package handlers

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"net/http"
	"timdevs.rest.api.com/m/v2/error"
	"timdevs.rest.api.com/m/v2/modal"
	"timdevs.rest.api.com/m/v2/vehicle"
)

func RetrieveVehicle(c *gin.Context) {
	vin := c.Param("vin")
	retrievedVehicle := vehicle.Vehicle{}

	getVehicleResponse, getVehicleError := modal.GetVehicle(vin)
	if getVehicleError != nil {
		error.DynamoDBError(c, getVehicleError)
		return
	}

	unMarshalError := dynamodbattribute.UnmarshalMap(getVehicleResponse.Item, &retrievedVehicle)
	if unMarshalError != nil {
		error.DynamoDBError(c, unMarshalError)
		return
	}

	if retrievedVehicle.Vin == "" {
		error.NotFoundError(c)
		return
	}

	c.JSON(http.StatusOK, retrievedVehicle)
}
