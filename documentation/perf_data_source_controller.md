# PERF data source Controller

**PERF data source** is the representation of PERF data source in the PERF that is used to create/update data source entity 

The main purpose of a PERF data source controller is to watch changes in the respective Kubernetes Custom Resource (PerfDataSource CR)
 and to ensure that the state in that resource is applied in EPAM Delivery Platform.
 
Inspect the main steps performed in the reconcile loop on the diagram below:

![arch](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/epmd-edp/perf-operator/master/documentation/puml/perf_data_source_chain.puml&raw=true)

The diagram above displays the following steps:

- *Put PerfServer owner to CR*. The controller tries to add PerfServer owner reference to CR. 
- *Create/Update(Activate) data source entity in PERF*. The controller tries to create data source entity in PERF if current doesn't exist, or firstly activates it (if not activated) and then updates entity
- *Update Status*. The status update in the respective PerfDataSource CR.