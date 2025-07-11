#!/bin/bash

# Set root folder name
ROOT_DIR="."

# Create directories
mkdir -p $ROOT_DIR/{config,models,routes,controllers,services,utils,middleware}

# Create files
touch $ROOT_DIR/main.go

touch $ROOT_DIR/config/db.go

touch $ROOT_DIR/models/invoice.go
touch $ROOT_DIR/models/user.go

touch $ROOT_DIR/routes/router.go

touch $ROOT_DIR/controllers/invoice.go
touch $ROOT_DIR/controllers/user.go

touch $ROOT_DIR/services/payment.go

touch $ROOT_DIR/utils/webhooks.go

touch $ROOT_DIR/middleware/auth.go

echo "✅ Golang project structure created at: $ROOT_DIR"
