#!/bin/bash
# filepath: /home/ayush/github/mindshelf/generate-certs.sh

# Create directory for certs if it doesn't exist
mkdir -p docker/nginx/certs

# Generate self-signed certificates
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout docker/nginx/certs/self-signed.key \
  -out docker/nginx/certs/self-signed.crt \
  -subj "/C=US/ST=State/L=City/O=Organization/OU=Unit/CN=localhost"

# Set proper permissions
chmod 600 docker/nginx/certs/self-signed.key
chmod 644 docker/nginx/certs/self-signed.crt

echo "Self-signed certificates generated successfully"
echo "Key: docker/nginx/certs/self-signed.key"
echo "Certificate: docker/nginx/certs/self-signed.crt"