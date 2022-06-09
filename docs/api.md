# API Reference

Packages:

- [v2.edp.epam.com/v1](#v2edpepamcomv1)
- [v2.edp.epam.com/v1alpha1](#v2edpepamcomv1alpha1)

# v2.edp.epam.com/v1

Resource Types:

- [PerfDataSourceGitLab](#perfdatasourcegitlab)

- [PerfDataSourceJenkins](#perfdatasourcejenkins)

- [PerfDataSourceSonar](#perfdatasourcesonar)

- [PerfServer](#perfserver)




## PerfDataSourceGitLab
<sup><sup>[↩ Parent](#v2edpepamcomv1 )</sup></sup>






PerfDataSourceGitLab is the Schema for the PerfDataSourceGitLabs API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>v2.edp.epam.com/v1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>PerfDataSourceGitLab</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcegitlabspec">spec</a></b></td>
        <td>object</td>
        <td>
          PerfDataSourceGitLabSpec defines the desired state of PerfDataSourceGitLab<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcegitlabstatus">status</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PerfDataSourceGitLab.spec
<sup><sup>[↩ Parent](#perfdatasourcegitlab)</sup></sup>



PerfDataSourceGitLabSpec defines the desired state of PerfDataSourceGitLab

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>codebaseName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcegitlabspecconfig">config</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>perfServerName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfDataSourceGitLab.spec.config
<sup><sup>[↩ Parent](#perfdatasourcegitlabspec)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>branches</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>repositories</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfDataSourceGitLab.status
<sup><sup>[↩ Parent](#perfdatasourcegitlab)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>status</b></td>
        <td>string</td>
        <td>
          Specifies a current status of PerfDataSourceGitLab.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>

## PerfDataSourceJenkins
<sup><sup>[↩ Parent](#v2edpepamcomv1 )</sup></sup>






PerfDataSourceJenkins is the Schema for the PerfDataSourceJenkinses API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>v2.edp.epam.com/v1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>PerfDataSourceJenkins</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcejenkinsspec">spec</a></b></td>
        <td>object</td>
        <td>
          PerfDataSourceJenkinsSpec defines the desired state of PerfDataSource<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcejenkinsstatus">status</a></b></td>
        <td>object</td>
        <td>
          PerfDataSourceJenkinsStatus defines the observed state of PerfDataSource<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PerfDataSourceJenkins.spec
<sup><sup>[↩ Parent](#perfdatasourcejenkins)</sup></sup>



PerfDataSourceJenkinsSpec defines the desired state of PerfDataSource

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>codebaseName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcejenkinsspecconfig">config</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>perfServerName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfDataSourceJenkins.spec.config
<sup><sup>[↩ Parent](#perfdatasourcejenkinsspec)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>jobNames</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfDataSourceJenkins.status
<sup><sup>[↩ Parent](#perfdatasourcejenkins)</sup></sup>



PerfDataSourceJenkinsStatus defines the observed state of PerfDataSource

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>status</b></td>
        <td>string</td>
        <td>
          Specifies a current status of PerfDataSourceJenkins.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>

## PerfDataSourceSonar
<sup><sup>[↩ Parent](#v2edpepamcomv1 )</sup></sup>






PerfDataSourceSonar is the Schema for the PerfDataSourceSonars API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>v2.edp.epam.com/v1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>PerfDataSourceSonar</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcesonarspec">spec</a></b></td>
        <td>object</td>
        <td>
          PerfDataSourceSonarSpec defines the desired state of PerfDataSourceSonar<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcesonarstatus">status</a></b></td>
        <td>object</td>
        <td>
          PerfDataSourceSonarStatus defines the observed state of PerfDataSourceSonar<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PerfDataSourceSonar.spec
<sup><sup>[↩ Parent](#perfdatasourcesonar)</sup></sup>



PerfDataSourceSonarSpec defines the desired state of PerfDataSourceSonar

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>codebaseName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcesonarspecconfig">config</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>perfServerName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfDataSourceSonar.spec.config
<sup><sup>[↩ Parent](#perfdatasourcesonarspec)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>projectKeys</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfDataSourceSonar.status
<sup><sup>[↩ Parent](#perfdatasourcesonar)</sup></sup>



PerfDataSourceSonarStatus defines the observed state of PerfDataSourceSonar

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>status</b></td>
        <td>string</td>
        <td>
          Specifies a current status of PerfDataSourceSonar.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>

## PerfServer
<sup><sup>[↩ Parent](#v2edpepamcomv1 )</sup></sup>






PerfServer is the Schema for the PerfServers API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>v2.edp.epam.com/v1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>PerfServer</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#perfserverspec">spec</a></b></td>
        <td>object</td>
        <td>
          PerfServerSpec defines the desired state of PerfServer<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#perfserverstatus">status</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PerfServer.spec
<sup><sup>[↩ Parent](#perfserver)</sup></sup>



PerfServerSpec defines the desired state of PerfServer

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>apiUrl</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>credentialName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>projectName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>rootUrl</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfServer.status
<sup><sup>[↩ Parent](#perfserver)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>available</b></td>
        <td>boolean</td>
        <td>
          This flag indicates neither Codebase are initialized and ready to work. Defaults to false.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>last_time_updated</b></td>
        <td>string</td>
        <td>
          Information when the last time the action were performed.<br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>detailed_message</b></td>
        <td>string</td>
        <td>
          Detailed information regarding action result which were performed<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>

# v2.edp.epam.com/v1alpha1

Resource Types:

- [PerfDataSourceGitLab](#perfdatasourcegitlab)

- [PerfDataSourceJenkins](#perfdatasourcejenkins)

- [PerfDataSourceSonar](#perfdatasourcesonar)

- [PerfServer](#perfserver)




## PerfDataSourceGitLab
<sup><sup>[↩ Parent](#v2edpepamcomv1alpha1 )</sup></sup>






PerfDataSourceGitLab is the Schema for the PerfDataSourceGitLabs API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>v2.edp.epam.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>PerfDataSourceGitLab</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcegitlabspec-1">spec</a></b></td>
        <td>object</td>
        <td>
          PerfDataSourceGitLabSpec defines the desired state of PerfDataSourceGitLab<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcegitlabstatus-1">status</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PerfDataSourceGitLab.spec
<sup><sup>[↩ Parent](#perfdatasourcegitlab-1)</sup></sup>



PerfDataSourceGitLabSpec defines the desired state of PerfDataSourceGitLab

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>codebaseName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcegitlabspecconfig-1">config</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>perfServerName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfDataSourceGitLab.spec.config
<sup><sup>[↩ Parent](#perfdatasourcegitlabspec-1)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>branches</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>repositories</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfDataSourceGitLab.status
<sup><sup>[↩ Parent](#perfdatasourcegitlab-1)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>status</b></td>
        <td>string</td>
        <td>
          Specifies a current status of PerfDataSourceGitLab.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>

## PerfDataSourceJenkins
<sup><sup>[↩ Parent](#v2edpepamcomv1alpha1 )</sup></sup>






PerfDataSourceJenkins is the Schema for the PerfDataSourceJenkinses API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>v2.edp.epam.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>PerfDataSourceJenkins</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcejenkinsspec-1">spec</a></b></td>
        <td>object</td>
        <td>
          PerfDataSourceJenkinsSpec defines the desired state of PerfDataSource<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcejenkinsstatus-1">status</a></b></td>
        <td>object</td>
        <td>
          PerfDataSourceJenkinsStatus defines the observed state of PerfDataSource<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PerfDataSourceJenkins.spec
<sup><sup>[↩ Parent](#perfdatasourcejenkins-1)</sup></sup>



PerfDataSourceJenkinsSpec defines the desired state of PerfDataSource

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>codebaseName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcejenkinsspecconfig-1">config</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>perfServerName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfDataSourceJenkins.spec.config
<sup><sup>[↩ Parent](#perfdatasourcejenkinsspec-1)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>jobNames</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfDataSourceJenkins.status
<sup><sup>[↩ Parent](#perfdatasourcejenkins-1)</sup></sup>



PerfDataSourceJenkinsStatus defines the observed state of PerfDataSource

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>status</b></td>
        <td>string</td>
        <td>
          Specifies a current status of PerfDataSourceJenkins.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>

## PerfDataSourceSonar
<sup><sup>[↩ Parent](#v2edpepamcomv1alpha1 )</sup></sup>






PerfDataSourceSonar is the Schema for the PerfDataSourceSonars API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>v2.edp.epam.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>PerfDataSourceSonar</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcesonarspec-1">spec</a></b></td>
        <td>object</td>
        <td>
          PerfDataSourceSonarSpec defines the desired state of PerfDataSourceSonar<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcesonarstatus-1">status</a></b></td>
        <td>object</td>
        <td>
          PerfDataSourceSonarStatus defines the observed state of PerfDataSourceSonar<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PerfDataSourceSonar.spec
<sup><sup>[↩ Parent](#perfdatasourcesonar-1)</sup></sup>



PerfDataSourceSonarSpec defines the desired state of PerfDataSourceSonar

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>codebaseName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#perfdatasourcesonarspecconfig-1">config</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>perfServerName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfDataSourceSonar.spec.config
<sup><sup>[↩ Parent](#perfdatasourcesonarspec-1)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>projectKeys</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfDataSourceSonar.status
<sup><sup>[↩ Parent](#perfdatasourcesonar-1)</sup></sup>



PerfDataSourceSonarStatus defines the observed state of PerfDataSourceSonar

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>status</b></td>
        <td>string</td>
        <td>
          Specifies a current status of PerfDataSourceSonar.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>

## PerfServer
<sup><sup>[↩ Parent](#v2edpepamcomv1alpha1 )</sup></sup>






PerfServer is the Schema for the PerfServers API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>v2.edp.epam.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>PerfServer</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#perfserverspec-1">spec</a></b></td>
        <td>object</td>
        <td>
          PerfServerSpec defines the desired state of PerfServer<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#perfserverstatus-1">status</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PerfServer.spec
<sup><sup>[↩ Parent](#perfserver-1)</sup></sup>



PerfServerSpec defines the desired state of PerfServer

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>apiUrl</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>credentialName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>projectName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>rootUrl</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PerfServer.status
<sup><sup>[↩ Parent](#perfserver-1)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>available</b></td>
        <td>boolean</td>
        <td>
          This flag indicates neither Codebase are initialized and ready to work. Defaults to false.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>last_time_updated</b></td>
        <td>string</td>
        <td>
          Information when the last time the action were performed.<br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>detailed_message</b></td>
        <td>string</td>
        <td>
          Detailed information regarding action result which were performed<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>