# PERF Server Controller

**PERF Server** is the representation of PERF that is used to communicate with PERF API to manage the **project** and 
**data source** entities. 

The main purpose of a PERF Server controller is to watch changes in the respective Kubernetes Custom Resource (PerfServer CR)
 and to ensure that the state in that resource is applied in EPAM Delivery Platform.
 
Inspect the main steps performed in the reconcile loop on the diagram below:

![arch](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/epmd-edp/perf-operator/master/documentation/puml/perf_server_chain.puml&raw=true)

The diagram above displays the following steps:

- *Ensure Connection to PerfServer*. The controller tries to log in to the specified URL using the spec.ApiUrl and spec.credentialName. 
If connection is not successful, the loop ends up with an error. 
- *Update Status*. The status update in the respective PerfServer CR.
- *Put EDP Component*. Registration of a new component in EDP.

### Related Articles

* [PERF Data Source Controller](../documentation/perf_data_source_controller.md)
* [PERF Integration](../documentation/perf_integration.md)
