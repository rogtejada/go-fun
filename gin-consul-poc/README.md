###Simple Consul integration with gin


Start consul

``
docker run --net=host consul:latest consul agent -dev -bind=127.0.0.1
``

Then run 

``
go run main.go
``

The service will be listening on HttpPort 8080, and sending heartbeat to consul

You can fetch consul info from this service by calling

``
curl 127.0.0.1:8500/v1/health/service/my-service
``

