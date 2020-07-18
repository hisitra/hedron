#!/bin/bash

# SECTION 1: Utility Methods.
if [[ "$#" -ne 1 ]]; then
  echo "Usage: <script-name> <tag>"
  exit 1
fi

# SECTION 2: Image Building.
# Comment this section if you don't want to build an image.
# In that case, the older image will be used.
echo "BUILDING DOCKER IMAGE..."
if ! docker build -t hedron-image:"$1" .; then
  echo "FAILED TO BUILD IMAGE!"
  exit 1
fi
echo "IMAGE SUCCESSFULLY BUILT."
echo "======================================="

# SECTION 3: Removing the older containers, or volumes.
echo "REMOVING OLD CONTAINER..."
docker rm -f hedron-container

# SECTION 4: Running the new container.
# You'll have to change the port here (see the publish option)
# if you have changed it in the configs.
echo "RUNNING NEW CONTAINER..."
if ! docker run -d \
  --name hedron-container \
  --publish 3000:3000 \
  hedron-image:"$1"; then
  echo "Failed to run container"
  exit 1
fi
echo "CONTAINER UP AND RUNNING."
echo "======================================="