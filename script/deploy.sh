# Set env vars
export IMAGE_NAME=mangohacks2019-api
export COMMIT=${TRAVIS_COMMIT::8}
export BRANCH=$(if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then echo $TRAVIS_BRANCH; else echo $TRAVIS_PULL_REQUEST_BRANCH; fi)
echo "TRAVIS_BRANCH=$TRAVIS_BRANCH, PR=$PR, BRANCH=$BRANCH"

# Decrypt our keys
openssl aes-256-cbc -K $encrypted_f24c28559e81_key -iv $encrypted_f24c28559e81_iv -in client-secret.json.enc -out client-secret.json -d

# If the SDK is not already cached, download it and unpack it.
gcloud version || true
if [ ! -d "$HOME/google-cloud-sdk/bin" ]; then rm -rf $HOME/google-cloud-sdk; export CLOUDSDK_CORE_DISABLE_PROMPTS=1; curl https://sdk.cloud.google.com | bash; fi

# Add gcloud to $PATH
source /home/travis/google-cloud-sdk/path.bash.inc
gcloud version

# Here we use the decrypted service account credentials to authenticate the command line tool.
gcloud auth activate-service-account --key-file client-secret.json

# Push to Docker registry.
gcloud docker -- push ${HOST_NAME}/${PROJECT_ID}/${IMAGE_NAME}:${BRANCH}-${COMMIT}

# Install Kubernetes.
gcloud components install kubectl

# Decrypt and set Kubeconfig
openssl aes-256-cbc -K $encrypted_f24c28559e81_key -iv $encrypted_f24c28559e81_iv -in config.enc -out config -d
mkdir $HOME/.kube
mv config $HOME/.kube/config

# Update on Kubernetes.
kubectl set image deployment/mangohacks2019-api mangohacks2019-api=${HOST_NAME}/${PROJECT_ID}/${IMAGE_NAME}:${BRANCH}-${COMMIT}
