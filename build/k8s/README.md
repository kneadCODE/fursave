# Kubernetes Configuration Management

This directory contains Kubernetes configurations divided into two main categories:

## Directory Structure

- `cluster-config/`: Cluster-level configurations
  - Includes service mesh, OPA, ingress controller, etc.
  - Managed with Kustomize across environments

- `cluster-operations/`: App developer self-service configurations
  - Namespace creation
  - Network policies
  - Shared operational configurations

## Purpose

This structure isolates:
- Cluster-wide infrastructure configurations
- Namespace and application-level operational configurations

## Workflow

1. Manage cluster-wide configurations in `cluster-config/`
2. Create namespace and app-specific configurations in `cluster-operations/`

## Best Practices

- Keep cluster-level and app-level configurations separate
- Use Kustomize for environment-specific customizations
- Maintain clear separation of concerns
