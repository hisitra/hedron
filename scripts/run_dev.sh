#!/bin/sh

echo "Sourcing the development configs..."
set -a
if ! . scripts/dev_configs.env; then
  set +a
  echo "Failed to source development configs"
  exit 1
fi
set +a
echo "Configs sourced."
echo "======================================="

echo "Building native binary..."
if ! make build --quiet; then
  echo "Failed to build native binary"
  exit 1
fi
echo "Binary built."
echo "======================================="

echo "Running binary..."
if ! bin/hedron-bin; then
  echo "Application binary exited with non-zero status code"
  exit 1
fi