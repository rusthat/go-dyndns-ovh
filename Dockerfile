# TODO: implement DynDNS OVH update docker container
# TODO: Test env var config
FROM alpine:latest
LABEL maintainer="Jerome Grassnick <grassnick@pm.me>"
WORKDIR /bin
COPY ./bin/ovh-dyndns /bin/ovh-dyndns
CMD ["/bin/ovh-dyndns"]