package main

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

type Red struct {
	msg string
}

func (r *Red) Run(ctx context.Context) error {
	color.Red(r.msg)
	return nil
}

type Blue struct {
	msg string
}

func (b *Blue) Run(ctx context.Context) error {
	color.Blue(b.msg)
	return nil
}

type Err struct {
	msg string
}

func (e *Err) Run(ctx context.Context) error {
	return fmt.Errorf(e.msg)
}
