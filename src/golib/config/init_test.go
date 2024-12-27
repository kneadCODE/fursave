package config

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func TestInit(t *testing.T) {
	type testCase struct {
		givenSentryEnabled                  bool
		mockRes                             *resource.Resource
		mockEnvStr                          string
		mockResErr                          error
		expApp                              App
		expErr                              error
		expNewOTELResourceFromEnvStubCalled bool
	}

	tcs := map[string]testCase{
		"res err": {
			mockResErr:                          errors.New("some err"),
			expErr:                              errors.New("some err"),
			expNewOTELResourceFromEnvStubCalled: true,
		},
		"invalid env": {
			mockRes:                             resource.NewWithAttributes(semconv.SchemaURL, semconv.DeploymentEnvironment("dev")),
			mockEnvStr:                          "dev",
			expErr:                              errors.New("invalid env: [dev]"),
			expNewOTELResourceFromEnvStubCalled: true,
		},
		"success": {
			mockRes:                             resource.NewWithAttributes(semconv.SchemaURL, semconv.DeploymentEnvironment("development")),
			expApp:                              App{Env: EnvDev, Res: resource.NewWithAttributes(semconv.SchemaURL, semconv.DeploymentEnvironment("development"))},
			mockEnvStr:                          "development",
			expNewOTELResourceFromEnvStubCalled: true,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			// Given:
			defer resetStubs()
			var newOTELResourceFromEnvStubCalled bool
			newOTELResourceFromEnvStub = func(ctx context.Context) (*resource.Resource, string, error) {
				newOTELResourceFromEnvStubCalled = true
				return tc.mockRes, tc.mockEnvStr, tc.mockResErr
			}

			// When:
			ctx, err := Init()

			// Then:
			require.Equal(t, tc.expNewOTELResourceFromEnvStubCalled, newOTELResourceFromEnvStubCalled)

			if tc.expErr != nil {
				require.Equal(t, tc.expErr, err)
				require.EqualValues(t, tc.expApp, AppFromContext(ctx))
			} else {
				require.Nil(t, err)

				app := AppFromContext(ctx)
				require.EqualValues(t, tc.expApp, app)
			}
		})
	}
}
