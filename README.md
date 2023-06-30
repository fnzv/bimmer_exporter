# Bimmer Metrics Exporter

![Project Logo](/imgs/grafana_mybmw_dashboard.png)

A straightforward BMW metrics exporter implemented in Golang and leveraging the <a href="https://github.com/bimmerconnected/bimmer_connected" target="_blank">bimmerconnected library</a>.

The exporter is designed to function with the MyBMW App APIs across various recent car models starting from 2014. However, it's important to note that any changes made by BMW to their APIs or application for recent models may affect its compatibility.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Kubernetes Installation](#installation)
- [References](#references)
- [Contribute](#Contribute)

## Introduction

The code/script within this repository enables you to easily export the metrics provided by the bimmerconnected library, which references the MyBMW application and BMW internal APIs.
<br>
The objective of this project was solely to facilitate the extraction of data in a more practical format, such as Prometheus, and subsequently visualize the outcomes on a Grafana dashboard. Additionally, the implementation includes alerting capabilities.


## Features

The exporter will offer the following available metrics and expose them via HTTP on port 1337, utilizing the default settings and the provided Dockerfile:

```
# curl http://bimmerexporter.ingress.local
location_lat{label="geo"}  00.0000
location_long{label="geo"} 00.0000
checkControlMessages_severity 0
checkControlMessages_type 00
combustionFuelLevel_range 000
gas_fuel_price 1.924

combustionFuelLevel_remainingFuelLiters 00
currentMileages 0000
range 000
doorsState_combinedSecurityState 0
doorsState_combinedState 0
location_heading 000
hood 0
leftFront 0
leftRear 0
rightFront 0
rightRear 0
trunk 0
windowsStatecombinedState 0
windowsStateleftFront 0
windowsStateleftRear 0
windowsStaterightFront 0
windowsStaterightRear 0
```

<br>
Note that some metrics may or may not be available depending on the car model
<br>
<br>

In this specific scenario, we made the decision to expose the exporter on a Kubernetes cluster and automate the rollout and restart of pods to fetch fresh data every 15 minutes.

To accomplish this, you can refer to the provided Terraform code in this repository as an example and make necessary adjustments according to your requirements.

However, if you don't intend to deploy the application on Kubernetes, you can still deploy it as a standalone application using Docker on any host. In this case, please ensure that you execute the "scrape" <a href="https://github.com/bimmerconnected/bimmer_connected" target="_blank">bimmerconnected</a> command prior to running the application.

Please note that the gas fuel price is updated using a random Italian gas brand. If you need to customize it, you can modify the Golang code and replace it with a fixed or dynamically retrieved price from another external endpoint.

The underlying functionality of this application can be described as follows:

1. When the container is started, the <a href="https://github.com/bimmerconnected/bimmer_connected" target="_blank">bimmerconnected</a> client is invoked to download all the data associated with your car. The data is saved in the /app/bmw.json file.

2. The application then loads and parses the relevant data using jq, a command-line JSON processor.

3. Finally, the extracted metrics are made accessible and exposed on port 1337 for retrieval and monitoring.

## Installation

To proceed with the deployment of the Kubernetes deployment, service account, and role necessary for automatic redeployment, you can execute the following Terraform code snippets:
- ![bimmer.tf](/k8s/bimmer.tf)
- ![restart-bimmer.tf](/k8s/restartbimmer.tf)

Please ensure that you provide the required environment variables, EMAIL and PASS (MyBMW App credentials), as follows:

- `var.bmw_email`
- `var.bmw_pass`

Once you have set up the environment variables, you should be ready to proceed. Upon successful deployment, you can verify that your pods are running smoothly. Please note that if the <a href="https://github.com/bimmerconnected/bimmer_connected" target="_blank">bimmerconnected</a> library encounters authentication failures, the pods will crash.

Additionally, a cronjob is set up to restart the deployment every 15 minutes. This ensures that the pods fetch updated data and expose fresh metrics.

Please note that these snippets have been tested with Kubernetes v1.25.5 and Terraform v1.4.6.

Once installed, in order to proceed and scrape those metrics, you will have to set up a scrape job on your Prometheus instance as follows:
```
    additionalScrapeConfigs: 
      - job_name: 'bmw_exporter'
        metrics_path: /
        static_configs:
            - targets: ['http://bimmerexporter.ingress.local']
```

On the Grafana side, you will need to create a dashboard or import the existing one:
- ![dashboard.json](/grafana/dashboard.json)



## References
- https://github.com/bimmerconnected/bimmer_connected
- https://bimmer-connected.readthedocs.io/en/latest/development/reverse_engineering_mybmw.html (This is a particularly interesting read for reverse engineering enthusiasts.)

These resources provide valuable information and insights for further understanding and exploring the project and its underlying technologies.


## Contribute
If you'd like to contribute to the project or report any issues, please feel free to open a pull request or submit an issue on the project's repository.
