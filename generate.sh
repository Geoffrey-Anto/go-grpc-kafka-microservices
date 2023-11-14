echo "Generating Proto files"

protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
protos/randomjoke.proto

rm -rf protos/randomjoke

mkdir protos/randomjoke

mv protos/randomjoke*.pb.go protos/randomjoke

echo "Generated Files"
