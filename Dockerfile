FROM alpine:3.21
RUN mkdir "/app"
WORKDIR "/app"
COPY cronServer /app/cronServer
COPY config.yaml /app/config.yaml
ENTRYPOINT ["./cronServer"]