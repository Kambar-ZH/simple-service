package logger

import (
	"context"
	"os"
	"time"

	"github.com/Kambar-ZH/simple-service/pkg/tools/tracing_tools"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Level = zapcore.Level

const (
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel Level = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel // 5
	DebugLevel Level = zap.DebugLevel // -1

	TracingMetadataKey string = "tracing_metadata"
)

type logger struct {
	l     *zap.Logger // zap ensure that zap.Logger is safe for concurrent use
	level Level
}

type Field = zap.Field

func (l *logger) withCtxFields(ctx context.Context) (fields []Field) {
	tracingMetadata, ok := ctx.Value(TracingMetadataKey).(tracing_tools.TracingMetadata)
	if ok {
		fields = append(fields, Any(TracingMetadataKey, tracingMetadata))
	}

	return
}

func (l *logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func (l *logger) WithCtx(ctx context.Context) Logger {
	l.l = l.l.With(l.withCtxFields(ctx)...)
	return l
}

var (
	Skip       = zap.Skip
	Binary     = zap.Binary
	Bool       = zap.Bool
	Boolp      = zap.Boolp
	ByteString = zap.ByteString
	String     = zap.String
	Float64    = zap.Float64
	Float64p   = zap.Float64p
	Float32    = zap.Float32
	Float32p   = zap.Float32p
	Any        = zap.Any
)

type LevelEnablerFunc func(lvl Level) bool

type RotateOptions struct {
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

type TeeOption struct {
	Filename string
	Ropt     RotateOptions
	Lef      LevelEnablerFunc
}

type Option = zap.Option

func NewTeeWithRotate(tops []TeeOption, opts ...Option) Logger {
	var cores []zapcore.Core
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02T15:04:05"))
	}

	for _, top := range tops {
		top := top

		lv := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return top.Lef(Level(lvl))
		})

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   top.Filename,
			MaxSize:    top.Ropt.MaxSize,
			MaxBackups: top.Ropt.MaxBackups,
			MaxAge:     top.Ropt.MaxAge,
			Compress:   top.Ropt.Compress,
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.AddSync(w),
			lv,
		)
		cores = append(cores, core)
	}

	cores = append(cores, zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(os.Stdout),
		InfoLevel,
	))

	lgr := &logger{
		l: zap.New(zapcore.NewTee(cores...), opts...),
	}
	return lgr
}
