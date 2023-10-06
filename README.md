# Go Microservice Lifecycle

Go is great at concurrency on multi-core machines and networking. I had run
across the [Go Pipelines and Cancellation][] blog post and decided to implement
it myself through the lens of building a microservice.

From the blog post, *Sameer Ajmani* describes the idea of what a Go pipeline is:

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

By breaking up code in this fashion, it becomes easier to migrate from a small
application that runs on a single machine, to running this as a microservice.
Adding features or additional downstream/upstream functions becomes much more
intuitive and decouples the functions further.

## The IColor Project

IColor is a simple example of how to split jobs between upstream and downstream
jobs. The premise is simple: IColor is an interface that implements `Run()`.
Upstream services generate IColor types and send them to downstream services
that execute `Run()`. The `Response` value is then sent to the executing
function.

![IColor Flow Chart](./docs/IColor-pipeline-light.png#gh-light-mode-only)
![IColor Flow Chart](./docs/IColor-pipeline-dark.png#gh-dark-mode-only)

The project is split into [5 versions][] that simulate a traditional program
progression. Each version has a `README.md` file that explains the iteration
further.

## Makefile Targets

The `Makefile` has targets to run each version. Requirements include [docker][]
and [Go][]. Example:

```bash
make v1 # Runs v1 of the project
```

[5 versions]: https://www.cortex.io/content/the-5-stages-of-the-microservice-life-cycle-and-the-best-tools-to-optimize-them
[docker]: https://www.docker.com/
[Go]: https://go.dev/
[Go Pipelines and Cancellation]: https://go.dev/blog/pipelines
