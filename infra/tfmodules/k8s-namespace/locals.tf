locals {
  common_labels = {
    "app.kubernetes.io/managed-by" = "terraform"
    "app.kubernetes.io/part-of"    = var.namespace_name
    "app.kubernetes.io/created-by" = "k8s-namespace-module"
  }

  role_labels = merge(local.common_labels, {
    "app.kubernetes.io/component" = "rbac"
  })

  service_account_labels = merge(local.common_labels, {
    "app.kubernetes.io/component" = "service-account"
  })
}
