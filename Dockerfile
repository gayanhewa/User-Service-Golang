# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

#Configure bunch
RUN mkdir /bunch
RUN mkdir /app

ENV GOPATH /bunch
WORKDIR /bunch
RUN go get github.com/dkulchenko/bunch
RUN go install github.com/dkulchenko/bunch
RUN pwd;ls bin
RUN cp /bunch/bin/bunch /usr/bin/bunch 
RUN which bunch

WORKDIR /app
ADD . /app
RUN bunch 


ENV GOPATH /app
RUN pwd

# Install my dependencies
RUN bunch update 

# Recompile the binaries
RUN bunch rebuild


# Launch the server 
ENTRYPOINT bunch go run /app/server.go

# Document that the service listens on port 8080.
EXPOSE 9091
