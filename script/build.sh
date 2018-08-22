source ./bootstrap.sh

# Docker build and tag image
docker build -t ${IMAGE_NAME}:${TAG} .
docker tag ${IMAGE_NAME}:${TAG} ${HOST_NAME}/${PROJECT_ID}/${IMAGE_NAME}:${TAG}