# Decrypt our keys
openssl aes-256-cbc -K $encrypted_f24c28559e81_key -iv $encrypted_f24c28559e81_iv -in client-secret.json.enc -out client-secret.json -d

# If the SDK is not already cached, download it and unpack it.
gcloud version || true
if [ ! -d "$HOME/google-cloud-sdk/bin" ]; then rm -rf $HOME/google-cloud-sdk; export CLOUDSDK_CORE_DISABLE_PROMPTS=1; curl https://sdk.cloud.google.com | bash; fi

# Add gcloud to $PATH
source /home/travis/google-cloud-sdk/path.bash.inc
gcloud version

tar -xzf credentials.tar.gz
mkdir -p lib

# Here we use the decrypted service account credentials to authenticate the command line tool.
gcloud auth activate-service-account --key-file client-secret.json

# Install Kubernetes.
curl -Lo kubectl https://storage.googleapis.com/kubernetes-release/release/v1.7.0/bin/linux/amd64/kubectl && \
    chmod +x kubectl && sudo mv kubectl /usr/local/bin/

# Push to Docker registry.
gcloud docker -- push ${HOST_NAME}/${PROJECT_ID}/MangoHacks2019-API:v1

# Run on Kubernetes.
kubectl run MangoHacks2019-API --image=${HOST_NAME}/${PROJECT_ID}/MangoHacks2019-API:v1 --port 9000
kubectl expose deployment hello-web --type=LoadBalancer --port 80 --target-port 9000