# Description

This is a playground for playing with deno and docker.

It currently builds a docker image with a natively compiles executable with ubuntu into distroless.

```shell
//builds it
docker build -t hello-deno-ts .
```

After that you can run it with docker:

```shell
//runs it
docker run --name hello-deno-ts hello-deno-ts
```

Thats all.
