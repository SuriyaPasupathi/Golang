# Event Booking Platform

This repository contains a proof of concept (POC) for an event booking platform using Go, gRPC, and a gRPC REST Gateway. The platform is containerized using Docker and deployed using Docker Swarm with Traefik as a reverse proxy.

## Structure
- `api/`: Contains protobuf definitions and generated code.
- `cmd/`: Main entry points for the services.
- `internal/`: Service-specific and shared application logic.
- `deployments/`: Docker and Swarm deployment configurations.
- `pkg/`: Reusable packages and clients.
- `scripts/`: Utility scripts for building, deploying, and generating code.
- `Makefile`: Simplifies common tasks like building, running, and deploying.

## Setup

1. **Build Services**
   ```sh
   make build
