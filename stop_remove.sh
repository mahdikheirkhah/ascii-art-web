#!/bin/bash

# this will stop the container if it is running and then remove it and then remove the image.

# Define variables for the container and image names
CONTAINER_NAME="ascii-art-web"
IMAGE_NAME="ascii-art-web"

echo "Stopping the container if it is running..."
if [ $(docker ps -q -f name=$CONTAINER_NAME) ]; then
    docker container stop $CONTAINER_NAME
    echo "Container $CONTAINER_NAME stopped."
else
    echo "No running container with name $CONTAINER_NAME."
fi

echo "Removing the container if it exists..."
if [ $(docker ps -aq -f name=$CONTAINER_NAME) ]; then
    docker container rm $CONTAINER_NAME
    echo "Container $CONTAINER_NAME removed."
else
    echo "No container with name $CONTAINER_NAME to remove."
fi

echo "Removing the image if it exists..."
if [ $(docker images -q $IMAGE_NAME) ]; then
    docker image rm $IMAGE_NAME
    echo "Image $IMAGE_NAME removed."
else
    echo "No image with name $IMAGE_NAME to remove."
fi

echo "Cleanup complete!"