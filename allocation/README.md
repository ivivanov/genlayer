# Allocation

## Instructions
```
go test -v
```

## Dependencies
"container/heap" - as part of the standard library

## Explanation
Min Heap would be proper DS since it's adding and removing of elements costs O(logn), while maintaining on top the data center with lowest risk. We can use this mechanics of the min heap to store each fragment in the next data center with lowest risk and fix the data center's risk every time a new fragment is stored. Fixing it's risk will re-arrange the heap...

## Assumptions
Assume that heap DS from standard lib "container/heap" is allowed to be used since it's not external dep