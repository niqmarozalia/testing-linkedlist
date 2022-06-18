package main

import (
	"errors"
	"fmt"
)

type node struct {
	data string
	next *node
}

type mySingleLinkedList struct {
	size int
	head *node
}

func (sl *mySingleLinkedList) checkDuplicate(node node)error{
	current := sl.head
	for current.next != nil{
		if current.next.data == node.data{
			return errors.New("tidak boleh sama")
		}
		current = current.next
	}
	return nil
}

func (sl *mySingleLinkedList) addToHead(name string) {
	newNode := &node{
		data: name,
	}
	if sl.head == nil {
		sl.head = newNode
	} else {
		err := sl.checkDuplicate(*newNode)
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		newNode.next = sl.head
		sl.head = newNode
	}
	sl.size++
}
func (sl *mySingleLinkedList) addToTail(name string) {
	newNode := &node{
		data: name,
	}
	if sl.head == nil {
		sl.head = newNode
	} else {
		err := sl.checkDuplicate(*newNode)
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		current := sl.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	sl.size++
}

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

func (sl *mySingleLinkedList) GetAt(pos int) *node{
	ptr := sl.head
	if pos < 0 {
		return ptr
	}
	if pos > (sl.size - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}
func (sl *mySingleLinkedList) DeleteAt(pos int)error{
	if sl.size == 0 {
		fmt.Println("No nodes in list")
		return errors.New("no nodes in list")
	}
	prevNode := sl.GetAt(pos-1)
	if prevNode == nil{
		fmt.Println("Node not found")
		return errors.New("node not found")
	}
	prevNode.next = sl.GetAt(pos).next
	return nil
}

func (s *mySingleLinkedList) iterateList() {
	for node := s.head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}


func newLinkedList() *mySingleLinkedList {
	return &mySingleLinkedList{}
}

func main() {
	singleList := newLinkedList()
	singleList.addToHead("11")
	singleList.addToHead("4")
	singleList.addToHead("12")
	singleList.addToTail("13")
	singleList.addToTail("7")
	singleList.addToTail("70")
	// singleList.DeleteAt(0)
	// singleList.DeleteAt(1)
	// singleList.DeleteAt(2)
	// singleList.DeleteAt(3)
	// singleList.DeleteAt(4)
	singleList.iterateList()
	fmt.Println(singleList.size)
	fmt.Printf("Position of  value '13' is : %v\n",singleList.Search("13"))
}