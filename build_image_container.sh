IMAGE_NAME="ascii-art-web"
CONTAINER_NAME="ascii-art-web"

echo "Building the Docker image..."
docker image build -f Dockerfile -t $IMAGE_NAME .

echo "Running the Docker container..."
docker container run -p 8080:8080 --detach --name $CONTAINER_NAME $IMAGE_NAME

echo "Container is running. Access the app at http://localhost:8080"

docker images
docker ps -a