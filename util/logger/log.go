package logger

import "go.uber.org/zap"

var (
	Zap, _ = zap.NewDevelopment()
	Info   = Zap.Info
	Debug  = Zap.Debug
	Warn   = Zap.Warn
	Error  = Zap.Error
	Fatal  = Zap.Fatal
	Panic  = Zap.Panic
	Sync   = Zap.Sync
	With   = Zap.With
)
