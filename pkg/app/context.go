package app

import (
	"context"
	"time"
)

// AppContext implements context.Context
type AppContext struct {
	ctx    context.Context
	dbUri  string
	dbName string
}

func (a *AppContext) Deadline() (deadline time.Time, ok bool) {
	return a.ctx.Deadline()
}

func (a *AppContext) Done() <-chan struct{} {
	return a.ctx.Done()
}

func (a *AppContext) Err() error {
	return a.ctx.Err()
}

func (a *AppContext) Value(key any) any {
	return a.ctx.Value(key)
}

// Context returns the original context default to context.Background if nil
func (a *AppContext) Context() context.Context {
	if a.ctx == nil {
		a.ctx = context.Background()
	}
	return a.ctx
}

func (a *AppContext) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *AppContext) DbUri() string {
	return a.dbUri
}

func (a *AppContext) SetDbUri(uri string) {
	a.dbUri = uri
}

func (a *AppContext) DbName() string {
	return a.dbName
}

func (a *AppContext) SetDbName(name string) {
	a.dbName = name
}
