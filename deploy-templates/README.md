# perf-operator

![Version: 2.13.0-SNAPSHOT](https://img.shields.io/badge/Version-2.13.0--SNAPSHOT-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.13.0-SNAPSHOT](https://img.shields.io/badge/AppVersion-2.13.0--SNAPSHOT-informational?style=flat-square)

A Helm chart for EDP Perf Operator

**Homepage:** <https://epam.github.io/edp-install/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| epmd-edp | <SupportEPMD-EDP@epam.com> | <https://solutionshub.epam.com/solution/epam-delivery-platform> |
| sergk |  | <https://github.com/SergK> |

## Source Code

* <https://github.com/epam/edp-perf-operator>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| annotations | object | `{}` |  |
| global.edpName | string | `""` | namespace or a project name (in case of OpenShift) |
| global.platform | string | `"openshift"` | platform type that can be "kubernetes" or "openshift" |
| image.repository | string | `"epamedp/perf-operator"` | EDP perf-operator Docker image name. The released image can be found on [Dockerhub](https://hub.docker.com/r/epamedp/perf-operator) |
| image.tag | string | `nil` | EDP perf-operator Docker image tag. The released image can be found on [Dockerhub](https://hub.docker.com/r/epamedp/perf-operator/tags) |
| imagePullPolicy | string | `"IfNotPresent"` |  |
| name | string | `"perf-operator"` | component name |
| nodeSelector | object | `{}` |  |
| perf.apiUrl | string | `"https://perf.delivery.example.com"` | API URL for development |
| perf.credentialName | string | `"perf-user"` | Name of a secret with credentials to the PERF server |
| perf.integration | bool | `false` | Flag to enable/disable PERF integration (e.g. true/false) |
| perf.luminate.apiUrl | string | `"https://api.example.luminatesec.com"` | API URL for development |
| perf.luminate.credentialName | string | `"luminate-secret"` | Name of a secret with Luminate credentials |
| perf.luminate.enabled | bool | `false` | Flag to enable/disable Luminate integration (e.g. true/false) |
| perf.name | string | `"perf"` | PerfServer CR name |
| perf.projectName | string | `"PROJECT-NAME"` | Name of a project in PERF |
| perf.rootUrl | string | `"https://perf.delivery.example.com"` | URL to PERF project |
| resources.limits.memory | string | `"192Mi"` |  |
| resources.requests.cpu | string | `"50m"` |  |
| resources.requests.memory | string | `"64Mi"` |  |
| tolerations | list | `[]` |  |

