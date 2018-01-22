FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install --yes software-properties-common
RUN add-apt-repository --yes ppa:bitcoin/bitcoin
RUN apt-get update

RUN apt-get install --yes bitcoind make

WORKDIR /var/btcrpc

EXPOSE 8334 8334

COPY bitcoin.conf .

CMD bitcoind -testnet -printtoconsole -conf=/var/btcrpc/bitcoin.conf
