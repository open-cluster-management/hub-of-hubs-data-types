[comment]: # ( Copyright Contributors to the Open Cluster Management project )

# hub-of-hubs-data-types

[![Go Report Card](https://goreportcard.com/badge/github.com/open-cluster-management/hub-of-hubs-data-types)](https://goreportcard.com/report/github.com/open-cluster-management/hub-of-hubs-data-types)
[![License](https://img.shields.io/github/license/open-cluster-management/hub-of-hubs-data-types)](/LICENSE)

The data-types module of [Hub-of-Hubs](https://github.com/open-cluster-management/hub-of-hubs).

## DeepCopy and CRD generation

1. DeepCopy - run ```make generate```. The file will be generated in the corresponding API folder
2. CRD - run ```make manifests```. The YAML files will be created under <root>/crd folder
