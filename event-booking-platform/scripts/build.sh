#!/bin/bash

# Script to build all services

echo "Building Event Service..."
go build -o ../bin/event-service ../cmd/event-service

echo "Building Booking Service..."
go build -o ../bin/booking-service ../cmd/booking-service

echo "Building Gateway Service..."
go build -o ../bin/gateway-service ../cmd/gateway-service

echo "Build completed successfully!"
