module github.com/ngunq/shippy/shippy-cli-consignment

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/micro/go-micro v1.18.0
	github.com/ngunq/shippy/shippy-service-consignment v0.0.0-20200908095726-39c7565d4e1c
)
