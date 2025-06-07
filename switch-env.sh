#!/bin/bash

# switch-env.sh - Switch between local and production environments

case "$1" in
    "local")
        cp .env.local .env
        echo "🏠 Switched to LOCAL environment"
        echo "   Database: localhost (your local PostgreSQL)"
        ;;
    "prod"|"production")
        cp .env.production .env
        echo "☁️  Switched to PRODUCTION environment"
        echo "   Database: Supabase (cloud)"
        echo "   ⚠️  You're now connected to LIVE data!"
        ;;
    "show"|"current")
        if [ -f .env ]; then
            ENV_TYPE=$(grep "APP_ENV=" .env | cut -d'=' -f2)
            DB_HOST=$(grep "DB_HOST=" .env | cut -d'=' -f2)
            echo "Current environment: $ENV_TYPE"
            echo "Database host: $DB_HOST"
        else
            echo "No .env file found"
        fi
        ;;
    *)
        echo "Usage: ./switch-env.sh [local|prod|show]"
        echo ""
        echo "Commands:"
        echo "  local - Switch to local PostgreSQL database"
        echo "  prod  - Switch to production Supabase database"
        echo "  show  - Show current environment"
        echo ""
        echo "Example:"
        echo "  ./switch-env.sh local"
        echo "  go run cmd/test/main.go"
        echo "  ./switch-env.sh prod"
        echo "  go run cmd/test/main.go"
        ;;
esac