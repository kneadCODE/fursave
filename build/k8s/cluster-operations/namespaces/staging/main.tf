module "ledgersvc" {
  source = "../../tfmodules/namespace"

  namespace_name = "ledgersvc"

  # resource_quotas = {
  #   requests_cpu    = "4"
  #   requests_memory = "8Gi"
  #   limits_cpu      = "8"
  #   limits_memory   = "16Gi"
  #   max_pods        = 20
  # }
}
