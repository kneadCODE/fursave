# Cluster Operations Configurations

## Purpose

This directory contains configurations for app developers to self-service their namespace and application-level configurations.

## Supported Configurations

- Namespace creation
- Network policies
- Resource quotas
- Service accounts
- Basic RBAC configurations

## Usage Guidelines

- Keep configurations minimal and focused
- Follow least-privilege principle
- Use templates for consistency
- Avoid cluster-wide configurations (those belong in `cluster-config/`)

## Recommended Workflow

1. Create namespace configuration
2. Define network policies
3. Set up service accounts
4. Configure resource quotas
5. Apply RBAC rules as needed
