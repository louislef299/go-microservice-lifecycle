# Pipeline Example

Go is great at concurrency on multi-core machines and networking. I had run
across the [Go Pipelines and Cancellation][] blog post and decided to implement
it myself for practice.

From the blog post, *Sameer Ajmani* describes the idea of what a Go pipeline is
in this context:

> There’s no formal definition of a pipeline in Go; it’s just one of many kinds
> of concurrent programs. Informally, a pipeline is a series of stages connected
> by channels, where each stage is a group of goroutines running the same
> function. In each stage, the goroutines\
> \
> -receive values from upstream via inbound channels\
> -perform some function on that data, usually producing new values\
> -send values downstream via outbound channels\
> \
> Each stage has any number of inbound and outbound channels, except the first
> and last stages, which have only outbound or inbound channels, respectively.
> The first stage is sometimes called the source or producer; the last stage,
> the sink or consumer.

The beauty of breaking up code in this fashion right away, is that migrating
from a small application that runs on a single machines, to running this as a
microservice becomes much easier. Let me explain.

## The IColor Project

IColor is a simple example of how to split jobs between upstream and downstream
jobs. This is the current project:

![IColor Flow Chart](./docs/IColor%20pipeline.png)

[Go Pipelines and Cancellation]: https://go.dev/blog/pipelines
