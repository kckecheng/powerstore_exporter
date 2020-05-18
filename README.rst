PowerStore Exporter
====================

About
------

Prometheus exporter for Dell PowerStore.

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

Run as a docker container
~~~~~~~~~~~~~~~~~~~~~~~~~~

::

  git clone https://github.com/kckecheng/powerstore_exporter.git
  cd powerstore_exporter
  cp config.sample.yml config.yml
  vim config.yml # Tune options
  docker build -t kckecheng/powerstore_exporter .
  docker run -d -p 9100:9100 --name powerstore_exporter kckecheng/powerstore_exporter

Run from CLI
~~~~~~~~~~~~~~

::

  cd powerstore_exporter
  go build
  cp config.sample.yml config.yml
  vim config.yml # Tune options
  ./powerstore_exporter -config config.yml
