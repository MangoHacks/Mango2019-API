# Set env vars
export IMAGE_NAME=mangohacks2019-api
export COMMIT=${TRAVIS_COMMIT::8}
export BRANCH=$(if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then echo $TRAVIS_BRANCH; else echo $TRAVIS_PULL_REQUEST_BRANCH; fi)
echo "TRAVIS_BRANCH=$TRAVIS_BRANCH, PR=$PR, BRANCH=$BRANCH"

# Docker build and tag image
docker build -t ${IMAGE_NAME}:${BRANCH}-${COMMIT} .
docker tag ${IMAGE_NAME}:${BRANCH}-${COMMIT} ${HOST_NAME}/${PROJECT_ID}/${IMAGE_NAME}:${BRANCH}-${COMMIT}