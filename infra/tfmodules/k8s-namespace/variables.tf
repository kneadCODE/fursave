variable "namespace_name" {
  description = "Name of the Kubernetes namespace"
  type        = string
}

variable "pod_security_level" {
  description = "Pod Security Admission level"
  type        = string
  default     = "restricted"
  validation {
    condition     = contains(["privileged", "baseline", "restricted"], var.pod_security_level)
    error_message = "Pod security level must be one of: privileged, baseline, restricted"
  }
}

variable "resource_quotas" {
  description = "Resource quotas for the namespace"
  type = object({
    requests_cpu    = optional(string, "1")
    requests_memory = optional(string, "1Gi")
    limits_cpu      = optional(string, "1")
    limits_memory   = optional(string, "1Gi")
    max_pods        = optional(number, 20)
  })
  default = {}
}
