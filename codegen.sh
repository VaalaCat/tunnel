#!/bin/bash

/usr/local/bin/protoc --go_out=plugins=grpc:. ./idl/*.proto -I ./