FROM alpine:3.11
ADD omo-msa-startkit /usr/bin/omo-msa-startkit
ENV MSA_REGISTRY_PLUGIN
ENV MSA_REGISTRY_ADDRESS
ENTRYPOINT [ "omo-msa-startkit" ]
