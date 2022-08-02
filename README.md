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
