FROM drmjo/terraform:0.11.6

USER root

ARG WORKDIR=/tfstack
RUN mkdir $WORKDIR

COPY . $WORKDIR

WORKDIR /tfstack
