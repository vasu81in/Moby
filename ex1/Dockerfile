FROM golang:1.7
RUN mkdir -p /app
ADD . /app/
WORKDIR /app/
RUN go build -o /app/wordbuilder
RUN ./wordbuilder
 
