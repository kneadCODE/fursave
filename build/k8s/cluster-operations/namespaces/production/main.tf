module "ledgersvc" {
  source = "../../tfmodules/namespace"

  namespace_name = "ledgersvc"

  # resource_quotas = {
  #   requests_cpu    = "8"
  #   requests_memory = "16Gi"
  #   limits_cpu      = "16"
  #   limits_memory   = "32Gi"
  #   max_pods        = 50
  # }
}
