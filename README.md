# Golang Linked List

## Search Node Function

```
func (sl *mySingleLinkedList) Search(name string) int {
	ptr := sl.head
	for i := 0; i < sl.size; i++ {
		if ptr.data == name {
			return i
		}
		ptr = ptr.next
	}
	return -1
}
```



Output Program : 


