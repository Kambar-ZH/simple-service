package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
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

	RequestIDKey string = "request_id"
	AppKey       string = "app"
)

type logger struct {
	ctx   context.Context
	l     *zap.Logger // zap ensure that zap.Logger is safe for concurrent use
	level Level
}

type Field = zap.Field

func (l *logger) lgrFields(fields ...Field) (lgrFields []Field) {

	if l.ctx != nil {
		requestID := l.ctx.Value(RequestIDKey).(string)
		lgrFields = append(lgrFields, String(RequestIDKey, requestID))

		app := l.ctx.Value(AppKey).(string)
		lgrFields = append(lgrFields, String(AppKey, app))
	}

	lgrFields = append(lgrFields, fields...)

	return
}

func (l *logger) Debug(msg string, fields ...Field) {
	lgrFields := l.lgrFields(fields...)
	l.l.Debug(msg, lgrFields...)
}

func (l *logger) Info(msg string, fields ...Field) {
	lgrFields := l.lgrFields(fields...)
	l.l.Info(msg, lgrFields...)
}

func (l *logger) Warn(msg string, fields ...Field) {
	lgrFields := l.lgrFields(fields...)
	l.l.Warn(msg, lgrFields...)
}

func (l *logger) Error(msg string, fields ...Field) {
	lgrFields := l.lgrFields(fields...)
	l.l.Error(msg, lgrFields...)
}

func (l *logger) Panic(msg string, fields ...Field) {
	lgrFields := l.lgrFields(fields...)
	l.l.Panic(msg, lgrFields...)
}

func (l *logger) Fatal(msg string, fields ...Field) {
	lgrFields := l.lgrFields(fields...)
	l.l.Fatal(msg, lgrFields...)
}

func (l *logger) WithCtx(ctx context.Context) Logger {
	l.ctx = ctx
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
