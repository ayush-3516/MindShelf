#!/bin/bash
# filepath: /home/ayush/github/mindshelf/mindshelf.sh

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to display usage
function show_help {
    echo -e "${BLUE}Mindshelf Development Helper${NC}"
    echo "Usage: $0 [command]"
    echo ""
    echo "Commands:"
    echo "  start           Start the development environment"
    echo "  stop            Stop the development environment"
    echo "  restart         Restart the development environment"
    echo "  logs            Show logs from containers"
    echo "  build           Rebuild containers"
    echo "  clean           Remove all containers and volumes"
    echo "  certs           Generate SSL certificates"
    echo "  help            Show this help message"
}

# Generate certificates
function generate_certs {
    echo -e "${BLUE}Generating SSL certificates...${NC}"
    bash ./generate-certs.sh
}

# Start development environment
function start {
    if [ ! -f "docker/nginx/certs/self-signed.crt" ] || [ ! -f "docker/nginx/certs/self-signed.key" ]; then
        generate_certs
    fi
    
    echo -e "${BLUE}Starting development environment...${NC}"
    docker-compose up -d
    
    echo -e "${GREEN}Mindshelf is running!${NC}"
    echo -e "API is available at ${BLUE}http://localhost:8080/api${NC}"
    echo -e "Frontend is available at ${BLUE}http://localhost${NC}"
}

# Stop development environment
function stop {
    echo -e "${BLUE}Stopping development environment...${NC}"
    docker-compose down
    echo -e "${GREEN}Stopped successfully!${NC}"
}

# Show logs
function show_logs {
    echo -e "${BLUE}Showing logs (Ctrl+C to stop)...${NC}"
    docker-compose logs -f
}

# Rebuild containers
function rebuild {
    echo -e "${BLUE}Rebuilding containers...${NC}"
    docker-compose down
    docker-compose build
    docker-compose up -d
    echo -e "${GREEN}Rebuild completed!${NC}"
}

# Clean everything
function clean {
    echo -e "${RED}Warning: This will remove all containers and volumes!${NC}"
    read -p "Are you sure you want to continue? (y/N) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo -e "${BLUE}Cleaning up...${NC}"
        docker-compose down -v
        echo -e "${GREEN}Cleanup completed!${NC}"
    fi
}

# Process command
case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        stop
        start
        ;;
    logs)
        show_logs
        ;;
    build)
        rebuild
        ;;
    clean)
        clean
        ;;
    certs)
        generate_certs
        ;;
    help|*)
        show_help
        ;;
esac