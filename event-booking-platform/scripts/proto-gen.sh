#!/bin/bash

# Script to generate Go code from protobuf files

PROTO_DIR=../api/proto
OUT_DIR=../api/proto

echo "Generating Go files from protobuf definitions..."

protoc --proto_path=$PROTO_DIR \
  --go_out=$OUT_DIR --go_opt=paths=source_relative \
  --go-grpc_out=$OUT_DIR --go-grpc_opt=paths=source_relative \
  $PROTO_DIR/event/*.proto \
  $PROTO_DIR/booking/*.proto \
  $PROTO_DIR/gateway/*.proto

echo "Protobuf generation completed!"
