FROM alpine:latest

RUN mkdir app
ADD dist/ app
RUN chmod +x app
COPY ./internal/config/config.yaml app

EXPOSE 3001
WORKDIR /app

ENTRYPOINT [ "./tech-challenge", "--config-dir", "." ]