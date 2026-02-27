#!/bin/bash

# Build script for nanotalon

echo "Building nanotalon..."

# Create bin directory if it doesn't exist
mkdir -p bin

# Build the main executable
go build -o bin/nanotalon ./cmd/main.go

if [ $? -eq 0 ]; then
    echo "✅ Successfully built nanotalon"
    echo "Run './bin/nanotalon --help' to get started"
else
    echo "❌ Build failed"
    exit 1
fi