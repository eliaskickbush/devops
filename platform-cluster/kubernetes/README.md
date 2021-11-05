# Kubernetes

This section of the repo describes the necessary steps to get a local Kubernetes cluster running.

# Prerequisites

You should install `kubectl`, and `curl`.

# Steps

This kubernetes cluster, being a local one, is ran using K3s:

```
# Install K3s (Idempotent)

curl -sfL https://get.k3s.io | sh -
```
Once K3s is installed, you could benefit from registering the following alias:

```
# Source k3sctl alias (Idempotent)
sudo cp -n /etc/rancher/k3s/k3s.yaml ~/.kube
sudo chmod 444 ~/.kube/k3s.yaml
command=$(echo "kubectl --kubeconfig='$HOME/.kube/k3s.yaml'")
grep "alias k3sctl=\"$command\"" ~/.zshrc 1&>/dev/null || echo "alias k3sctl=\"$command\"" >> ~/.zshrc
```

Now we should have a running cluster! Confirm with:

```
k3sctl version
```

Which should return the client and server versions.

Hurray! Now you can begin deploying software.