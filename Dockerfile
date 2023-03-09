FROM ubuntu
WORKDIR /home
RUN apt-get update -y 

RUN apt install -y build-essential golang rustc vim git
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y

#RUN source '$HOME/.cargo/env'
#RUN rustup target add wasm32-unknown-unknown

COPY ./bin/wasmd /usr/bin/
RUN chmod 777 /usr/bin/wasmd

COPY ./pkg/  $GOPATH/pkg/

COPY ./init_script .

CMD ["tail", "-f", "/dev/null"]