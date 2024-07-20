package LinearList

import "fmt"

/*

1.创建链表(创建一个表头表示整个链表)
2.创建结点
3.插入结点
4.删除结点
5.打印遍历结点(测试)

*/

// 结点结构体
type Node struct {
	data int
	next *Node
}

// 初始化链表(带头结点)
func createList() *Node {
	headNode := new(Node) //头结点指针
	headNode.data = 0
	headNode.next = nil //循环链表时指向头结点,非空链表的尾结点指针指向头结点
	return headNode
}

// 创建结点
func createNode(data int) *Node {
	newNode := new(Node)
	newNode.data = data
	newNode.next = nil
	return newNode
}

// 打印遍历结点(测试)
func printList(headNode *Node) {
	pMove := headNode.next
	for pMove != nil {
		fmt.Printf("%v\t", pMove.data)
		pMove = pMove.next
	}
}

// (头插法)插入结点
func insertNode(headNode *Node, data int) {
	//先创建好再插入
	newNode := createNode(data)
	newNode.next = headNode.next
	headNode.next = newNode
}

// (尾插法)插入结点
func tailInsertNode(headNode *Node, data int) {
	newNode := createNode(data)
	pInsert := headNode.next
	if pInsert == nil { //空链表的情况
		headNode.next = newNode
	} else {
		for pInsert.next != nil {
			pInsert = pInsert.next
		}
		pInsert.next = newNode
	}
}

// 删除结点
func deleteNode(headNode *Node, posData int) {
	posNode := headNode.next //目标删除结点
	posNodeFront := headNode //目标删除结点的前一个结点
	if posNode == nil {
		fmt.Println("The linkList is empty")
		return
	}
	//遍历删除结点
	for posNode.data != posData {
		posNodeFront = posNode
		posNode = posNodeFront.next
		if posNode == nil {
			fmt.Println("The posData isn't in the list")
			return
		}
	}
	posNodeFront.next = posNode.next
	posNode.data = 0
	posNode.next = nil
}

/*
func main() {
	linkList := createList()
	tailInsertNode(linkList, 1)
	insertNode(linkList, 2)
	tailInsertNode(linkList, 3)
	insertNode(linkList, 4)
	printList(linkList)
	deleteNode(linkList, 2)
	printList(linkList)
}
*/
