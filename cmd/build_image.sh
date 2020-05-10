#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: <script-name> <version>"
    exit 1
fi

docker build -t shivansh-hedron:"$1" .