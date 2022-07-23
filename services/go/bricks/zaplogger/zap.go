package zaplogger

import (
	"go.uber.org/zap"
)

// ZapLoggingConfig is a necessary configuration for creation a Logger.
type ZapLoggingConfig struct {
	Level                    string
	Development              bool
	DisableRedirectStdLog    bool
	DisableReplaceGlobals    bool
	DisableSampling          bool
	DisableStacktrace        bool
	DisableCaller            bool
	Encoding                 string
	SamplingConfigInitial    int
	SamplingConfigThereafter int
	Options                  []zap.Option
	InitialFields            []zap.Field
}

// NewZapLogger creates zap.Logger.
func NewZapLogger(cfg ZapLoggingConfig) (*zap.Logger, error) {
	prodCfg := zap.NewProductionConfig()
	atomicLevel := zap.NewAtomicLevel()
	if err := atomicLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, err
	}
	prodCfg.Level = atomicLevel
	prodCfg.Development = cfg.Development
	prodCfg.DisableStacktrace = cfg.DisableStacktrace
	prodCfg.DisableCaller = cfg.DisableCaller

	// Default encoding is JSON
	if cfg.Encoding != "" {
		prodCfg.Encoding = cfg.Encoding
	}

	if cfg.DisableSampling {
		prodCfg.Sampling = nil
	} else {
		if cfg.SamplingConfigInitial != 0 {
			// Default Initial Sampling is 100
			prodCfg.Sampling.Initial = cfg.SamplingConfigInitial
		}
		if cfg.SamplingConfigThereafter != 0 {
			// Default Thereafter Sampling is 100
			prodCfg.Sampling.Thereafter = cfg.SamplingConfigThereafter
		}
	}

	logger, err := prodCfg.Build(cfg.Options...)
	if err != nil {
		return nil, err
	}
	logger = logger.With(cfg.InitialFields...)

	if !cfg.DisableRedirectStdLog {
		zap.RedirectStdLog(logger)
	}

	if !cfg.DisableReplaceGlobals {
		zap.ReplaceGlobals(logger)
	}

	return logger, nil
}
