FROM golang:alpine as build

WORKDIR /app

COPY . ./

RUN go build -ldflags "-s -w" -o /api-todo

FROM scratch
COPY --from=build /api-todo /api-todo
EXPOSE 9000
ENTRYPOINT ["/api-todo"]
