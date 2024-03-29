<a name="unreleased"></a>
## [Unreleased]

### Testing

- Add mocks folder to sonar.exclusions [EPMDEDP-11716](https://jiraeu.epam.com/browse/EPMDEDP-11716)
- Refactor, increase test coverage [EPMDEDP-11716](https://jiraeu.epam.com/browse/EPMDEDP-11716)

### Routine

- Update current development version [EPMDEDP-11472](https://jiraeu.epam.com/browse/EPMDEDP-11472)
- Update current development version [EPMDEDP-11826](https://jiraeu.epam.com/browse/EPMDEDP-11826)
- Add templates for github issues [EPMDEDP-11928](https://jiraeu.epam.com/browse/EPMDEDP-11928)
- Upgrade alpine image version to 3.18.0 [EPMDEDP-12085](https://jiraeu.epam.com/browse/EPMDEDP-12085)


<a name="v2.13.0"></a>
## [v2.13.0] - 2023-03-25
### Features

- Added a stub linter [EPMDEDP-10536](https://jiraeu.epam.com/browse/EPMDEDP-10536)
- Updated Operator SDK version [EPMDEDP-11168](https://jiraeu.epam.com/browse/EPMDEDP-11168)
- Updated EDP components [EPMDEDP-11206](https://jiraeu.epam.com/browse/EPMDEDP-11206)

### Code Refactoring

- Apply golangci-lint [EPMDEDP-10629](https://jiraeu.epam.com/browse/EPMDEDP-10629)
- Removed old api [EPMDEDP-11206](https://jiraeu.epam.com/browse/EPMDEDP-11206)

### Routine

- Update current development version [EPMDEDP-10274](https://jiraeu.epam.com/browse/EPMDEDP-10274)
- Update git-chglog for perf-operator [EPMDEDP-11518](https://jiraeu.epam.com/browse/EPMDEDP-11518)
- Bump golang.org/x/net from 0.5.0 to 0.8.0 [EPMDEDP-11578](https://jiraeu.epam.com/browse/EPMDEDP-11578)
- Upgrade alpine image version to 3.16.4 [EPMDEDP-11764](https://jiraeu.epam.com/browse/EPMDEDP-11764)

### Documentation

- Update chart and application version in Readme file [EPMDEDP-11221](https://jiraeu.epam.com/browse/EPMDEDP-11221)


<a name="v2.12.0"></a>
## [v2.12.0] - 2022-08-26
### Features

- Switch to use V1 apis of EDP components [EPMDEDP-10081](https://jiraeu.epam.com/browse/EPMDEDP-10081)
- Download required tools for Makefile targets [EPMDEDP-10105](https://jiraeu.epam.com/browse/EPMDEDP-10105)
- Switch to V1 [EPMDEDP-9220](https://jiraeu.epam.com/browse/EPMDEDP-9220)

### Bug Fixes

- PerfServer CRD metadata updated. [EPMDEDP-9515](https://jiraeu.epam.com/browse/EPMDEDP-9515)

### Code Refactoring

- Use repository and tag for image reference in chart [EPMDEDP-10389](https://jiraeu.epam.com/browse/EPMDEDP-10389)

### Routine

- Upgrade go version to 1.18 [EPMDEDP-10110](https://jiraeu.epam.com/browse/EPMDEDP-10110)
- Fix Jira Ticket pattern for changelog generator [EPMDEDP-10159](https://jiraeu.epam.com/browse/EPMDEDP-10159)
- Update alpine base image to 3.16.2 version [EPMDEDP-10274](https://jiraeu.epam.com/browse/EPMDEDP-10274)
- Update alpine base image version [EPMDEDP-10280](https://jiraeu.epam.com/browse/EPMDEDP-10280)
- Change 'go get' to 'go install' for git-chglog [EPMDEDP-10337](https://jiraeu.epam.com/browse/EPMDEDP-10337)
- Remove VERSION file [EPMDEDP-10387](https://jiraeu.epam.com/browse/EPMDEDP-10387)
- Add gcflags for go build artifact [EPMDEDP-10411](https://jiraeu.epam.com/browse/EPMDEDP-10411)
- Update current development version [EPMDEDP-8832](https://jiraeu.epam.com/browse/EPMDEDP-8832)
- Update chart annotation [EPMDEDP-9515](https://jiraeu.epam.com/browse/EPMDEDP-9515)

### Documentation

- Align README.md [EPMDEDP-10274](https://jiraeu.epam.com/browse/EPMDEDP-10274)


<a name="v2.11.0"></a>
## [v2.11.0] - 2022-05-25
### Features

- Update Makefile changelog target [EPMDEDP-8218](https://jiraeu.epam.com/browse/EPMDEDP-8218)
- Generate CRDs and helm docs automatically [EPMDEDP-8385](https://jiraeu.epam.com/browse/EPMDEDP-8385)

### Bug Fixes

- Fix changelog generation in GH Release Action [EPMDEDP-8468](https://jiraeu.epam.com/browse/EPMDEDP-8468)

### Routine

- Update release CI pipelines [EPMDEDP-7847](https://jiraeu.epam.com/browse/EPMDEDP-7847)
- Add automatic GitHub Release Action [EPMDEDP-7847](https://jiraeu.epam.com/browse/EPMDEDP-7847)
- Populate chart with Artifacthub annotations [EPMDEDP-8049](https://jiraeu.epam.com/browse/EPMDEDP-8049)
- Update changelog [EPMDEDP-8227](https://jiraeu.epam.com/browse/EPMDEDP-8227)
- Update base docker image to alpine 3.15.4 [EPMDEDP-8853](https://jiraeu.epam.com/browse/EPMDEDP-8853)
- Update changelog [EPMDEDP-9185](https://jiraeu.epam.com/browse/EPMDEDP-9185)


<a name="v2.10.0"></a>
## [v2.10.0] - 2021-12-06
### Features

- Provide operator's build information [EPMDEDP-7847](https://jiraeu.epam.com/browse/EPMDEDP-7847)

### Bug Fixes

- Changelog links [EPMDEDP-7847](https://jiraeu.epam.com/browse/EPMDEDP-7847)

### Code Refactoring

- Expand perf-operator role [EPMDEDP-7279](https://jiraeu.epam.com/browse/EPMDEDP-7279)
- Add namespace field in roleRef in OKD RB, aling CRB name [EPMDEDP-7279](https://jiraeu.epam.com/browse/EPMDEDP-7279)
- Replace cluster-wide role/rolebinding to namespaced [EPMDEDP-7279](https://jiraeu.epam.com/browse/EPMDEDP-7279)
- Disable perf integration by default [EPMDEDP-7812](https://jiraeu.epam.com/browse/EPMDEDP-7812)
- Address golangci-lint issues [EPMDEDP-7945](https://jiraeu.epam.com/browse/EPMDEDP-7945)

### Formatting

- Add pointer to MockPerfClient methods [EPMDEDP-7943](https://jiraeu.epam.com/browse/EPMDEDP-7943)

### Routine

- Add changelog generator [EPMDEDP-7847](https://jiraeu.epam.com/browse/EPMDEDP-7847)
- Add codecov report [EPMDEDP-7885](https://jiraeu.epam.com/browse/EPMDEDP-7885)
- Update docker image [EPMDEDP-7895](https://jiraeu.epam.com/browse/EPMDEDP-7895)
- Use custom go build step for operator [EPMDEDP-7932](https://jiraeu.epam.com/browse/EPMDEDP-7932)
- Update go to version 1.17 [EPMDEDP-7932](https://jiraeu.epam.com/browse/EPMDEDP-7932)

### Documentation

- Update the links on GitHub [EPMDEDP-7781](https://jiraeu.epam.com/browse/EPMDEDP-7781)


<a name="v2.9.0"></a>
## [v2.9.0] - 2021-12-03

<a name="v2.8.0"></a>
## [v2.8.0] - 2021-12-03

<a name="v2.7.1"></a>
## [v2.7.1] - 2021-12-03

<a name="v2.7.0"></a>
## v2.7.0 - 2021-12-03

[Unreleased]: https://github.com/epam/edp-perf-operator/compare/v2.13.0...HEAD
[v2.13.0]: https://github.com/epam/edp-perf-operator/compare/v2.12.0...v2.13.0
[v2.12.0]: https://github.com/epam/edp-perf-operator/compare/v2.11.0...v2.12.0
[v2.11.0]: https://github.com/epam/edp-perf-operator/compare/v2.10.0...v2.11.0
[v2.10.0]: https://github.com/epam/edp-perf-operator/compare/v2.9.0...v2.10.0
[v2.9.0]: https://github.com/epam/edp-perf-operator/compare/v2.8.0...v2.9.0
[v2.8.0]: https://github.com/epam/edp-perf-operator/compare/v2.7.1...v2.8.0
[v2.7.1]: https://github.com/epam/edp-perf-operator/compare/v2.7.0...v2.7.1
