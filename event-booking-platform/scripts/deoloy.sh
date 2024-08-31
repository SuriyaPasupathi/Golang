#!/bin/bash

# Script to deploy the stack to Docker Swarm

STACK_NAME=event-booking-stack

echo "Deploying stack to Docker Swarm..."
docker stack deploy -c ../deployments/swarm/swarm-stack.yaml $STACK_NAME

echo "Stack deployed successfully!"
