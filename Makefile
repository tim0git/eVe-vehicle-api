#Table commands
start_db:
	docker run --detach --name dynamodb -p 8000:8000 amazon/dynamodb-local:1.22.0 -jar DynamoDBLocal.jar -sharedDb -inMemory

delete_db:
	docker rm -f dynamodb

create_table:
	aws dynamodb create-table --table-name Vehicles --attribute-definitions AttributeName=vin,AttributeType=S --key-schema AttributeName=vin,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 --endpoint-url http://localhost:8000

#Testing commands
test: start_db
	go test -count=1 --tags=docker ./... ; make delete_db

debug: start_db
	go test -count=1 --tags=docker ./... -v ; make delete_db

coverage: start_db
	go test -coverprofile=coverage.out ./... ; go tool cover -html=coverage.out ; make delete_db

integration: compose_up
	make newman ; make compose_down ; make clean

newman:
	docker run --rm --network host -t postman/newman run https://api.postman.com/collections/12717734-83accfac-9b69-47e1-8cd8-59e7179da108?access_key=${POSTMAN_COLLECTION_API_KEY}

#Go binary commands
dev:
	PORT="8443" TABLE_NAME="Vehicles" AWS_ACCESS_KEY_ID="fakeKey" AWS_SECRET_ACCESS_KEY="fakeSecretAccessKey" DYNAMODB_ENDPOINT="http://localhost:8000" go run main.go

build:
	go build -ldflags="-w -s" -o build

run:
	PORT="8443" TABLE_NAME="Vehicles" AWS_ACCESS_KEY_ID="fakeKey" AWS_SECRET_ACCESS_KEY="fakeSecretAccessKey" DYNAMODB_ENDPOINT="http://localhost:8000" ./build

#Docker commands
package:
	docker build -t vehicles-api -f Dockerfile .

run_package:
	docker run --add-host localhost:host-gateway -p8443:8443 --env-file .env.local vehicles-api

dive:
	CI=true dive vehicles-api --ci-config docker/.dive.yaml

compose_up:
	docker-compose up -d && aws dynamodb create-table --table-name Vehicles --attribute-definitions AttributeName=vin,AttributeType=S --key-schema AttributeName=vin,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 --endpoint-url http://localhost:8000 --region us-east-1 2>&1 > /dev/null

compose_down:
	docker-compose down

clean:
	rm -rf build ; docker rm -f dynamodb ; docker rmi timdevs-go-rest-api-vehicle-api:latest ; docker rmi vehicles-api:latest ; docker-compose down ; docker image prune -f ; docker builder prune -f;

#Documentation commands
generate_swagger_docs:
	swag init
