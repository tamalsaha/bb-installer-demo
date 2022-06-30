# Install ACE

## Install Prerequisites

- Create New Google Cloud Project appscode-ace
- Add Billing account
- Gave eng@appscode.com owner access

- create buckets gs://ace-avatars, gs://ace-invoices
- TODO: create bucket for kubepack
- created new service account
- Add "Storage Object Creator" permission to the buckets

- Get token from Cloudflare for appscode.cloud Domain
- not using KMS


---
helm upgrade -i kubedb appscode/kubedb \
  --version v2022.05.24 \
  --namespace kubedb --create-namespace \
  --set kubedb-provisioner.enabled=true \
  --set kubedb-ops-manager.enabled=false \
  --set kubedb-autoscaler.enabled=false \
  --set kubedb-dashboard.enabled=false \
  --set kubedb-schema-manager.enabled=false \
  --set-file global.license=/Users/tamal/Downloads/kubedb-enterprise-license-20aae10d-67db-4041-bcdf-fe46f58d9231.txt

helm upgrade -i stash appscode/stash \
  --version v2022.05.18 \
  --namespace stash --create-namespace \
  --set features.enterprise=true \
  --set-file global.license=/Users/tamal/Downloads/kubedb-enterprise-license-20aae10d-67db-4041-bcdf-fe46f58d9231.txt

helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.8.0 \
  --set installCRDs=true

helm upgrade -i kube-prometheus-stack prometheus-community/kube-prometheus-stack \
  --namespace monitoring --create-namespace
---

## Deply ACE


---
# add helm repository appscode
helm repo add appscode https://charts.appscode.com/stable/
helm repo update
helm search repo appscode/ace
# install chart appscode/ace
helm upgrade --install ace appscode/ace \
  --namespace ace --create-namespace \
  --values=values.yaml

---