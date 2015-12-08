# ConfigurationProcessor-go

To run the docker container:

1. mkdir -p /var/tmp/ConfigurationProcessor/output
2. cp -r ./config /var/tmp/ConfigurationProcessor
3. docker build -t ftcjeff/cfgproc .
4. docker run -v /var/tmp/ConfigurationProcessor:/opt/ConfigurationProcessor ftcjeff/cfgproc
