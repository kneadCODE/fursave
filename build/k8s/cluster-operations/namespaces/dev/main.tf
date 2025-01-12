module "ledgersvc" {
  source = "../../../../../infra/tfmodules/k8s-namespace"

  namespace_name = "ledgersvc"
}
