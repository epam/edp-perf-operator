# PERF Integration

To successfully integrate with PERF, follow the steps below:

1. Create PerfServer CR in a namespace:

    ```bash
    apiVersion: v2.edp.epam.com/v1alpha1
    kind: PerfServer
    metadata:
      name: <perf_cr_name>
      namespace: <namespace>
    spec:
      apiUrl: '<api_url_to_perf>'
      credentialName: '<credential_name>'
      projectName: '<project_name>'
      rootUrl: '<url_to_project_in_perf>'
    ```
    
    >_**NOTE**: As soon as the connection is established, the following information will be displayed in the status parameter:_
    >```bash
    >status:
    >  available: true
    >detailed_message: connected
    >```

2. Create the **luminatesec-conf** config map and a secret as at this point, the integration with PERf works on 
the top of Luminate service:

    ```bash
    apiVersion: v1
    data:
      apiUrl: '<luminate_api_hostname>'
      credentialName: <secret_name>
    kind: ConfigMap
    metadata:
      name: luminatesec-conf
      namespace: <namespace>
    
    apiVersion: v1
    data:
      password: <password>    
      username: <username>
    kind: Secret
    metadata:
      name: <secret_name>
      namespace: <namespace>
    type: kubernetes.io/basic-auth
    ```    
    >_**NOTE**: Pay attention that the secret name must be the same as the 'credentialName' property in the **luminatesec-conf** config map._

3. As a result, it becomes available to add PERF data sources during the codebase creation.

### Related Articles

* [PERF Data Source Controller](../documentation/perf_data_source_controller.md)
* [PERF Server Controller](../documentation/perf_server_controller.md)