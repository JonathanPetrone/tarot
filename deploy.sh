#!/bin/bash

# deploy.sh - Deployment script for AI Tarot app

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🚀 AI Tarot Deployment Script${NC}"
echo "================================"

# Function to print colored output
print_step() {
    echo -e "${GREEN}✓${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠️${NC} $1"
}

print_error() {
    echo -e "${RED}❌${NC} $1"
}

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    print_error "go.mod not found. Are you in the project root directory?"
    exit 1
fi

# Parse command line arguments
COMMAND=${1:-"deploy"}
IMAGE_TAG=${2:-"latest"}

case "$COMMAND" in
    "build")
        print_step "Building Docker image..."
        docker build -t tarot-webserver:$IMAGE_TAG .
        print_step "Docker image built successfully as tarot-webserver:$IMAGE_TAG"
        ;;
    
    "deploy")
        print_step "Starting deployment to Fly.io..."
        
        # Check if fly CLI is available
        if ! command -v fly &> /dev/null; then
            print_error "Fly CLI not found. Please install it first."
            exit 1
        fi
        
        # Check if we're logged in to Fly.io
        if ! fly auth whoami &> /dev/null; then
            print_warning "Not logged in to Fly.io. Please run: fly auth login"
            exit 1
        fi
        
        print_step "Deploying to Fly.io..."
        fly deploy
        
        print_step "Deployment completed!"
        echo -e "${GREEN}🎉 Your app is live at: https://aitarot.fly.dev${NC}"
        ;;
    
    "full")
        print_step "Running full build and deploy..."
        
        print_step "1/2 Building Docker image..."
        docker build -t tarot-webserver:$IMAGE_TAG .
        
        print_step "2/2 Deploying to Fly.io..."
        fly deploy
        
        print_step "Full deployment completed!"
        echo -e "${GREEN}🎉 Your app is live at: https://aitarot.fly.dev${NC}"
        ;;
    
    "logs")
        print_step "Fetching recent logs..."
        fly logs
        ;;
    
    "status")
        print_step "Checking app status..."
        fly status
        ;;
    
    "health")
        print_step "Checking app health..."
        echo "Environment configuration:"
        curl -s https://aitarot.fly.dev/health || print_error "Health check failed"
        ;;
    
    *)
        echo "Usage: ./deploy.sh [COMMAND] [IMAGE_TAG]"
        echo ""
        echo "Commands:"
        echo "  build    - Build Docker image only"
        echo "  deploy   - Deploy to Fly.io only (default)"
        echo "  full     - Build image + deploy"
        echo "  logs     - Show recent logs"
        echo "  status   - Show app status"
        echo "  health   - Check app health endpoint"
        echo ""
        echo "Examples:"
        echo "  ./deploy.sh              # Just deploy"
        echo "  ./deploy.sh build        # Just build Docker image"
        echo "  ./deploy.sh full         # Build + deploy"
        echo "  ./deploy.sh deploy v1.2  # Deploy with specific tag"
        echo "  ./deploy.sh logs         # View logs"
        echo "  ./deploy.sh health       # Check if app is healthy"
        ;;
esac