swagger-format:
	swag fmt

generate-swagger-full: swagger-format
	swag init \
	--instanceName full \
	--generalInfo pkg/api/swagger/full.go \
	--outputTypes go \
	--quiet \
	--dir ./

generate: generate-swagger-full
