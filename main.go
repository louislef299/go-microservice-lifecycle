package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

type IColor interface {
	Run(ctx context.Context) *Response
}

type Response struct {
	err  error
	resp string
}

func main() {
	// Set the SIGINT context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Run 5 parallel processes to read incoming IColors
	resps := parallel(ctx, 5, gen(generateRandomValues(100)...))

	// Set up the pipeline and consume the output
	for resp := range resps {
		if resp.err != nil {
			log.Println("got an error:", resp.err)
		} else {
			log.Println(resp.resp)
		}
	}
}

// Does the work of printing responses found from the input channel
func printColor(ctx context.Context, in <-chan IColor, num int) <-chan *Response {
	resps := make(chan *Response)
	fmt.Printf("thread %d reporting for duty!\n", num+1)
	go func() {
		defer close(resps)
		for c := range in {
			select {
			case resps <- c.Run(ctx):
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			case <-ctx.Done():
				return
			}
		}
	}()
	return resps
}

// Begins sending the IColors on the channel
func gen(colors ...IColor) <-chan IColor {
	out := make(chan IColor)
	go func() {
		defer close(out)
		for _, n := range colors {
			out <- n
		}
	}()
	return out
}

// Runs num parallel printers and returns a single channel that returns the responses
func parallel(ctx context.Context, num int, in <-chan IColor) <-chan *Response {
	var ch []<-chan *Response
	fmt.Printf("running %d parallel readers\n", num)
	for n := 0; n < num; n++ {
		ch = append(ch, printColor(ctx, in, n))
	}
	return merge(ctx, ch...)
}

// Funnels all channels into one
func merge(ctx context.Context, cs ...<-chan *Response) <-chan *Response {
	var wg sync.WaitGroup
	out := make(chan *Response)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done
	output := func(c <-chan *Response) {
		defer wg.Done()
		for err := range c {
			select {
			case out <- err:
			case <-ctx.Done():
				log.Println("closing out merge")
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Generates num random IColors
func generateRandomValues(num int) []IColor {
	var colors []IColor
	for n := 0; n < num; n++ {
		var c IColor
		switch rand.Intn(3) {
		case 0:
			c = &Red{msg: fmt.Sprintf("this is red %d", n)}
		case 1:
			c = &Blue{msg: fmt.Sprintf("this is blue %d", n)}
		default:
			c = &Err{msg: fmt.Sprintf("this is error %d", n)}
		}
		colors = append(colors, c)
	}
	return colors
}
