echo "Generating Proto files"

protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
protos/logger.proto

protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
protos/randomjoke.proto

rm -rf protos/randomjoke
rm -rf protos/logger

mkdir protos/randomjoke
mkdir protos/logger

mv protos/logger*.pb.go protos/logger
mv protos/randomjoke*.pb.go protos/randomjoke

echo "Generated Files"
