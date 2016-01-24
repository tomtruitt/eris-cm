FROM quay.io/eris/base
MAINTAINER Eris Industries <support@erisindustries.com>

#-----------------------------------------------------------------------------
# dependencies

#-----------------------------------------------------------------------------
# install eris-chainmaker

ENV REPO $GOPATH/src/github.com/eris-ltd/eris-chainmaker
COPY . $REPO
WORKDIR $REPO/cmd/eris-chainmaker
RUN go build -o /usr/local/bin/eris-chainmaker
RUN chown --recursive $USER:$USER $REPO

#-----------------------------------------------------------------------------
# persist data, set user
RUN chown --recursive $USER:$USER /home/$USER
VOLUME /home/$USER/.eris
WORKDIR /home/$USER/.eris
USER $USER
ENTRYPOINT ["eris-chainmaker" ]
