FROM golang:1.21-alpine as base

#ENV WORKING_DIR=/src
#ENV TARGET_DIR=/app
#ENV EXEC=fincal
#ENV CONFIG_FILE=config.json
#ENV TOKENS_DIR=tokens

FROM base as dev

WORKDIR /app/backend
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./ ./

RUN apk --no-cache -U add build-base

# See https://mt165.co.uk/blog/static-link-go/
RUN go build -ldflags "-linkmode 'external' -extldflags '-static'"  -o fincal

#FROM scratch
#WORKDIR /app/backend
#COPY --from=build /src/.env .
#COPY --from=build /src/config.json .
#COPY --from=build /src/tokens/* ./tokens/
#COPY --from=build /src/fincal .

EXPOSE 9000

#ENTRYPOINT ["./fincal", "serve"]
CMD ["/bin/sh", "-c", "ls", "-l", "/app/backend"]
#CMD ["/bin/sh", "-c", "./fincal serve"]