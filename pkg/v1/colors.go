package v1

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

type Red struct {
	msg string
}

func (r *Red) Run(ctx context.Context) *Response {
	red := color.New(color.FgRed)
	return &Response{
		resp: red.Sprintf("%s", r.msg),
	}
}

type Blue struct {
	msg string
}

func (b *Blue) Run(ctx context.Context) *Response {
	blue := color.New(color.FgBlue)
	return &Response{
		resp: blue.Sprintf("%s", b.msg),
	}
}

type Err struct {
	msg string
}

func (e *Err) Run(ctx context.Context) *Response {
	return &Response{
		err: fmt.Errorf(e.msg),
	}
}
