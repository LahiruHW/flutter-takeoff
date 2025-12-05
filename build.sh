#!/bin/bash
# Build script for Flutter Takeoff
# Sets version information at build time

set -e

OUTPUT_NAME="${1:-flutter-installer}"

# Get version info
VERSION="1.0.0"
BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "unknown")

echo -e "\033[36mBuilding Flutter Takeoff v${VERSION}\033[0m"
echo -e "\033[90m  Build Date: ${BUILD_DATE}\033[0m"
echo -e "\033[90m  Git Commit: ${GIT_COMMIT}\033[0m"
echo -e "\033[90m  Git Branch: ${GIT_BRANCH}\033[0m"
echo ""

# Build with version information
LDFLAGS="-X 'flutter_takeoff/pkg/version.BuildDate=${BUILD_DATE}' \
         -X 'flutter_takeoff/pkg/version.GitCommit=${GIT_COMMIT}' \
         -X 'flutter_takeoff/pkg/version.GitBranch=${GIT_BRANCH}'"

go build -ldflags "$LDFLAGS" -o "${OUTPUT_NAME}" .

echo -e "\033[32m✓ Build successful: ${OUTPUT_NAME}\033[0m"
echo ""
echo -e "\033[33mTo run: ./${OUTPUT_NAME}\033[0m"
