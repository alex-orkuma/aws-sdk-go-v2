module github.com/aws/aws-sdk-go-v2/service/dynamodb/internal/benchmark

go 1.15

require (
	github.com/aws/aws-sdk-go v1.34.15
	github.com/aws/aws-sdk-go-v2 v0.0.0-20200923184724-5013332bb7ed
	github.com/aws/aws-sdk-go-v2/service/dynamodb v0.0.0-00010101000000-000000000000
)

replace github.com/aws/aws-sdk-go-v2 => ../../../../

replace github.com/aws/aws-sdk-go-v2/service/dynamodb => ../../
