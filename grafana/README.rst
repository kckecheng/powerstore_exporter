About
=======

Introduction
--------------

This is the sample dashboard for PowerStore with "rollup" as true. For "rollup" as false, the metrics are different, please define a new dashboard with the supported metrics like this sample.

**Notes**: when there are a large num. of volumes or file systems, it is recommended to select just a few instead of all since the dashboard needs more time to render diagrams which eats CPU and memory at the browser side that lead to computer slow response.

Prometheus Configuration
--------------------------

A unique label named "powerstore" needs to be assigned for each monitored PowerStore. The label value will be used to filter PowerStore arrays.

::

  - job_name: 'powerstore_exporter'
    static_configs:
      - targets:
          - 192.168.10.10:9100
        labels:
          powerstore: fnm1010
