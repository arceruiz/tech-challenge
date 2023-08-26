FROM golang as builder

RUN mkdir app
COPY ./ app
WORKDIR app
RUN go build -o dist/tech-challenge cmd/client/main.go

FROM alpine as runner

RUN mkdir app
COPY --from=builder go/app/dist/tech-challenge app/tech-challenge
RUN chmod +x app
COPY ./internal/config/config.yaml app

EXPOSE 3001
WORKDIR /app

ENTRYPOINT [ "./tech-challenge", "--config-dir", "." ]