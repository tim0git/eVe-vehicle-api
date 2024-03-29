package modal

import (
	"eve.vehicle.api.com/m/v2/database"
	"eve.vehicle.api.com/m/v2/utils"
	"eve.vehicle.api.com/m/v2/vehicle"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func UpdateVehicle(vehicle vehicle.Update, vin string) (*dynamodb.UpdateItemOutput, error) {
	client := database.DynamoDB()

	updateRequest := &dynamodb.UpdateItemInput{
		TableName: aws.String(utils.GetTableName()),
		Key: map[string]*dynamodb.AttributeValue{
			"vin": {
				S: aws.String(vin),
			},
		},
		ConditionExpression:       aws.String("attribute_exists(vin)"),
		UpdateExpression:          aws.String(getUpdateExpression()),
		ExpressionAttributeValues: getUpdateExpressionAttributeValues(vehicle),
		ExpressionAttributeNames:  getUpdateExpressionAttributeNames(),
	}

	return client.UpdateItem(updateRequest)
}

func getUpdateExpression() string {
	return "set #manufacturer = :manufacturer, #model = :model, #year = :year, #color = :color, #batteryCapacity = :batteryCapacity"
}

func getUpdateExpressionAttributeValues(vehicle vehicle.Update) map[string]*dynamodb.AttributeValue {
	expressionAttributeValues := map[string]*dynamodb.AttributeValue{
		":manufacturer": {
			S: aws.String(vehicle.Manufacturer),
		},
		":model": {
			S: aws.String(vehicle.Model),
		},
		":year": {
			N: aws.String(fmt.Sprintf("%d", vehicle.Year)),
		},
		":color": {
			S: aws.String(vehicle.Color),
		},
		":batteryCapacity": {
			M: map[string]*dynamodb.AttributeValue{
				"value": {
					N: aws.String(fmt.Sprintf("%d", vehicle.Capacity.Value)),
				},
				"unit": {
					S: aws.String(fmt.Sprintf("%s", vehicle.Capacity.Unit)),
				},
			},
		},
	}
	return expressionAttributeValues
}

func getUpdateExpressionAttributeNames() map[string]*string {
	expressionAttributeNames := map[string]*string{
		"#model":           aws.String("model"),
		"#year":            aws.String("year"),
		"#color":           aws.String("color"),
		"#batteryCapacity": aws.String("capacity"),
		"#manufacturer":    aws.String("manufacturer"),
	}
	return expressionAttributeNames
}
