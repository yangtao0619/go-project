package main

import (
	"fmt"
	"strconv"
)

//链表需要有节点和链表两个结构体类型来表示

/*
节点需要有数据和指向下一个节点的指针
 */
type Node struct {
	data string
	next *Node
}

/*
链表需要存储总长度,第一个节点,最后一个节点
 */
type NodeList struct {
	head *Node
	size int
	tail *Node
}

func (list *NodeList) Init() {
	//初始化链表
	list.size = 0
	list.head = nil
	list.tail = nil
}

//在链表最后位置插入一个节点
func (list *NodeList) AppendNode(node *Node) {
	//判断当前链表的长度
	if list.size == 0 {
		list.head = node
		list.tail = node
		list.size = 1
	} else {
		//如果链表长度大于等于1的话需要将tail指针指向新插入的节点
		oldTail := list.tail
		list.tail = node
		oldTail.next = node
		list.size += 1
	}
}

//链表需要有增删改查的方法
func (list *NodeList) InsertNode(node *Node, index int) {
	//判断输入数据的正确性
	if index <= -1 {
		index = 0
	}
	if index > list.size {
		index = list.size
	}
	//插入的时候判断当前链表的长度
	if list.size == 0 {
		list.head = node
		list.tail = node
		list.size = 1
	} else {
		//如果存在多个元素的时候,按index进行插入
		list.insertToList(node, index)
	}
}
func (list *NodeList) insertToList(newNode *Node, index int) {
	if index == 0 {
		oldHead := list.head
		newNode.next = oldHead
		list.head = newNode
	} else if index == list.size {
		//如果链表长度大于等于1的话需要将tail指针指向新插入的节点
		oldTail := list.tail
		list.tail = newNode
		oldTail.next = newNode
	} else {
		//循环到第index - 1个元素,将这个元素原来的next保留,设置新的next为传入的node
		var node *Node
		node = list.head
		for i := 0; i < index-1; i++ {
			node = node.next
		}
		//循环后得到的是第index个元素
		oldNext := node.next
		node.next = newNode
		newNode.next = oldNext
	}
	list.size += 1
}

//删除链表中某个特定元素
func (list *NodeList) DeleteNode(index int) {
	//找到第index个元素
	var node *Node
	node = list.head
	var preNode *Node
	for i := 1; i <= index; i++ {
		if i == index-1 {
			preNode = node.next
		}
		node = node.next
	}
	//要拿到第index个节点和第index-1个节点
	oldNext := node.next
	node.next = nil
	preNode.next = oldNext
	list.size -= 1
}

//得到链表上某个位置的节点
func (list *NodeList) GetNode(index int) *Node {
	var node *Node
	node = list.head
	for i := 1; i <= index; i++ {
		node = node.next
	}
	return node
}

func (list *NodeList) PrintListData() {
	var node *Node
	node = list.head
	for i := 0; i < list.size; i++ {
		fmt.Println("data is", node.data)
		node = node.next
	}
}

func main() {

}

func testInsert() {
	var list NodeList
	list.Init()
	for i := 0; i < 10; i++ {
		list.AppendNode(&Node{"我是链表的第" + strconv.Itoa(i) + "个节点", nil})
	}
	newNode := &Node{"我是插入的节点", nil}
	list.insertToList(newNode, 10)
	list.PrintListData()
}
