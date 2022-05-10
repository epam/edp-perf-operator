<a name="unreleased"></a>
## [Unreleased]

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

[Unreleased]: https://github.com/epam/edp-perf-operator/compare/v2.10.0...HEAD
[v2.10.0]: https://github.com/epam/edp-perf-operator/compare/v2.9.0...v2.10.0
[v2.9.0]: https://github.com/epam/edp-perf-operator/compare/v2.8.0...v2.9.0
[v2.8.0]: https://github.com/epam/edp-perf-operator/compare/v2.7.1...v2.8.0
[v2.7.1]: https://github.com/epam/edp-perf-operator/compare/v2.7.0...v2.7.1
