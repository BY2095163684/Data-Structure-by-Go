package LinearList

import "fmt"

//双链表
type DNode struct {
	data  int
	prior *DNode //前驱
	next  *DNode //后继
}

// 初始化链表(带头结点)
func createDList() *DNode {
	headNode := new(DNode) //头结点指针
	headNode.data = 0
	headNode.prior = nil //循环链表时指向尾结点,尾结点的.next指向头结点,形成两个方向的循环
	headNode.next = nil
	return headNode
}

// 创建结点
func createDNode(data int) *DNode {
	newNode := new(DNode)
	newNode.data = data
	newNode.next = nil
	return newNode
}

// (头插法)插入结点
func insertDNode(headNode *DNode, data int) {
	//先创建好再插入
	newNode := createDNode(data)
	newNode.prior = headNode
	newNode.next = headNode.next
	headNode.next = newNode
	if newNode.next.prior != nil { //判断链表空与否
		newNode.next.prior = newNode
	}
}

// (尾插法)插入结点
func tailInsertDNode(headNode *DNode, data int) {
	newNode := createDNode(data)
	pInsert := headNode.next
	if pInsert == nil { //空链表的情况
		newNode.prior = headNode
		headNode.next = newNode
	} else {
		for pInsert.next != nil {
			pInsert = pInsert.next
		}
		pInsert.next = newNode
		newNode.prior = pInsert
	}
}

//(正向)遍历
func printDList(headNode *DNode) {
	temp := headNode.next
	for temp != nil {
		fmt.Println(*temp)
		temp = temp.next
	}
}

/*
(逆向)遍历
func rprintList(headNode *DNode) {
	temp := headNode.next
	// for i := 0; i < 3; i++ {//假设正序遍历到了第三个结点
	// 	fmt.Println(*temp)
	// 	temp = temp.next
	// }
	//要求从当前结点逆向遍历
	for temp != headNode {
		fmt.Println(*temp)
		temp = temp.prior
	}
}
*/

/*
func main() {
	headNode := createDList()

	tailInsertDNode(headNode, 1)
	tailInsertDNode(headNode, 2)
	tailInsertDNode(headNode, 3)
	tailInsertDNode(headNode, 4)
	insertDNode(headNode, 0)
	insertDNode(headNode, -1)
	insertDNode(headNode, -2)
	insertDNode(headNode, -3)

	printList(headNode)
}
*/
