FROM golang

ADD . /go/src/github.com/ftcjeff/ConfigurationProcessor

RUN cd /go/src/github.com/ftcjeff/ConfigurationProcessor && ./build.sh

ENTRYPOINT /go/bin/ConfigurationProcessor -input=/opt/ConfigurationProcessor/config/definition.json -output=/opt/ConfigurationProcessor/output
