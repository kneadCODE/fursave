package cfg

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// NewOTELResourceFromEnv creates a new OpenTelemetry resource from environment variables.
// It loads various resource attributes from the environment, including service, container,
// Kubernetes, and cloud resources. It also sets the deployment environment attribute.
//
// Parameters:
//   - ctx: The context for resource creation.
//
// Returns:
//   - res: The created OpenTelemetry resource.
//   - env: The deployment environment value.
//   - err: An error if any occurred during resource creation or if required environment variables are missing.
func NewOTELResourceFromEnv(ctx context.Context) (res *resource.Resource, env string, err error) {
	attrs, err := loadServiceResourceFromEnv()
	if err != nil {
		return
	}

	env = getOTELEnvVar(semconv.DeploymentEnvironmentKey)
	if env == "" {
		err = fmtMissingEnvVarError(semconv.DeploymentEnvironmentKey)
		return
	}

	attrs = append(attrs, semconv.DeploymentEnvironment(env))
	attrs = append(attrs, loadContainerResourceFromEnv()...)
	attrs = append(attrs, loadK8sResourceFromEnv()...)
	attrs = append(attrs, loadCloudResourceFromEnv()...)

	res, err = resource.New(
		ctx,
		resource.WithSchemaURL(semconv.SchemaURL),
		withTelemetrySDKStub(),
		withOSStub(),
		withHostStub(),
		withContainerStub(),
		withProcessStub(),
		resource.WithAttributes(attrs...),
	)
	if err != nil {
		err = fmt.Errorf("otel:err creating resource: %w", err)
		return
	}

	return
}

func loadServiceResourceFromEnv() ([]attribute.KeyValue, error) {
	attrs := []attribute.KeyValue{
		semconv.ServiceInstanceID(getOTELEnvVar(semconv.ServiceInstanceIDKey)),
	}

	if v := getOTELEnvVar(semconv.ServiceNameKey); v == "" {
		return nil, fmtMissingEnvVarError(semconv.ServiceNameKey)
	} else {
		attrs = append(attrs, semconv.ServiceName(v))
	}

	if v := getOTELEnvVar(semconv.ServiceNamespaceKey); v == "" {
		// OTEL considers this optional, but we will consider it mandatory to avoid mistakes
		return nil, fmtMissingEnvVarError(semconv.ServiceNamespaceKey)
	} else {
		attrs = append(attrs, semconv.ServiceNamespace(v))
	}

	if v := getOTELEnvVar(semconv.ServiceVersionKey); v == "" {
		// OTEL considers this optional, but we will consider it mandatory to avoid mistakes
		return nil, fmtMissingEnvVarError(semconv.ServiceVersionKey)
	} else {
		attrs = append(attrs, semconv.ServiceVersion(v))
	}

	return attrs, nil
}

func loadContainerResourceFromEnv() []attribute.KeyValue {
	return []attribute.KeyValue{
		semconv.ContainerName(getOTELEnvVar(semconv.ContainerNameKey)),
		semconv.ContainerImageName(getOTELEnvVar(semconv.ContainerImageNameKey)),
		semconv.ContainerImageTags(getOTELEnvVar(semconv.ContainerImageTagsKey)),
		semconv.ContainerRuntime(getOTELEnvVar(semconv.ContainerRuntimeKey)),
	}
}

func loadK8sResourceFromEnv() []attribute.KeyValue {
	attrs := []attribute.KeyValue{
		semconv.K8SClusterName(getOTELEnvVar(semconv.K8SClusterNameKey)),
		semconv.K8SNodeName(getOTELEnvVar(semconv.K8SNodeNameKey)),
		semconv.K8SNodeUID(getOTELEnvVar(semconv.K8SNodeUIDKey)),
		semconv.K8SNamespaceName(getOTELEnvVar(semconv.K8SNamespaceNameKey)),
		semconv.K8SPodName(getOTELEnvVar(semconv.K8SPodNameKey)),
		semconv.K8SPodUID(getOTELEnvVar(semconv.K8SPodUIDKey)),
		semconv.K8SContainerName(getOTELEnvVar(semconv.K8SContainerNameKey)),
		semconv.K8SDeploymentName(getOTELEnvVar(semconv.K8SDeploymentNameKey)),
		semconv.K8SJobName(getOTELEnvVar(semconv.K8SJobNameKey)),
		semconv.K8SCronJobName(getOTELEnvVar(semconv.K8SCronJobNameKey)),
	}

	intV, _ := strconv.Atoi(getOTELEnvVar(semconv.K8SContainerRestartCountKey)) // Intentionally suppressing the err since nothing to do
	attrs = append(attrs, semconv.K8SContainerRestartCount(intV))

	return attrs
}

func loadCloudResourceFromEnv() []attribute.KeyValue {
	return []attribute.KeyValue{
		semconv.CloudProviderKey.String(getOTELEnvVar(semconv.CloudProviderKey)),
		semconv.CloudRegion(getOTELEnvVar(semconv.CloudRegionKey)),
		semconv.CloudAvailabilityZone(getOTELEnvVar(semconv.CloudAvailabilityZoneKey)),
		semconv.CloudPlatformKey.String(getOTELEnvVar(semconv.CloudPlatformKey)),
	}
}

func getOTELEnvVar(key attribute.Key) string {
	return os.Getenv(getOTELEnvVarKey(key))
}

func getOTELEnvVarKey(key attribute.Key) string {
	// Convert key into OTEL_<envvar> format and replace dots with underscore
	return fmt.Sprintf("OTEL_%s",
		strings.ToUpper(
			strings.Replace(string(key), ".", "_", -1),
		),
	)
}

func fmtMissingEnvVarError(key attribute.Key) error {
	return fmt.Errorf("otel:[%s] not provided", getOTELEnvVarKey(key))
}
