FROM golang:1.10.2-stretch

ARG USER=go
RUN adduser --disabled-password --gecos "" $USER

ARG WORKDIR=/home/$USER/app
WORKDIR $WORKDIR

# make sure the github directory exists
RUN mkdir -p src/github.com
RUN chown -R $USER:$USER $WORKDIR

USER $USER

ENV GOPATH $WORKDIR
ENV GOBIN $WORKDIR/bin
ENV PATH $GOBIN:$PATH
# enable color prompt
RUN sed -i 's/#force_color_prompt/force_color_prompt/' $HOME/.bashrc
