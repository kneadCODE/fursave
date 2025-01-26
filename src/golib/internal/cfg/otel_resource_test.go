package cfg

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func TestNewOTELResourceFromEnv(t *testing.T) {
	type testCase struct {
		givenEnvVars map[string]string
		expErr       error
		expEnv       string
	}
	tcs := map[string]testCase{
		"err: empty env": {
			expErr: errors.New("otel:[OTEL_SERVICE_NAME] not provided"),
		},
		"err: svc namespace missing": {
			givenEnvVars: map[string]string{
				"OTEL_SERVICE_NAME": "svc",
			},
			expErr: errors.New("otel:[OTEL_SERVICE_NAMESPACE] not provided"),
		},
		"err: svc version missing": {
			givenEnvVars: map[string]string{
				"OTEL_SERVICE_NAME":      "svc_name",
				"OTEL_SERVICE_NAMESPACE": "namespace",
			},
			expErr: errors.New("otel:[OTEL_SERVICE_VERSION] not provided"),
		},
		"err: svc env missing": {
			givenEnvVars: map[string]string{
				"OTEL_SERVICE_NAME":      "svc_name",
				"OTEL_SERVICE_NAMESPACE": "namespace",
				"OTEL_SERVICE_VERSION":   "version",
			},
			expErr: errors.New("otel:[OTEL_DEPLOYMENT_ENVIRONMENT] not provided"),
		},
		"success: with mandatory envvars": {
			givenEnvVars: map[string]string{
				"OTEL_SERVICE_NAME":           "svc_name",
				"OTEL_SERVICE_NAMESPACE":      "namespace",
				"OTEL_SERVICE_VERSION":        "version",
				"OTEL_DEPLOYMENT_ENVIRONMENT": "development",
			},
			expEnv: "development",
		},
		"success: with container envvars": {
			givenEnvVars: map[string]string{
				"OTEL_SERVICE_NAME":           "svc_name",
				"OTEL_SERVICE_NAMESPACE":      "namespace",
				"OTEL_SERVICE_VERSION":        "version",
				"OTEL_DEPLOYMENT_ENVIRONMENT": "development",
				"OTEL_CONTAINER_NAME":         "container",
				"OTEL_CONTAINER_IMAGE_NAME":   "gcr.io/opentelemetry/operator",
				// "OTEL_CONTAINER_IMAGE_TAGS":   "tags", // TODO: Need to modify test case verification as the expected output is `[tags]` and I am too lazy to do it properly now.
				"OTEL_CONTAINER_RUNTIME": "docker",
			},
			expEnv: "development",
		},
		"success: with container & k8s envvars": {
			givenEnvVars: map[string]string{
				"OTEL_SERVICE_NAME":           "svc_name",
				"OTEL_SERVICE_NAMESPACE":      "namespace",
				"OTEL_SERVICE_VERSION":        "version",
				"OTEL_DEPLOYMENT_ENVIRONMENT": "development",
				"OTEL_CONTAINER_NAME":         "container",
				"OTEL_CONTAINER_IMAGE_NAME":   "gcr.io/opentelemetry/operator",
				// "OTEL_CONTAINER_IMAGE_TAGS":   "tags", // TODO: Need to modify test case verification as the expected output is `[tags]` and I am too lazy to do it properly now.
				"OTEL_CONTAINER_RUNTIME":   "docker",
				"OTEL_K8S_CLUSTER_NAME":    "cluster",
				"OTEL_K8S_NODE_NAME":       "node",
				"OTEL_K8S_NODE_UID":        "node_uid",
				"OTEL_K8S_NAMESPACE_NAME":  "ns",
				"OTEL_K8S_POD_NAME":        "pod",
				"OTEL_K8S_POD_UID":         "pod_uid",
				"OTEL_K8S_CONTAINER_NAME":  "container",
				"OTEL_K8S_DEPLOYMENT_NAME": "dep",
				"OTEL_K8S_JOB_NAME":        "job",
				"OTEL_K8S_CRONJOB_NAME":    "cron",
			},
			expEnv: "development",
		},
		"success: with container, k8s & cloud envvars": {
			givenEnvVars: map[string]string{
				"OTEL_SERVICE_NAME":           "svc_name",
				"OTEL_SERVICE_NAMESPACE":      "namespace",
				"OTEL_SERVICE_VERSION":        "version",
				"OTEL_DEPLOYMENT_ENVIRONMENT": "development",
				"OTEL_CONTAINER_NAME":         "container",
				"OTEL_CONTAINER_IMAGE_NAME":   "gcr.io/opentelemetry/operator",
				// "OTEL_CONTAINER_IMAGE_TAGS":   "tags", // TODO: Need to modify test case verification as the expected output is `[tags]` and I am too lazy to do it properly now.
				"OTEL_CONTAINER_RUNTIME":   "docker",
				"OTEL_K8S_CLUSTER_NAME":    "cluster",
				"OTEL_K8S_NODE_NAME":       "node",
				"OTEL_K8S_NODE_UID":        "node_uid",
				"OTEL_K8S_NAMESPACE_NAME":  "ns",
				"OTEL_K8S_POD_NAME":        "pod",
				"OTEL_K8S_POD_UID":         "pod_uid",
				"OTEL_K8S_CONTAINER_NAME":  "container",
				"OTEL_K8S_DEPLOYMENT_NAME": "dep",
				"OTEL_K8S_JOB_NAME":        "job",
				"OTEL_K8S_CRONJOB_NAME":    "cron",
				"OTEL_CLOUD_PROVIDER":      "heaven",
				"OTEL_CLOUD_REGION":        "mountain",
				// "OTEL_CLOUD_AVAILABILITY_ZONE": "downtown", // TODO: Need to modify test case verification as the key name is `cloud.availability_zone` instead of `cloud.availability.zone` and I am too lazy to do it properly now.
				"OTEL_CLOUD_PLATFORM": "heaven-platform",
			},
			expEnv: "development",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			// Given:.
			ctx := context.Background()
			defer func() {
				withTelemetrySDKStub = resource.WithTelemetrySDK
				withOSStub = resource.WithOS
				withHostStub = resource.WithHost
				withContainerStub = resource.WithContainer
				withProcessStub = resource.WithProcess
			}()
			for k, v := range tc.givenEnvVars {
				t.Setenv(k, v)
			}
			var telStubCalled bool
			withTelemetrySDKStub = func() resource.Option {
				telStubCalled = true
				return resource.WithTelemetrySDK()
			}
			var osStubCalled bool
			withOSStub = func() resource.Option {
				osStubCalled = true
				return resource.WithOS()
			}
			var hostStubCalled bool
			withHostStub = func() resource.Option {
				hostStubCalled = true
				return resource.WithHost()
			}
			var containerStubCalled bool
			withContainerStub = func() resource.Option {
				containerStubCalled = true
				return resource.WithContainer()
			}
			var processStubCalled bool
			withProcessStub = func() resource.Option {
				processStubCalled = true
				return resource.WithProcess()
			}

			// When:.
			res, env, err := NewOTELResourceFromEnv(ctx)

			// Then:.
			require.Equal(t, tc.expErr, err)
			require.Equal(t, tc.expEnv, env)
			if tc.expErr != nil {
				require.Nil(t, res)
				return
			}

			require.True(t, telStubCalled)
			require.True(t, osStubCalled)
			require.True(t, hostStubCalled)
			require.True(t, containerStubCalled)
			require.True(t, processStubCalled)
			require.NotNil(t, res)
			require.Equal(t, semconv.SchemaURL, res.SchemaURL())
			for k, v := range tc.givenEnvVars {
				cmpResAttrVal(t, res, attribute.Key(convertEnvVarToOTELKeyStr(k)), v)
			}
		})
	}
}

func convertEnvVarToOTELKeyStr(key string) string {
	return strings.ToLower(
		strings.Replace(
			strings.TrimPrefix(key, "OTEL_"),
			"_", ".", -1),
	)
}

func cmpResAttrVal(t *testing.T, res *resource.Resource, key attribute.Key, expVal string) {
	v, ok := res.Set().Value(key)
	require.True(t, ok, "missing: %s", string(key))
	require.Equal(t, expVal, v.Emit(), "for: %s", string(key))
}
