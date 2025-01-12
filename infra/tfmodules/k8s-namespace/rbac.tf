

resource "kubernetes_role_v1" "admin" {
  metadata {
    name      = "${kubernetes_namespace_v1.this.metadata[0].name}-admin-role"
    namespace = kubernetes_namespace_v1.this.metadata[0].name
    labels = merge(local.role_labels, {
      "app.kubernetes.io/name" = "${kubernetes_namespace_v1.this.metadata[0].name}-admin-role"
      "app.kubernetes.io/role" = "admin"
    })
  }

  rule {
    api_groups = ["*"]
    resources  = ["*"]
    verbs      = ["*"]
  }
}

resource "kubernetes_role_v1" "readonly" {
  metadata {
    name      = "${kubernetes_namespace_v1.this.metadata[0].name}-readonly-role"
    namespace = kubernetes_namespace_v1.this.metadata[0].name
    labels = merge(local.role_labels, {
      "app.kubernetes.io/name" = "${kubernetes_namespace_v1.this.metadata[0].name}-readonly-role"
      "app.kubernetes.io/role" = "readonly"
    })
  }

  rule {
    api_groups = ["*"]
    resources  = ["*"]
    verbs      = ["get", "list", "watch"]
  }

  rule {
    api_groups = [""]
    resources  = ["secrets"]
    verbs      = ["get", "list"]
  }
}

resource "kubernetes_role_v1" "app" {
  metadata {
    name      = "${kubernetes_namespace_v1.this.metadata[0].name}-app-role"
    namespace = kubernetes_namespace_v1.this.metadata[0].name
    labels = merge(local.role_labels, {
      "app.kubernetes.io/name" = "${kubernetes_namespace_v1.this.metadata[0].name}-app-role"
      "app.kubernetes.io/role" = "app"
    })
  }

  rule {
    api_groups = ["apps", "extensions"]
    resources  = ["deployments", "statefulsets", "daemonsets"]
    verbs      = ["get", "list", "create", "update", "patch", "delete"]
  }

  rule {
    api_groups = [""]
    resources  = ["pods", "services", "configmaps", "persistentvolumeclaims"]
    verbs      = ["get", "list", "create", "update", "patch", "delete"]
  }

  rule {
    api_groups = ["networking.k8s.io"]
    resources  = ["ingresses"]
    verbs      = ["get", "list", "create", "update", "patch", "delete"]
  }

  rule {
    api_groups = [""]
    resources  = ["secrets"]
    verbs      = ["get", "list"]
  }
}

resource "kubernetes_service_account_v1" "app" {
  metadata {
    name      = "${kubernetes_namespace_v1.this.metadata[0].name}-app-sa"
    namespace = kubernetes_namespace_v1.this.metadata[0].name
    labels = merge(local.service_account_labels, {
      "app.kubernetes.io/name" = "${kubernetes_namespace_v1.this.metadata[0].name}-app-sa"
      "app.kubernetes.io/role" = "app"
    })
  }
}

resource "kubernetes_role_binding_v1" "app" {
  metadata {
    name      = "${kubernetes_service_account_v1.app.metadata[0].name}-rb"
    namespace = kubernetes_namespace_v1.this.metadata[0].name
    labels = merge(local.common_labels, {
      "app.kubernetes.io/name"      = "${kubernetes_service_account_v1.app.metadata[0].name}-rb"
      "app.kubernetes.io/component" = "role-binding"
      "app.kubernetes.io/role"      = "app"
    })
  }

  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "Role"
    name      = kubernetes_role_v1.admin.metadata[0].name
  }

  subject {
    kind      = "ServiceAccount"
    name      = kubernetes_service_account_v1.app.metadata[0].name
    namespace = kubernetes_namespace_v1.this.metadata[0].name
  }
}
