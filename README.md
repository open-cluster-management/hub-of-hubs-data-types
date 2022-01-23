[comment]: # ( Copyright Contributors to the Open Cluster Management project )

# hub-of-hubs-data-types

[![Go Report Card](https://goreportcard.com/badge/github.com/stolostron/hub-of-hubs-data-types)](https://goreportcard.com/report/github.com/stolostron/hub-of-hubs-data-types)
[![Go Reference](https://pkg.go.dev/badge/github.com/stolostron/hub-of-hubs-data-types.svg)](https://pkg.go.dev/github.com/stolostron/hub-of-hubs-data-types)
[![License](https://img.shields.io/github/license/stolostron/hub-of-hubs-data-types)](/LICENSE)

The data-types module of [Hub-of-Hubs](https://github.com/stolostron/hub-of-hubs).

Go to the [Contributing guide](CONTRIBUTING.md) to learn how to get involved.

## Getting Started

## DeepCopy and CRD generation

1. DeepCopy - run ```make generate```. The file will be generated in the corresponding API folder
2. CRD - run ```make manifests```. The YAML files will be created under <root>/crd folder. The YAML files must be copied then to: https://github.com/stolostron/hub-of-hubs-crds/tree/main/crds and CR example must be updated at: https://github.com/stolostron/hub-of-hubs-crds/tree/main/cr-examples
