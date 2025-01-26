package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitBase(t *testing.T) {
	tests := []struct {
		name     string
		setupEnv func()
		wantErr  bool
	}{
		{
			name:    "success - all env vars set",
			wantErr: false,
		},
		{
			name: "failure - missing OTEL_SERVICE_NAME",
			setupEnv: func() {
				t.Setenv("OTEL_SERVICE_NAME", "")
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupEnv != nil {
				tt.setupEnv()
			}

			ctx, cleanup, err := initBase()
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			defer cleanup()
			require.NotNil(t, ctx, "context should not be nil")
		})
	}
}

func TestInitServer(t *testing.T) {
	log.Println("hello world")
	log.Println(os.Getenv("OTEL_SERVICE_NAME"))
	tests := []struct {
		name    string
		port    string
		wantErr bool
	}{
		{
			name:    "valid port",
			port:    "8080",
			wantErr: false,
		},
		{
			name:    "invalid port",
			port:    "invalid",
			wantErr: true,
		},
	}

	ctx, cleanup, err := initBase()
	require.NoError(t, err, "setup should not fail")
	defer cleanup()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("PORT", tt.port)

			_, err := initServer(ctx)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
