#!/bin/bash
# Get repository URI
REPO_NAME="https://$(aws ecr describe-repositories | jq -r '.repositories[] | select(.repositoryName == "stateless") | .repositoryUri')"

# Get repository credentials (In case of ECR, only a password is required)
REPO_PASSWORD=$(aws ecr get-login-password --region us-east-1)

# Create secret 
kubectl --kubeconfig ~/.kube/k3s.yml create secret docker-registry stateless-registry --docker-server="https://$REPO_NAME" --docker-username=AWS --docker-password="$REPO_PASSWORD"
