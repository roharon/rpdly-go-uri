protoc -I . -I /usr/local/include --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true ./protobuf/uri/v1/uri_exchange.proto