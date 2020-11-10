# PERF integrating

To integrate with PERF follow those steps:
1. Create PerfServer CR in namespace:
```
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

Once connection is established, you will see information in status:
```
status:
  available: true
  detailed_message: connected
```

2. At this moment, integration with PERf works on top of Luminate service, so you need to create config map **luminatesec-conf** and secret:
```
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
      
>_**NOTE**: Pay attention that secret name must be the same as 'credentialName' property in **luminatesec-conf** config map._

* As soon as everything is configured you will able to add PERF data sources during codebase creation.
