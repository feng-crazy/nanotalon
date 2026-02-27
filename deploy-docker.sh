#!/bin/bash

# nanotalon Docker Deployment Script

set -e

echo "_nanotalon Docker Deployment Script_"
echo

# Function to show usage
usage() {
    echo "Usage: $0 [build|run|up|down|logs|shell|help]"
    echo
    echo "Commands:"
    echo "  build    - Build the Docker image"
    echo "  run      - Run the container in detached mode"
    echo "  up       - Start the service using docker-compose"
    echo "  down     - Stop the service using docker-compose"
    echo "  logs     - Show container logs"
    echo "  shell    - Open a shell in the running container"
    echo "  help     - Show this help message"
    echo
    exit 1
}

# Check if docker is installed
if ! command -v docker &> /dev/null; then
    echo "Error: Docker is not installed. Please install Docker first."
    exit 1
fi

# Check if docker-compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo "Warning: docker-compose is not installed. Some commands may not work."
fi

case "${1:-help}" in
    build)
        echo "Building nanotalon Docker image..."
        docker build -t nanotalon .
        echo "✓ Docker image built successfully!"
        ;;
    run)
        echo "Running nanotalon container..."
        docker run -d \
            --name nanotalon \
            -p 18790:18790 \
            -v ~/.nanotalon:/home/nonroot/.nanotalon \
            -v ~/nanotalon-workspace:/workspace \
            nanotalon
        echo "✓ nanotalon container is running!"
        echo "✓ Access the service at http://localhost:18790"
        ;;
    up)
        if command -v docker-compose &> /dev/null; then
            echo "Starting nanotalon service with docker-compose..."
            docker-compose up -d
            echo "✓ nanotalon service started!"
            echo "✓ Access the service at http://localhost:18790"
        else
            echo "Error: docker-compose is not installed."
            exit 1
        fi
        ;;
    down)
        if command -v docker-compose &> /dev/null; then
            echo "Stopping nanotalon service..."
            docker-compose down
            echo "✓ nanotalon service stopped!"
        else
            echo "Stopping nanotalon container..."
            docker stop nanotalon 2>/dev/null || true
            docker rm nanotalon 2>/dev/null || true
            echo "✓ nanotalon container stopped!"
        fi
        ;;
    logs)
        if command -v docker-compose &> /dev/null; then
            docker-compose logs -f nanotalon
        else
            docker logs -f nanotalon
        fi
        ;;
    shell)
        if command -v docker-compose &> /dev/null; then
            docker-compose exec nanotalon sh
        else
            docker exec -it nanotalon sh
        fi
        ;;
    help|"")
        usage
        ;;
    *)
        echo "Error: Unknown command '$1'"
        echo
        usage
        ;;
esac