FROM nvidia/cuda:10.0-base

WORKDIR /

COPY bin/GMT /usr/local/bin

CMD ["GMT"]