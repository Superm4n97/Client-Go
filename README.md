# Kubernetes verbs
Kubernetes API verbs are different form HTTP verbs (get, create, apply, update, patch, delete and proxy). Kubernetes verbs
are represented in small letters.

# Single resource API
Kubernetes APIs  are single resource API **get, create, apply, update, patch, delete and proxy**. When a client (including kubectl) acts on a set of resource, it sequentially 
performs a series of single resource API request, and merge the response if needed. <br>

But by contrast **watch** and **list** are allowed for getting multiple resources. **Deletecollection** are also allowed for
deleting multiple resources. 

# Dry-run
Dry-run is used when you have to perform HTTP verbs that can modify resources.

# Server side apply

# Resource version
It is a server version of an object. It is a string type, not something numeric. Client api can check whether two resource 
versions are equal or not. You can not compare greater than or less than operation between two resource version. 
<br>
The **get**,**list** and **watch** operation support resource version.
Client find resource version in resources. `metadata.resourceVersion` field identifies the resource version.
### Any
### Most recent
### Not older than
### Exact

# Kubernetes Authentication
[link](https://kubernetes.io/docs/reference/access-authn-authz/authentication/)

# Client-go
[link](https://github.com/kubernetes/client-go)
