package feature_flag

import "github.com/sirupsen/logrus"

type FeatureFlag struct {
	dumpConfig       bool
	trace            bool
	setTestUserToCtx bool
}

var Instance *FeatureFlag

func Init(
	dumpConfig,
	traceEnabled,
	setTestUserToCtx bool) {
	Instance = &FeatureFlag{
		dumpConfig:       dumpConfig,
		trace:            traceEnabled,
		setTestUserToCtx: setTestUserToCtx,
	}
}

func Get() *FeatureFlag {
	if Instance == nil {
		Init(false, false, false)
		logrus.Warn("feature_flag - GetFeatureFlags - Instance is nil, init default")
	}
	return Instance
}

func (f *FeatureFlag) DumpConfigEnabled() bool {
	return f.dumpConfig
}

func (f *FeatureFlag) TraceEnabled() bool {
	return f.trace
}

func (f *FeatureFlag) TestUserEmbeddedInContext() bool {
	return f.setTestUserToCtx
}
