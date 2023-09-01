FROM alpine as runner

RUN mkdir app
COPY dist/tech-challenge app/
RUN chmod +x app
COPY ./internal/config/config.yaml app

EXPOSE 3001
WORKDIR /app

ENTRYPOINT [ "./tech-challenge", "--config-dir", "." ]