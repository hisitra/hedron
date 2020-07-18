#!/bin/sh

echo "Sourcing the production configs..."
set -a
if ! . scripts/prod_configs.env; then
  set +a
  echo "Failed to source production configs"
  exit 1
fi
set +a
echo "Configs sourced."
echo "======================================="

echo "Building native binary..."
if ! go build -o ./bin/hedron-bin ./main.go; then
  echo "Application binary exited with non-zero status code"
  exit 1
fi
echo "Binary built."
echo "======================================="

echo "Running binary..."
if ! bin/hedron-bin; then
  echo "Application binary exited with non-zero status code"
  exit 1
fi