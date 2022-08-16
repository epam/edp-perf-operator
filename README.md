[![codecov](https://codecov.io/gh/epam/edp-perf-operator/branch/master/graph/badge.svg?token=T30RCA4QFD)](https://codecov.io/gh/epam/edp-perf-operator)

# PERF Operator

| :heavy_exclamation_mark: Please refer to [EDP documentation](https://epam.github.io/edp-install/) to get the notion of the main concepts and guidelines. |
| --- |

Get acquainted with the PERF Operator and the installation process as well as the local development, and architecture scheme.

## Overview

PERF Operator is an EDP operator that is responsible for for integration with the Project Performance Board ([PERF board](https://kb.epam.com/display/EPMDMO/Project+Performance+Board)), maintenance, and creation of the data source in the delivery metrics. Operator installation can be applied on two container orchestration platforms: OpenShift and Kubernetes.

_**NOTE:** Operator is platform-independent, that is why there is a unified instruction for deploying._

## Prerequisites

1. Linux machine or Windows Subsystem for Linux instance with [Helm 3](https://helm.sh/docs/intro/install/) installed;
2. Cluster admin access to the cluster;
3. EDP project/namespace is deployed by following the [Install EDP](https://epam.github.io/edp-install/operator-guide/install-edp/) instruction.

## Installation

In order to install the PERF Operator, follow the steps below:

1. To add the Helm EPAMEDP Charts for local client, run "helm repo add":
     ```bash
     helm repo add epamedp https://epam.github.io/edp-helm-charts/stable
     ```
2. Choose available Helm chart version:
     ```bash
     helm search repo epamedp/perf-operator -l
     NAME                           CHART VERSION   APP VERSION     DESCRIPTION
     epamedp/perf-operator          2.11.0          2.11.0          A Helm chart for EDP Perf Operator
     epamedp/perf-operator          2.10.0          2.10.0          A Helm chart for EDP Perf Operator
     ```

    _**NOTE:** It is highly recommended to use the latest released version._

3. Create manually the corresponding secrets:

    3.1 OpenShift:
    ```bash
    oc -n <edp-project> create secret generic <perf.credentialName> --from-literal=username=<username_to_perf> --from-literal=password=<password_to_perf>

    oc -n <edp-project> create secret generic <perf.luminate.credentialName> --from-literal=username=<username_to_luminate> --from-literal=password=<password_to_luminate>
    ```

    3.2 Kubernetes:
    ```bash
    kubectl -n <edp-project> create secret generic <perf.credentialName> --from-literal=username=<username_to_perf> --from-literal=password=<password_to_perf>

    kubectl -n <edp-project> create secret generic <perf.luminate.credentialName> --from-literal=username=<username_to_luminate> --from-literal=password=<password_to_luminate>
    ```
    >_**INFO**: The `<perf.credentialName>` and `<perf.luminate.credentialName>` parameters are described below._

    **IMPORTANT**: Pay attention that at this point, the PERF integration works only on the top of Luminate service so it is required to create the Luminate secret.

4. Create config map with luminate data:

    4.1 OpenShift:
    ```bash
    oc -n <edp-project> create configmap luminatesec-conf --from-literal=apiUrl=<api_url_to_get_luminate_token> --from-literal=credentialName=<perf.luminate.credentialName>
    ```

    4.2 Kubernetes:
    ```bash
    kubectl -n <edp-project> create configmap luminatesec-conf --from-literal=apiUrl=<api_url_to_get_luminate_token> --from-literal=credentialName=<perf.luminate.credentialName>
    ```

5. Create PerfServer CR:

    ```bash
    apiVersion: v2.edp.epam.com/v1
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
    oc -n <edp-project> create secret generic gitlab-admin-password --from-literal=username=<username_to_gitlab> --from-literal=password=<password_to_gitlab>

    oc -n <edp-project> create secret generic jenkins-admin-token --from-literal=username=<username_to_jenkins> --from-literal=password=<password_to_jenkins>

    oc -n <edp-project> create secret generic sonar-admin-password --from-literal=username=<username_to_sonar> --from-literal=password=<password_to_sonar>
    ```

    6.2 Kubernetes:
    ```bash
    kubectl -n <edp-project> create secret generic gitlab-admin-password --from-literal=username=<username_to_gitlab> --from-literal=password=<password_to_gitlab>

    kubectl -n <edp-project> create secret generic jenkins-admin-token --from-literal=username=<username_to_jenkins> --from-literal=password=<password_to_jenkins>

    kubectl -n <edp-project> create secret generic sonar-admin-password --from-literal=username=<username_to_sonar> --from-literal=password=<password_to_sonar>
    ```

7. Full chart parameters available in [deploy-templates/README.md](deploy-templates/README.md).

8. Install operator in the <edp-project> namespace with the helm command; find below the installation command example:
    ```bash
        helm install perf-operator epamedp/perf-operator --version <chart_version> --namespace <edp-project> \
        --set name=perf-operator \
        --set global.edpName=<edp-project> \
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
9. Check the <edp-project> namespace that should contain operator deployment with your operator in a running status.

## Local Development

In order to develop the operator, first set up a local environment. For details, please refer to the [Local Development](https://epam.github.io/edp-install/developer-guide/local-development/) page.

Development versions are also available, please refer to the [snapshot helm chart repository](https://epam.github.io/edp-helm-charts/snapshot/) page.

### Related Articles

* [Architecture Scheme of PERF Operator](documentation/arch.md)
* [Install EDP](https://epam.github.io/edp-install/operator-guide/install-edp/)
* [PERF Data Source Controller](documentation/perf_data_source_controller.md)
* [PERF Server Controller](documentation/perf_server_controller.md)