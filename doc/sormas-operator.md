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
After reconcilation, Kubernetesativ Objects will be created:
### Sormas Application Deployment Object
This object defines the behavior of kubernetes with the SORMAS application.
This _might_ be replicatable, when the hazelcast cache is outside of the main program.  
#### Options: 
* name
* replicas
* image
* environment variables (from configmap and secret)
  - SORMAS_POSTGRES_USER
  - SORMAS_POSTGRES_PASSWORD
  - SORMAS_SERVER_URL
  - DB_HOST
  - DOMAIN_NAME
  - DB_NAME
  - DB_NAME_AUDIT
  - MAIL_HOST
  - MAIL_FROM
  - SORMAS_VERSION
  - LOCALE
  - EPIDPREFIX
  - SEPARATOR
  - EMAIL_SENDER_ADDRESS
  - EMAIL_SENDER_NAME
  - LATITUDE
  - LONGITUDE
  - MAP_ZOOM
  - TZ
  - JVM_MAX
  - GEO_UUID
  - DEVMODE

#### Healthcheck
```["CMD", "curl", "-f", "-I", "http://localhost:6080/sormas-ui/login"]```
#### Constants
command: ```-c 'config_file=/etc/postgresql/postgresql.conf'```  
replica: 1  

#### Volumes
* /opt/sormas/custom

### Postgres StatefulSet
This Object defines the behavior of k8s towards the database.

#### Options:
* name
* image
* environment variables (from configmap and secret)
  - POSTGRES_PASSWORD
  - DB_NAME
  - DB_NAME_AUDIT
  - SORMAS_POSTGRES_PASSWORD
  - SORMAS_POSTGRES_USER
  - TZ

#### Volumes
* /var/lib/postgresql/data

#### Healthcheck
```["CMD", "psql", "-U", "${SORMAS_POSTGRES_USER}", "-c", "SELECT 1;", "${DB_NAME}"]```

#### Constants
command: ```-c 'config_file=/etc/postgresql/postgresql.conf'```  
replica: 1  

### YAMLs
#### Server Deployment
```yaml

```
#### Postgres StatefulSet

#### Sormas ConfigMap

#### Sormas Secret

#### Sormas Server Service

#### Sormas Postgres Service

#### Server Autoscaler