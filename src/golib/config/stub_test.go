package config

import (
	"github.com/kneadCODE/fursave/src/golib/internal/cfg"
)

func resetStubs() {
	newOTELResourceFromEnvStub = cfg.NewOTELResourceFromEnv
}
