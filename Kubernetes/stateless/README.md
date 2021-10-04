# Stateless App

This stateless application can be ran as a single pod without the need of any secrets, configuration values, or additional dependencies.

The first step towards using this application is building the image and pushing it to a private registry.

`docker build -ti stateless .`

This will tag the image with the `latest` tag (you may want to override this and use your own tags).

After your image is built, you need to authenticate to the remote image repository and push the newly created image.

> Note: We will use my personal account for this, you may need to replace the endpoints for your repository

For that we need to:

1. Authenticate to the remote repository: `aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 297813884257.dkr.ecr.us-east-1.amazonaws.com` 
2. Tag the latest image to have the ECR remote: `docker tag stateless:latest 297813884257.dkr.ecr.us-east-1.amazonaws.com/stateless:latest`
3. Push the image: `docker push 297813884257.dkr.ecr.us-east-1.amazonaws.com/stateless:latest`

After we have our image living in the remote repository, now we need to configure our kubernetes cluster with the necessary credentials to pull that image.

In order to do this, we will create a secret in the K8s cluster, with an existing command for that: `kubectl create secret docker-registry`

> Note: Although the command is named `docker-registry`, you can configure the kubernetes cluster to pull from all sorts of registries. The main ones being: Google Container Registry, AWS Elastic Container Registry, Harbor, Azure Container Registry, etc.

The following script creates the secret, given valid AWS credentials to access the ECR, and [`jq`](https://stedolan.github.io/jq/).

```
#!/bin/bash
# Get repository URI
REPO_NAME="https://$(aws ecr describe-repositories | jq -r '.repositories[] | select(.repositoryName == "stateless") | .repositoryUri')"

# Get repository credentials (In case of ECR, only a password is required)
REPO_PASSWORD=$(aws ecr get-login-password --region us-east-1)

# Create secret 
kubectl --kubeconfig=~/.kube/k3s.yml create secret docker-registry stateless-registry --docker-server="https://$REPO_NAME" --docker-username=AWS --docker-password="$REPO_PASSWORD"
```