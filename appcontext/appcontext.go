package appcontext

import (
	"context"

	"github.com/CaSe-Notification/go-utilities/logger"
	"github.com/segmentio/ksuid"
)

type contextKey int

const (
	userContextKey contextKey = iota
	ipContextKey
	langContextKey
)

type AppContext struct {
	requestID string
	traceID   string
	logger    *logger.Logger
	context   context.Context
}

type Fields = logger.Fields

func newWithSource(ctx context.Context, source string) *AppContext {
	var (
		requestID = generateID()
		traceID   = generateID()
	)

	return &AppContext{
		requestID: requestID,
		traceID:   traceID,
		logger:    logger.NewLogger(logger.Fields{"requestId": requestID, "traceId": traceID, "source": source}),
		context:   ctx,
	}
}

func NewRest(ctx context.Context) *AppContext {
	return newWithSource(ctx, "rest")
}

func NewGRPC(ctx context.Context) *AppContext {
	return newWithSource(ctx, "grpc")
}

func NewWorker(ctx context.Context) *AppContext {
	return newWithSource(ctx, "worker")
}

func (appCtx *AppContext) SetTraceID(traceID string) {
	appCtx.traceID = traceID
	appCtx.logger.AddData(logger.Fields{"traceId": traceID})
}

func (appCtx *AppContext) GetTraceID() string {
	return appCtx.traceID
}

func (appCtx *AppContext) AddLogData(fields Fields) {
	appCtx.logger.AddData(fields)
}

func (appCtx *AppContext) Logger() *logger.Logger {
	return appCtx.logger

}

func (appCtx *AppContext) Context() context.Context {
	return appCtx.context
}

func (appCtx *AppContext) SetContext(ctx context.Context) {
	appCtx.context = ctx
}

func (appCtx *AppContext) SetUserID(id string) {
	appCtx.context = context.WithValue(appCtx.context, userContextKey, id)
}

func (appCtx *AppContext) GetUserID() string {
	id, ok := appCtx.context.Value(userContextKey).(string)
	if !ok {
		return ""
	}
	return id
}

func (appCtx *AppContext) SetIP(ip string) {
	appCtx.context = context.WithValue(appCtx.context, ipContextKey, ip)
}

func (appCtx *AppContext) GetIP() string {
	ip, ok := appCtx.context.Value(ipContextKey).(string)
	if !ok {
		return ""
	}
	return ip
}

func (appCtx *AppContext) SetLang(lang string) {
	appCtx.context = context.WithValue(appCtx.context, langContextKey, lang)
}

func (appCtx *AppContext) GetLang() Language {
	lang, ok := appCtx.context.Value(langContextKey).(string)
	if !ok {
		return Vietnamese
	}

	dLang := ToLanguage(lang)
	if !dLang.IsValid() {
		return Vietnamese
	}

	return dLang
}

func generateID() string {
	return ksuid.New().String()
}
