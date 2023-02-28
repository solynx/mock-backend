FROM golang:latest
WORKDIR /Day1
COPY . .
RUN go build -o main .
CMD ["/Day1/main"]