#!/bin/bash

which protoc >/dev/null
if [ $? -ne 0 ]; then
  echo "Protoc not installed"
  exit 1
fi



protoc -I api/proto --go_opt=paths=source_relative --go_out=plugins=grpc:pkg/api api/proto/*.proto