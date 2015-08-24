# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

#Configure bunch
WORKDIR /go
RUN go get github.com/dkulchenko/bunch
RUN go install github.com/dkulchenko/bunch
RUN cp /go/bin/bunch /usr/bin;chmod 755 /usr/bin/bunch
RUN which bunch

RUN mkdir /app
WORKDIR /app
# Setup the app source
ADD . /app
RUN cp /go/bin/bunch /app/bin/bunch

RUN export GOPATH=/app
RUN cd /app
# Install my dependencies
RUN /app/bin/bunch update 

# Recompile the binaries
RUN /app/bin/bunch rebuild


# Launch the server 
ENTRYPOINT /app/bin/bunch go run /app/server.go

# Document that the service listens on port 8080.
EXPOSE 9091
