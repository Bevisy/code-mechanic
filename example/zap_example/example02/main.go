package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	url := "https://zbbx.ml:30443"

	// When performance and type safety are critical, use the Logger
	logger.Info("failed to fetch URL", zap.String("url", url), zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
	logger.Debug("failed to fetch URL", zap.String("url", url))

	// In contexts where performance is nice, but not critical
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)
	sugar.Infof("Failed to fetch URL: %s", url)
}
