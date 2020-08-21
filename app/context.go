package app

import (
	"context"
)

type Context struct {
	context.Context
	cancel context.CancelFunc
}

func NewContext() *Context {
	ctx, cancel := context.WithCancel(context.Background())
	return &Context{ctx, cancel}
}
