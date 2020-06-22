# Intialisation
of the project-directiory with the operator-sdk:
```bash
operator-sdk new sormas-operator --repo=github.com/Netzlink/sormas-operator
```
# Operator build
on a  local maschine:
```bash
make local
```
with a registry:
```bash
make push
```
## Requisites:
* make
* A container engine
  - docker
  - podman
  - buildah
* operator-sdk
# Objects
## Sormas
Create a new API-object via the operator-sdk:
```bash
operator-sdk add api --api-version=sormas.netzlink.com/v1alpha1 --kind=Sormas
```
Since the files for the object should be created, generate a controller:
```bash
operator-sdk add controller --api-version=sormas.netzlink.com/v1alpha1 --kind=Sormas
```
