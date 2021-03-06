# PERF Operator

Get acquainted with the PERF Operator and the installation process as well as the local development, 
and architecture scheme.

## Overview

PERF Operator is an EDP operator that is responsible for for integration with the Project Performance Board ([PERF board](https://kb.epam.com/display/EPMDMO/Project+Performance+Board)), 
maintenance, and creation of the data source in the delivery metrics. 
Operator installation can be applied on two container orchestration platforms: OpenShift and Kubernetes.

_**NOTE:** Operator is platform-independent, that is why there is a unified instruction for deploying._

## Prerequisites
1. Linux machine or Windows Subsystem for Linux instance with [Helm 3](https://helm.sh/docs/intro/install/) installed;
2. Cluster admin access to the cluster;
3. EDP project/namespace is deployed by following one of the instructions: [EDP Installation on OpenShift](https://github.com/epam/edp-install/blob/master/documentation/openshift_install_edp.md#edp-installation-on-openshift) or [EDP Installation on Kubernetes](https://github.com/epam/edp-install/blob/master/documentation/kubernetes_install_edp.md#edp-installation-on-kubernetes).

## Installation
In order to install the PERF Operator, follow the steps below:

1. To add the Helm EPAMEDP Charts for local client, run "helm repo add":
     ```bash
     helm repo add epamedp https://chartmuseum.demo.edp-epam.com/
     ```
2. Choose available Helm chart version:
     ```bash
     helm search repo epamedp/perf-operator
     NAME                           CHART VERSION   APP VERSION     DESCRIPTION
     epamedp/perf-operator      v2.6.0                          Helm chart for Golang application/service deplo...
     ```

    _**NOTE:** It is highly recommended to use the latest released version._

3. Create manually the corresponding secrets:  

    3.1 OpenShift:
    ```bash
    oc -n <edp_cicd_project> create secret generic <perf.credentialName> --from-literal=username=<username_to_perf> --from-literal=password=<password_to_perf>
   
    oc -n <edp_cicd_project> create secret generic <perf.luminate.credentialName> --from-literal=username=<username_to_luminate> --from-literal=password=<password_to_luminate>
    ```

    3.2 Kubernetes: 
    ```bash
    kubectl -n <edp_cicd_project> create secret generic <perf.credentialName> --from-literal=username=<username_to_perf> --from-literal=password=<password_to_perf>
   
    kubectl -n <edp_cicd_project> create secret generic <perf.luminate.credentialName> --from-literal=username=<username_to_luminate> --from-literal=password=<password_to_luminate>
    ```
    >_**INFO**: The `<perf.credentialName>` and `<perf.luminate.credentialName>` parameters are described below._
    
    **IMPORTANT**: Pay attention that at this point, the PERF integration works only on the top of Luminate service so it is required to create the Luminate secret.
    
4. Create config map with luminate data:

    4.1 OpenShift:
    ```bash
    oc -n <edp_cicd_project> create configmap luminatesec-conf --from-literal=apiUrl=<api_url_to_get_luminate_token> --from-literal=credentialName=<perf.luminate.credentialName>
    ```

    4.2 Kubernetes: 
    ```bash
    kubectl -n <edp_cicd_project> create configmap luminatesec-conf --from-literal=apiUrl=<api_url_to_get_luminate_token> --from-literal=credentialName=<perf.luminate.credentialName>
    ```
    
5. Create PerfServer CR:

    ```bash
    apiVersion: v2.edp.epam.com/v1alpha1
    kind: PerfServer
    metadata:
      name: <perf_cr_name>
      namespace: <namespace>
    spec:
      apiUrl: '<perf.apiUrl>'
      credentialName: '<perf.credentialName>'
      projectName: '<perf.projectName>'
      rootUrl: '<perf.rootUrl>'
    ```
    
    >_**NOTE**: As soon as the connection is established, the following information will be displayed in the status parameter:_
    >```bash
    >status:
    >  available: true
    >detailed_message: connected
    >```
    
6. Create secrets with administrative rights to integrate the PERF data source with services (_e.g. Jenkins, Sonar, GitLab_):

    6.1 OpenShift:
    ```bash
    oc -n <edp_cicd_project> create secret generic gitlab-admin-password --from-literal=username=<username_to_gitlab> --from-literal=password=<password_to_gitlab>
   
    oc -n <edp_cicd_project> create secret generic jenkins-admin-token --from-literal=username=<username_to_jenkins> --from-literal=password=<password_to_jenkins>
   
    oc -n <edp_cicd_project> create secret generic sonar-admin-password --from-literal=username=<username_to_sonar> --from-literal=password=<password_to_sonar>
    ```

    6.2 Kubernetes: 
    ```bash
    kubectl -n <edp_cicd_project> create secret generic gitlab-admin-password --from-literal=username=<username_to_gitlab> --from-literal=password=<password_to_gitlab>
       
    kubectl -n <edp_cicd_project> create secret generic jenkins-admin-token --from-literal=username=<username_to_jenkins> --from-literal=password=<password_to_jenkins>
   
    kubectl -n <edp_cicd_project> create secret generic sonar-admin-password --from-literal=username=<username_to_sonar> --from-literal=password=<password_to_sonar>
    ```

7. Deploy operator:
  
     Full available chart parameters list:
     
   ```bash
     - chart_version                                 # a version of the PERF operator Helm chart;
     - global.edpName                                # a namespace or a project name (in case of OpenShift);
     - global.platform                               # OpenShift or Kubernetes;
     - image.name                                    # EDP image. The released image can be found on [Dockerhub](https://hub.docker.com/r/epamedp/perf-operator);
     - image.version                                 # EDP tag. The released image can be found on [Dockerhub](https://hub.docker.com/r/epamedp/perf-operator/tags);
     - perf.integration                              # Flag to enable/disable PERF integration (e.g. true/false);
     - perf.name                                     # PerfServer CR name;
     - perf.apiUrl                                   # API URL for development;
     - perf.rootUrl                                  # URL to PERF project;
     - perf.credentialName                           # Name of a secret with credentials to the PERF server;
     - perf.projectName                              # Name of a project in PERF;
     - perf.luminate.enabled                         # Flag to enable/disable Luminate integration (e.g. true/false);
     - perf.luminate.apiUrl                          # API URL for development;
     - perf.luminate.credentialName                  # Name of a secret with Luminate credentials;
   ```
   
8. Install operator in the <edp_cicd_project> namespace with the helm command; find below the installation command example:
    ```bash
        helm install perf-operator epamedp/perf-operator --version <chart_version> --namespace <edp_cicd_project> \
        --set name=perf-operator \
        --set global.edpName=<edp_cicd_project> \
        --set global.platform=openshift \
        --set perf.integration=true \
        --set perf.name=<perf_server_name> \
        --set perf.apiUrl=<api_url> \
        --set perf.rootUrl=<URL_to_project_in_perf> \
        --set perf.credentialName=<credential_name> \
        --set perf.projectName=<project_name_in_perf> \
        --set perf.luminate.enabled=true \
        --set perf.luminate.apiUrl=<api_url> \
        --set perf.luminate.credentialName=<credential_name> \
    ```
9. Check the <edp_cicd_project> namespace that should contain operator deployment with your operator in a running status.

## Local Development

In order to develop the operator, first set up a local environment. For details, please refer to the [Local Development](documentation/local_development.md) page.

### Related Articles

* [Architecture Scheme of PERF Operator](documentation/arch.md)
* [PERF Data Source Controller](documentation/perf_data_source_controller.md)
* [PERF Server Controller](documentation/perf_server_controller.md)