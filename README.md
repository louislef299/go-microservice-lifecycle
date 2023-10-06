# Go Microservice Lifecycle

Go is great at concurrency on multi-core machines and networking. I had run
across the [Go Pipelines and Cancellation][] blog post and decided to implement
it myself through the lens of building a microservice and the benefits channels
provide for quickly designing a microservice. Optional: get a [refresher on
microservices][].

From the Go blog post, *Sameer Ajmani* describes the idea of what a Go pipeline
is:

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

## The Color Product

You and your team have been given the objective to migrate the existing Color
service from a monolith into a microservice architecture. The Color service is
very simple and takes incoming text and colors it according to the request.

The product is split into [5 versions][] that simulate a program progression and
changing business requirements. Each version has a separate doc that explains
the iteration further.

## Makefile Targets

The `Makefile` has targets to run each version. Requirements include [docker][]
and [Go][]. Example:

```bash
make v1
```

[5 versions]: https://www.cortex.io/content/the-5-stages-of-the-microservice-life-cycle-and-the-best-tools-to-optimize-them
[docker]: https://www.docker.com/
[Go]: https://go.dev/
[Go Pipelines and Cancellation]: https://go.dev/blog/pipelines
[refresher on microservices]: https://martinfowler.com/articles/microservices.html
