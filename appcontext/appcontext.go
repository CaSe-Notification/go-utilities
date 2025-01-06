package appcontext

import (
	"context"

	"github.com/CaSe-Notification/go-utilities/logger"
	"github.com/segmentio/ksuid"
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

func generateID() string {
	return ksuid.New().String()
}
