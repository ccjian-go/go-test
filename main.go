package main

import (
	"fmt"
)


// 常见的时间复杂度有：常数阶O(1)，对数阶O(log2n)，线性阶O(n)，线性对数阶O(nlog2n)，平方阶O(n2)，立方阶O(n3)，
// k次方阶O(nk)，指数阶O(2n)。随着问题规模n的不断增大，上述时间复杂度不断增大，算法的执行效率越低。

//计算时间复杂度
// 去掉运行时间中的所有加法常数。
// 只保留最高阶项。
// 如果最高阶项存在且不是1，去掉与这个最高阶相乘的常数得到时间复杂度

//递归情况下的空间复杂度：递归深度为N*每次递归的辅助空间大小，如果每次递归的辅助空间为常数，则空间复杂度为O(N)。
// 对于递归的二分查找，递归深度是log2^n，每次递归的辅助空间为常数，所以空间复杂度为O(log2N)（2为底数下标）

// 非递归情况：在这个过程中，辅助空间为常数级别，所以空间复杂度为O(1)

func main() {
	arr := []int {5,2,8,9,3,7,1,4,6}
	fmt.Printf("%#v",arr)
	//QuickSort(arr,0,len(arr)-1)
	//HeapSort(arr)
	MergerSort(arr, 0, len(arr))
	//MergeSort(arr, 0, len(arr))
	fmt.Printf("%#v",arr)
}


// 1 冒泡排序
// 比较相邻的元素。如果第一个比第二个大，就交换他们两个。 [1]
// 对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数。 [1]
// 针对所有的元素重复以上的步骤，除了最后一个。 [1]
// 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
func BubbleSort(arr []int){
	for i:=0; i<len(arr); i++{
		exchange := false
		for j:=0; j<len(arr)-i-1; j++{
			if arr[j] > arr[j+1]  {
				arr[j], arr[j+1] = arr[j+1],arr[j]
				exchange =true
			}
		}
		if exchange {
			return
		}
	}
}

// 2 选择排序
// 工作原理是每一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
// 以此类推，直到全部待排序的数据元素排完。
func SelectSort(li [] int){
	for i := 0; i<len(li)-1; i++ {
		pos := i
		for j := i+1; j <len(li); j++ {
			if li[pos] < li[j] {
				pos = j
			}
		}
		li[i], li[pos] = li[pos], li[i]
	}
}

// 3 插入排序
// 它的工作原理是通过构建有序序列,对于未排序数据,在已排序序列中从后向前扫描,找到相应位置并插入
// 实现自述：获得将插入的数据，一旦在有序序列中找到相应位置，将后面的数全部右移一位，然后在空出来的位置上放入该插入的数据
func InsertSort(li []int){
	num := 0
	for i:=1; i<len(li); i++ {
		tmp := li[i]
		j := i-1
		for j>=0 && tmp < li[j]  {
			num +=1
			li[j+1] = li[j]
			j--
		}
		li[j+1] = tmp
	}
	fmt.Println(num)
}


// 4 希尔排序
// 直接插入排序的改进版，直接插入排序每次只能移动一位，而希尔排序根据特定的步长，移动距离大大增加。
// 希尔排序就是减少增量的排序方法，给出初始设定一个增量，然后根据这个增量将整个序列进行划分，划分成若干个子序列，然后对每个子序列进行直接插入排序。
// 然后再将增量减少，然后再将整个序列划分成若干个子序列，然后在对每个子序列进行直接插入排序，依次类推，直到增量为1时，
// 即对整个序列进行直接插入排序，因为现在的序列已经基本有序，则直接插入排序的效率较高，这样完成后，我们的整个序列就完全有序。
// 实现自述：直接插入排序算法的改进版，通过分区， 减少重复参与排序的数量
func ShellSort(li []int){
	num := 0
	for gap:=len(li)/2; gap>0; gap /=2 {
		for i:=gap; i<len(li); i+=gap {
			tmp := li[i]
			j := i-gap
			for j>=0 && tmp<li[j] {
				num +=1
				li[j+gap] = li[j]
				j -= gap
			}
			li[j+gap] = tmp
		}
	}
	fmt.Println(num)
}


// 5 快速排序
// 快排是利用分治的思想，如果要排序的数组中下标是从p到r之间的一组数据，我们选择p到r之间的任意一个数据作为分区点（pivot（分区点)),
// 我们遍历p到r之间的数据，将小于pivot放到左边，将大于pivot放到右边，将pivot放到中间，这样经过这一步骤之后，数组p到r之间的数据分成了三部分。
// 前面p到q-1都是小于pivot的，中间pivot， 后面的q+1到r之间都是大于pivot的。

// 双边循环法-实现自述：使用指针找到应该交换位置的值，交换指针对应的值，完成排序。
// 选择一个分区点，（左边的指针）放小的，（右边的指针）放大的，
// 那么当（左边的指针）对应的值小于或等于分区点的值时，指针右移，寻找到不适合该条件的情况时，停止移动，找到的值就是大于分区点的值。等待被交换到右边。
// 同样，指针从右边开始，向左移动，找到小于分区点的值，等到被交换到左边。

func QuickSort(arr []int, startIndex, endIndex int) {
	if startIndex>=endIndex{
		return
	}
	pivotIndex := partition(arr,startIndex,endIndex)
	QuickSort(arr,startIndex,pivotIndex-1)
	QuickSort(arr,pivotIndex+1,endIndex)
}

