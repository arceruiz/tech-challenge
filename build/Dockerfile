FROM golang as builder

RUN mkdir app
COPY ./ app
WORKDIR app
RUN CGO_ENABLED=0 go build -o dist/tech-challenge cmd/client/main.go

FROM alpine as runner

RUN mkdir app
COPY --from=builder go/app/dist/tech-challenge app/
RUN chmod +x app

EXPOSE 3001
WORKDIR /app

ENTRYPOINT [ "./tech-challenge", "--config-dir", "." ]