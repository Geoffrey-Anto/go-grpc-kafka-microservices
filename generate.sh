echo "Generating Proto files"

protoc -I protos/ protos/logger.proto --go_out=plugins=grpc:protos

echo "Generated Files"
