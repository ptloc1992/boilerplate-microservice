FROM golang:latest
LABEL maintainer="David Pham <ptloc.cnpm@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8687
CMD ["./main", "start"]