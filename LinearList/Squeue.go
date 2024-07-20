package LinearList

const Maxsize = 10

// 顺序队列
type Squeue struct {
	data  [Maxsize]int //静态数组存放数据
	front int          //队尾
	rear  int          //队头
}

// 队列初始化
func InitQueue(Q *Squeue) {
	for i := 0; i < Maxsize; i++ {
		Q.data[i] = 0
	}
	Q.front, Q.rear = 0, 0
}

// 入队(这种判满方式会让队列牺牲一个位置)
func EnQueue(Q *Squeue, x int) bool {
	//判满
	if (Q.rear+1)%Maxsize == Q.front {
		return false
	}
	Q.data[Q.rear] = x
	Q.rear = (Q.rear + 1) % Maxsize //取模运算让队列的逻辑结构变为了环状(循环队列),当Q.rear=9时,下一次入队操作将使Q.rear=0,指回队头
	return true
}

// 出队
func DeQueue(Q *Squeue, x *int) bool {
	//判空
	if Q.front == Q.rear {
		return false
	}
	*x = Q.data[Q.front]
	Q.data[Q.front] = 0
	Q.front = (Q.front + 1) % Maxsize
	return true
}

/*
func main() {
	var Q Squeue
	fmt.Println(Q)

	for i := 0; i < 20; i++ {
		EnQueue(&Q, i)
	}
	fmt.Println(Q)

	x := 0
	for i := 0; i < 4; i++ {
		DeQueue(&Q, &x)
		fmt.Println(x)
	}
	fmt.Println(Q)

	for i := 4; i < 6; i++ {
		EnQueue(&Q, i)
	}
	fmt.Println(Q)
}
*/
