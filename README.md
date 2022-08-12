# client-go
Creates a client using client-go. 

### Deployment
#### Create
```shell
> create depoyment
> [Name]
```
#### Update
```shell
> update deployment 
> [Name] [replicas]
```
#### Get
```shell
> get deployment
> [Name]
```
All deployments
```shell
> get deployments
```

#### Delete
```shell
> delete deployment
> [Name]
```

### Service
#### create
```shell
> create service
> [Name]
```

#### Get
```shell
> get services
```
#### Delete
```shell
> delete service
> [Name]
```





## Kubernetes verbs
Kubernetes API verbs are different form HTTP verbs (get, create, apply, update, patch, delete and proxy). Kubernetes verbs
are represented in small letters.

## Single resource API
Kubernetes APIs  are single resource API **get, create, apply, update, patch, delete and proxy**. When a client (including kubectl) acts on a set of resource, it sequentially 
performs a series of single resource API request, and merge the response if needed. <br>

But by contrast **watch** and **list** are allowed for getting multiple resources. **Deletecollection** are also allowed for
deleting multiple resources. 

## Dry-run
Dry-run is used when you have to perform HTTP verbs that can modify resources.

## Server side apply

## Resource version
It is a server version of an object. It is a string type, not something numeric. Client api can check whether two resource 
versions are equal or not. You can not compare greater than or less than operation between two resource version. 
<br>
The **get**,**list** and **watch** operation support resource version.
Client find resource version in resources. `metadata.resourceVersion` field identifies the resource version.
#### Any
#### Most recent
#### Not older than
#### Exact

### [Kubernetes Authentication](https://kubernetes.io/docs/reference/access-authn-authz/authentication/)

## Repositories
There are three repositories provided by kubernetes. 


### [Client-go](https://github.com/kubernetes/client-go)
The _k8s.io/client-go_ library that supports all API types that are officially part of kubernetes. It can be used to execute 
the usual kubernetes REST verbs. The _tool_ package of this library creates actual kubernetes client for with the help of 
kube config file.

### Kubernetes API

### API Machinery


## 3. Basics of client-go
* The Repositories
  * client
  * Kubernetes API
  * API Machinery
* Kubernetes Objects on Go
  * TypeMeta
  * ObjectMeta
* Client Sets
  * Status Subresource
  * Listing and Deletions
  * Watches
  * Client Expansion
  * Client Options
* Informers and Caching
  * Work Queue
* API Machinery in Depth
  * Kinds
  * Resources
  * Rest Mapping
  * Scheme
* Vendoring
  * Glide
  * dep
  * Go Modules

### TypeMeta
TypeMeta contains
* apiVersion
* kind

### ObjectMeta

ObjectMeta basically describe the field of metadata. In JSON or YAML 
these fields are under metadata. 

## 4. Using Custom Resources
* Discovery Information
* Type definitions
* Advanced Features of Custom Resources
  * Validating Custom resources
  * Short Names and Categories
  * Printer Columns
  * Subresources
* A Developer's view on Custom Resources
  * Dynamic Client
  * Typed Clients
  
## 5. Automating Code Generation
* Why Code generation
* Calling the Generator
* Controlling the Generators with Tags
  * Global Tags
  * Local Tags
  * deepcopy-gen Tags
  * runtime.Object and DeepCopyObject
  * client-gen Tags
  * informer-gen and lister-gen
  
## Tag
A tag is a specially formatted Go comment in the following form:
<br>

// +some-tag<br>
// +some-other-tag

There are two kind of tags:
1. Global 
2. Local
