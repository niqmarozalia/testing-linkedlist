#GOLANG LinkedList

Pada projek ini kita akan menggunakan linkedlist dengan fitur tambahan yaitu : search node dan add after value

Search Node with Value we can use :

Function Search Node 
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

Function Add After Value we can use :
```
func (sl *mySingleLinkedList) Insert(name string, pst int)error{
	newNode := &node{
		data: name,
	}
	if pst < 0 {
		err := errors.New("position cant be less than 0")
		if err != nil{
			fmt.Println(err.Error())
			return nil
		}
	}
	if pst == 0 {
		sl.head = newNode
		sl.size++
		return nil
	}
	if pst > sl.size{
		return nil
	}
	err := sl.checkDuplicate(*newNode) // error check : input cant be the same
		if err != nil{
			fmt.Println(err.Error())
			return nil
		}
	current := sl.GetAt(pst)
	newNode.next = current
	prevNode := sl.GetAt(pst-1)
	prevNode.next = newNode
	sl.size++
	
	return nil
}
```
