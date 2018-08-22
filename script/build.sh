this_dir=`dirname $0`

source $this_dir/bootstrap.sh

# Docker build and tag image
docker build -t ${IMAGE_NAME}:${TAG} .
docker tag ${IMAGE_NAME}:${TAG} ${HOST_NAME}/${PROJECT_ID}/${IMAGE_NAME}:${TAG}