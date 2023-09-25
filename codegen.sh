#!/bin/bash

protoc --go_out=plugins=grpc:. ./idl/*.proto -I ./