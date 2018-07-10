export IMAGE_NAME=mangohacks2019-api

# Extract information for tagging
export COMMIT=${TRAVIS_COMMIT::8}
export BRANCH=$(if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then echo $TRAVIS_BRANCH; else echo $TRAVIS_PULL_REQUEST_BRANCH; fi)
echo "TRAVIS_BRANCH=$TRAVIS_BRANCH, PR=$PR, BRANCH=$BRANCH"

docker build -t ${IMAGE_NAME}:${BRANCH}-${COMMIT} .