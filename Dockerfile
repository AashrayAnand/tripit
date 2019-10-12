FROM golang:latest
LABEL maintainer="aashrayanand01@gmail.com"
RUN go get github.com/AashrayAnand/tripit/
EXPOSE 8080
CMD ["TripIt"]
