FROM ubuntu:16.04

VOLUME /root/.bitcoin

RUN apt-get update
RUN apt-get install --yes software-properties-common
RUN add-apt-repository --yes ppa:bitcoin/bitcoin
RUN apt-get update

RUN apt-get install --yes bitcoind make

WORKDIR /root

EXPOSE 8334 8334

CMD bitcoind -testnet -printtoconsole -conf=/root/.bitcoin/bitcoin.conf
