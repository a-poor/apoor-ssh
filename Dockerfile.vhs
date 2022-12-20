FROM ghcr.io/charmbracelet/vhs

RUN apt-get update && \
  apt-get upgrade -y && \
  apt-get install -y openssh-client

COPY ./demo.tape /vhs/demo.tape

ENV VHS_PORT=1976
ENV VHS_HOST=0.0.0.0
EXPOSE 1976

CMD [ "/vhs/demo.tape" ]
