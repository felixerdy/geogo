FROM golang:onbuild

FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
RUN go get github.com/graphql-go/graphql
RUN go build -o main . 
CMD ["/app/main"]