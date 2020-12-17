## We specify the base image we need for our go application
FROM golang:1.14-alpine
RUN mkdir /WebApp
ADD . /WebApp
WORKDIR /WebApp
## Add go mod download to pull all dependencies
RUN go mod download
## Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .
## newly created binary executable
CMD ["/WebApp/main"]