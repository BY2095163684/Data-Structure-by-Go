package LinearList

//数组实现顺序表(静态分配),定义最大顺序表长度
const MAXSIZE = 10

//静态定义顺序表
type SqList1 struct {
	data   [MAXSIZE]int //int数组
	length int          //数组长度
}

//(静态)顺序表初始化
func (L *SqList1) InitList1() {
	for i := 0; i < MAXSIZE; i++ {
		L.data[i] = 0 //顺序表内元素初始值为0
	}
	L.length = 0 //顺序表长度初始值为0
}

//动态定义顺序表
type SqList2 struct {
	data   []int //int切片
	length int
}

//(动态)顺序表初始化
func (L *SqList2) InitList2(capacity int) {
	L.data = make([]int, 0, capacity) //分配容量、分配内存
	L.length = 0
}

//顺序表的插入,以静态定义为例
func ListInsert(L *SqList1, index int, e int) bool {
	//判断index值有效性
	if index < 0 || index > L.length {
		return false
	}
	//判断顺序表是否已满
	if L.length >= MAXSIZE {
		return false
	}
	//第index个及之后的元素后移
	for j := L.length; j >= index; j-- {
		L.data[j] = L.data[j-1]
	}
	L.data[index-1] = e //位置index放入e
	L.length++          //长度+1
	return true
}

//顺序表的删除,以静态定义为例
func ListDelete(L *SqList1, index int, e *int) bool {
	//判断index值有效性
	if index < 1 || index > L.length {
		return false
	}
	*e = L.data[index-1] //取出所删的值
	//第index个及之后的元素前移
	for j := index; j < L.length; j++ {
		L.data[j-1] = L.data[j]
	}
	L.data[L.length-1] = 0 //原顺序表最后一个值置为0
	L.length--             //长度-1
	return true
}

//顺序表(按值)查找
func LocateElem(L SqList1, e int) int {
	for i := 0; i < L.length; i++ {
		if L.data[i] == e {
			return i + 1 //返回其位序
		}
	}
	return 0 //没找到,返回0
}

/*
func main() {
	//声明顺序表
	var L1 SqList1
	var L2 SqList2

	//初始化顺序表
	L1.InitList1()
	L2.InitList2(10)

	//插入操作
	for i := 0; i < 5; i++ {
		L1.data[i] = i
		L1.length++
	}
	ListInsert(&L1, 3, 6)
	fmt.Println(L1)

	//删除操作
	e := 0
	if ListDelete(&L1, 2, &e) {
		fmt.Printf("Finish Deleting %v\n", e)
	} else {
		fmt.Println("Failed Deleting")
	}
	fmt.Println(L1)

	//查找操作
	index := LocateElem(L1, 2)
	fmt.Println(index)
}
*/
