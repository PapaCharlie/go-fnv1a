# Golang FNV-1a hash

This implementation of the FNV-1a hash algorithm can be used to hash objects. This hash can then be used as the map
key.
For example, the following type can be used to hash unhashable objects:
```golang
map[fnv1a.Hash][]T
```
