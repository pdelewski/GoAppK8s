FROM golang:1.13-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/

RUN CGO_ENABLED=0 go build -o /bin/logwriter

FROM scratch
COPY --from=build /bin/logwriter /bin/logwriter
ENTRYPOINT ["/bin/logwriter"]
