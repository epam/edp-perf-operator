# PERF Data Source Controller

**PERF data source** is the representation of a PERF data source that is used to create/update data source entity. 

The main purpose of a PERF data source controller is to watch changes in the respective Kubernetes Custom Resource (PerfDataSource CR)
 and to ensure that the state in that resource is applied in EPAM Delivery Platform.
 
Inspect the main steps performed in the reconcile loop on the diagram below that di:

![arch](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/epmd-edp/perf-operator/master/documentation/puml/perf_data_source_chain.puml&raw=true)

The diagram above displays the general workflow for the *PerfDataSourceJenkins/Sonar/GitLab* controllers and contains the following steps:

- *Put PerfServer Owner to CR*. The controller tries to add PerfServer owner reference to CR. 
- *Create/Update(Activate) Data Source Entity in PERF*. The controller tries to create data source entity in 
PERF if the current doesn't exist, or the controller activates it (_if not activated_) and then updates the data source entity.
- *Update Status*. The status update in the respective PerfDataSource CR.

### Related Articles

* [PERF Server Controller](../documentation/perf_server_controller.md)