package LinearList

import "fmt"

type ElemType interface{}

// 结点
type linkNode struct {
	data ElemType
	next *linkNode
}

// 链式队列
type linkQueue struct {
	front *linkNode
	rear  *linkNode
}

// 初始化
func initQueue(Q *linkQueue) {
	Q.front = new(linkNode)
	Q.rear = Q.front
	Q.front.data = nil
	Q.front.next = nil
}

// rear入队
func enQueue(Q *linkQueue, x interface{}) {
	newNode := new(linkNode)
	newNode.data = x
	newNode.next = nil
	Q.rear.next = newNode //新结点插入rear之后
	Q.rear = newNode      //表尾指针后移
}

// front出队
func deQueue(Q *linkQueue, x *interface{}) bool {
	if Q.front == Q.rear { //判空
		return false
	}
	p := Q.front.next     //要出队的是头结点的后一个结点
	*x = p.data           //获取出队的数据
	Q.front.next = p.next //出队对应结点后连接剩余结点
	if Q.rear == p {
		Q.rear = Q.front //判断出队结点是否是最后一个结点
	}
	p.data = nil
	p.next = nil
	return true
}

//遍历打印
func printQueue(Q *linkQueue) {
	p := Q.front
	for p.next != nil {
		fmt.Println(p)
		p = p.next
	}
	fmt.Println(Q.rear)
}

/*
func main() {
	var Q linkQueue
	initQueue(&Q)

	for i := 0; i < 6; i++ {
		enQueue(&Q, i)
	}
	printQueue(&Q)

	var x interface{}
	for i := 0; i < 3; i++ {
		deQueue(&Q, &x)
		fmt.Println(x)
	}
	printQueue(&Q)
}
*/
