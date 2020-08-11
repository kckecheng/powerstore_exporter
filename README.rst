PowerStore Exporter
====================

About
------

Prometheus exporter for Dell EMC PowerStore.

Usage
-------------

The exporter is able to expose metrics for:

- cluster
- appliance
- node
- fc_port
- eth_port
- volume
- file_system

Run from CLI
~~~~~~~~~~~~~~

::

  cd powerstore_exporter
  go build
  vim config.yml # Tune options
  ./powerstore_exporter -config config.yml

Run as a docker container
~~~~~~~~~~~~~~~~~~~~~~~~~~

::

  git clone https://github.com/kckecheng/powerstore_exporter.git
  cd powerstore_exporter
  vim config.yml # Tune options
  docker run -d -p 9100:9100 --name powerstore_exporter -v $PWD:/etc/powerstore_exporter quay.io/kckecheng/powerstore_exporter

Run on Kubernetes
~~~~~~~~~~~~~~~~~~

::

  export POWERSTORE_SN='fnm1111'
  export POWERSTORE_ADDRESS='192.168.10.10'
  export POWERSTORE_USER='admin'
  export POWERSTORE_PASSWORD='password'
  envsubst < powerstore-exporter.yml | kubectl apply -f -
