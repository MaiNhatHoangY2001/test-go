#!/bin/bash
# Consolidation script - removes duplicate directories
# This script removes the old duplicate architecture directories
# as part of the code consolidation process

echo "=========================================="
echo "Todo App Code Consolidation Script"
echo "=========================================="
echo ""
echo "This script will remove duplicate code directories:"
echo "  - internal/api/"
echo "  - internal/application/"
echo "  - internal/domain/"
echo ""
echo "The canonical feature-based structure will be preserved:"
echo "  - internal/features/"
echo "  - internal/shared/"
echo "  - internal/infrastructure/"
echo ""
read -p "Continue? (y/n) " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Consolidation cancelled."
    exit 1
fi

echo ""
echo "Backing up directories..."
BACKUP_DIR="backup_$(date +%Y%m%d_%H%M%S)"
mkdir -p "$BACKUP_DIR"
cp -r internal/api "$BACKUP_DIR/"
cp -r internal/application "$BACKUP_DIR/"
cp -r internal/domain "$BACKUP_DIR/"
echo "Backup created in: $BACKUP_DIR"

echo ""
echo "Removing duplicate directories..."
rm -rf internal/api
rm -rf internal/application
rm -rf internal/domain
echo "Duplicate directories removed."

echo ""
echo "=========================================="
echo "Consolidation Complete!"
echo "=========================================="
echo ""
echo "All code is now centralized in:"
echo "  - internal/features/auth"
echo "  - internal/features/todo"
echo "  - internal/shared"
echo "  - internal/infrastructure"
echo ""
echo "Next steps:"
echo "1. Run: make deps"
echo "2. Run: make test"
echo "3. Run: make build"
echo ""
