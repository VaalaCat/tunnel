#!/bin/bash
while true; do cat request | nc localhost 7002; done