resource "kubernetes_namespace_v1" "this" {
  metadata {
    name = var.namespace_name
    labels = merge(
      local.common_labels,
      {
        "app.kubernetes.io/name"             = var.namespace_name
        "pod-security.kubernetes.io/enforce" = var.pod_security_level
        "pod-security.kubernetes.io/audit"   = var.pod_security_level
        "pod-security.kubernetes.io/warn"    = var.pod_security_level
      }
    )
  }
}

resource "kubernetes_resource_quota_v1" "this" {
  metadata {
    name      = "${var.namespace_name}-quota"
    namespace = kubernetes_namespace_v1.this.metadata[0].name
    labels = merge(local.common_labels, {
      "app.kubernetes.io/name"      = "${var.namespace_name}-quota"
      "app.kubernetes.io/component" = "resource-management"
    })
  }

  spec {
    hard = {
      "requests.cpu"    = var.resource_quotas.requests_cpu
      "requests.memory" = var.resource_quotas.requests_memory
      "limits.cpu"      = var.resource_quotas.limits_cpu
      "limits.memory"   = var.resource_quotas.limits_memory
      "pods"            = var.resource_quotas.max_pods
    }
  }
}
