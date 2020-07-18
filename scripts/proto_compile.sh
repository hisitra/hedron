#!/bin/bash

PROTO_DIR="src/protos"

echo "======================================="
echo "Compiling proto files..."
if ! protoc -I $PROTO_DIR/ $PROTO_DIR/*.proto --go_out=plugins=grpc:$PROTO_DIR/; then
  echo "Failed to compile proto files."
  exit 1
fi
echo "Proto files successfully compiled."
echo "======================================="