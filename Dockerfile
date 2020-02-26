FROM alpine:3.11
ADD omo-msa-startkit /usr/bin/omo-msa-startkit
ENTRYPOINT [ "omo-msa-startkit" ]
