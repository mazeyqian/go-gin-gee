#!/bin/bash

# Define usage function
usage() {
  echo "Usage: docker-run.sh [OPTIONS] IMAGE_TAG [ENV_VARS...]"
  echo "Run a Docker container from the specified IMAGE_TAG with the specified environment variables."
  echo ""
  echo "Options:"
  echo "  -h, --help    Print this help message and exit"
  echo ""
  echo "Environment variables:"
  echo "  Any additional arguments passed to the script will be passed as environment variables to the Docker container."
  echo ""
}

# Parse command-line options
while [[ $# -gt 0 ]]; do
  case $1 in
    -h|--help)
      usage
      exit 0
      ;;
    *)
      break
      ;;
  esac
done

# Check for required arguments
if [[ $# -lt 1 ]]; then
  echo "Error: IMAGE_TAG argument is required."
  usage
  exit 1
fi

# Get command-line arguments
combinedVersion=$1
shift
envVars=("$@")

# Set port numbers
visitPort="3000"
innerPort="3000"

# Print message
echo "Run Docker: ${combinedVersion} with ${envVars[@]}"

# Stop running containers
echo "Stop Docker Containers"
docker container ls --filter "name=go-gin-gee" | awk '{if (NR!=1) print $1}' | xargs docker container stop

# Remove all containers
echo "Remove Docker Containers"
docker container ls -a --filter "name=go-gin-gee" | awk '{if (NR!=1) print $1}' | xargs docker container rm

# Remove all images
# echo "Remove Docker Images"
# docker image ls --filter "reference=go-gin-gee" | awk '{if (NR!=1) print $3}' | xargs docker image rm

# Pull the specified image
echo "Pull Docker Image: ${combinedVersion}"
docker image pull ${combinedVersion}

# Build the docker run command
# dockerRunCommand="docker run --name go-gin-gee"
ENV_VARS=""
for envVar in "${envVars[@]}"; do
  ENV_VARS+=" -e ${envVar}"
done

# dockerRunCommand+=" -d -p ${visitPort}:${innerPort} ${combinedVersion}"

# Print the docker run command
# echo "Docker Run Command: ${dockerRunCommand}"

# Run the container
echo "Run Docker Container"
# eval "${dockerRunCommand}"
echo "Environment variables: $ENV_VARS"
docker run --name go-gin-gee ${ENV_VARS} -d -p ${visitPort}:${innerPort} ${combinedVersion}

# Print message
echo "Complete, Visit: http://localhost:${visitPort}/api/ping"
