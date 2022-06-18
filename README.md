# Golang Linkedlist

## Add after Value

```
// fungsi add after value at given position

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

Contoh output program :
![Output add after value](https://github.com/niqmarozalia/testing-linkedlist/blob/add-after-valuee/addaftervalue.png)

