FROM golang:latest
LABEL maintainer="aashrayanand01@gmail.com"
RUN go get github.com/AashrayAnand/Bill-List/
EXPOSE 8080
CMD ["Bill-List"]