//双边循环，从右侧开始
func partition(arr []int, startIndex, endIndex int) int {
	var (
		pivot = arr[startIndex] // 选择一个数，作为中间值，亦称分区点，用于比较，小的放它左边，大的放它右边
		left = startIndex
		right = endIndex
	)
	for left != right {
		for left<right && arr[right]>pivot{ //从右边开始向左找，找到值比分区点的值小的位置，它需要被换到分区点的左边，称为左值
			right--
		}
		for left<right && arr[left]<=pivot{ //从左边开始向右找，找到值比分区点的值大的位置，它需要被换到分区点的右边，称为右值
			left++
		}
		if left<right{
			arr[left],arr[right] = arr[right],arr[left] // 交换左右值
		}
	}
	arr[left],arr[startIndex] = arr[startIndex],arr[left] // 将左值换到分区点的左边
	return left 										// 交换后的左值的位置被返回
}

// 6 堆排序
//Base-最大堆、最小堆、堆顶、堆的自我调整（插入节点、删除节点、构建二叉堆）

// 上浮调整
// 下沉调整
// 构建堆

//堆排序
func HeapSort(values []int) {
	buildHeap(values)  // 第一次是把最大值筛选到了第一位
	fmt.Printf("%#v",values)
	for i := len(values); i > 1; i-- {
		values[0], values[i-1] = values[i-1], values[0] // 筛出最大值到顶点结点0，将最大值换到 i-1，后续保持其不变
		adjustHeap(values[:i-1], 0) // i-1及其前面不参与筛选最大值，保持不变
		fmt.Println("the heap is ", values)
	}
}

func buildHeap(values []int) {
	for i := len(values); i >= 0; i-- { //////一定得从后往前调整，
		adjustHeap(values, i)
	}
}

func adjustHeap(values []int, pos int) { ///////// 调整pos位置的结点
	node := pos
	length := len(values)
	for node < length {
		var child int = 0
		if 2*node+2 < length { ///////// 比较左右子节点找出大子节点
			if values[2*node+1] > values[2*node+2] {
				child = 2*node + 1
			} else {
				child = 2*node + 2
			} //////// 选出大子节点
		} else if 2*node+1 < length {
			child = 2*node + 1
		}
		if child > 0 && values[child] > values[node] { ///////// 比较大子节点和本当前节点，当前子节点小则交换位置，node获取三者中最大值
			values[node], values[child] = values[child], values[node] // values[node] 将会是所能接触到的最大值
			node = child // 第二大的值对应的结点比较周围的值进入下个循环
		} else {
			break
		}
	}
}



// 7 归并排序
// 归并排序法（Merge Sort，以下简称MS）是分治法思想运用的一个典范。其主要算法操作可以分为以下步骤：
//Step 1：将n个元素分成两个含n/2元素的子序列
//Step 2：用MS将两个子序列递归排序（最后可以将整个原序列分解成n个子序列）
//Step 3：合并两个已排序好的序列
//首先将数组一份为二，分别为左数组和右数组
//若左数组的长度大于1，那么对左数组实施归并排序
//若右数组的长度大于1， 那么对右数组实施归并排序
//将左右数组进行合并
// 因为是递归，所以要一开始假设是两个数的情况
func MergeSort(arr[]int, a, b int){
	if b-a <= 1{
		return
	}
	c := (a+b)/2
	MergeSort(arr,a,c)
	MergeSort(arr,c,b)
	arrLeft := make([]int, c-a)
	arrRight := make([]int, b-c)
	copy(arrLeft, arr[a:c])
	copy(arrRight, arr[c:b])
	i := 0
	j := 0
	for k:=a; k<b; k++{
		if i == c-a {
			arr[k] = arrRight[j]
			j++
		}else if j == b-c{
			arr[k] = arrLeft[i]
			i++
		}else if arrLeft[i] < arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		}else{
			arr[k] = arrRight[j]
			j++
		}
	}
}









func MergerSort(arr []int, a, b int) {
	if b-a <= 1 {
		return
	}

	c := (a + b) / 2
	MergerSort(arr, a, c)
	MergerSort(arr, c, b)
	arrLeft := make([]int, c-a)
	arrRight := make([]int, b-c)
	copy(arrLeft, arr[a:c])
	copy(arrRight, arr[c:b])
	i := 0 // (a-b)/2的大小
	j := 0
	for k := a; k < b; k++ { //a~c、c~b  a-b步平分到i和j   取j的数 走i的步
		// 利用左右数组排序，再换到目标中
		if i == c-a { // a+i == c 从a走i步到达c  i走到了最后  取j的数 走i的步 则i已经取完 只取j即可
			arr[k] = arrRight[j]
			j++ // j走一步
		} else if j == b-c { // j+c == b 从c走j步到达b  j走到了最后
			arr[k] = arrLeft[i]
			i++ // i走一步
		} else if arrLeft[i] < arrRight[j] { // 先放小的到左边
			arr[k] = arrLeft[i]
			i++ // i走一步
		} else {
			arr[k] = arrRight[j] // 再放大的到右边
			j++ // j走一步
		}
	}
}


// -------------------------------------------------------


// 8 计数排序


// 9 桶排序


// 10 基数排序


// 用堆排解决top_k问题




































