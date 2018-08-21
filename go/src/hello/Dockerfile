FROM golang:latest
MAINTAINER vladiir.rod96@gmail.com
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main .
EXPOSE 8080 
CMD ["/app/main"]


