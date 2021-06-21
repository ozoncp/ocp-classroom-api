# Ozon Code Platform Classroom API 

by Aleksandr Kuzminykh

Is a gRPC sevice that provides access to operate with classrooms.

It has logging, metrics and tracing.

## Run
```cmd
docker-compose up
```

## [Prometheus UI](http://localhost:9090)

## [Jaeger UI](http://localhost:16686)

## [Swagger UI](http://localhost:80/swagger)

### Run
```cmd
docker run -d -p 80:8080 -e BASE_URL=/swagger -e SWAGGER_JSON=/swagger/api.swagger.json -v %cd%/swagger:/swagger swaggerapi/swagger-ui
```

## gRPC client
Is a simple user-interactive gRPC client to test service

### Run locally
```cmd
go run .\cmd\grpc-client\
```

## Kafka Apache consumer
Is a simple Kafka consumer to test service's logging

### Run locally
```cmd
go run .\cmd\kafka-consumer\
```