#!/bin/bash
# filepath: /home/ayush/github/mindshelf/deploy-prod.sh

# Exit on error
set -e

echo "Starting Mindshelf deployment..."

# Check if .env file exists
if [ ! -f ".env" ]; then
    echo "Error: .env file not found. Please create one with the required environment variables."
    exit 1
fi

# Generate certificates if they don't exist
if [ ! -f "docker/nginx/certs/self-signed.crt" ] || [ ! -f "docker/nginx/certs/self-signed.key" ]; then
    echo "Generating SSL certificates..."
    bash ./generate-certs.sh
fi

# Create logs directory if it doesn't exist
mkdir -p docker/nginx/logs

# Build and start the containers
echo "Building and starting Docker containers..."
docker-compose -f docker-compose.prod.yml down
docker-compose -f docker-compose.prod.yml build
docker-compose -f docker-compose.prod.yml up -d

echo "Deployment completed successfully!"
echo "Your application should be running at https://localhost"