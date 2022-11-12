FROM golang:alpine as build

WORKDIR /app

COPY . ./

RUN go build -o /api-todo

FROM gcr.io/distroless/base-debian10
COPY --from=build /api-todo /api-todo
EXPOSE 9000
ENTRYPOINT ["/api-todo"]
