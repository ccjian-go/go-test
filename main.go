package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 时间复杂度、空间复杂度、算法稳定性
func main() {
	arr := []int {5,9,2,8,3,7,4,6,1}
	fmt.Printf("%#v",arr)
	QuickSort(arr,0,len(arr)-1)
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
	rand.Seed(time.Now().UnixNano())

	var (
		pivot = arr[startIndex]
		left = startIndex
		right = endIndex
	)
	for left != right {
		for left<right && arr[right]>pivot{
			right--
		}
		for left<right && arr[left]<=pivot{
			left++
		}
		if left<right{
			arr[left],arr[right] = arr[right],arr[left]
		}
	}
	arr[left],arr[startIndex] = arr[startIndex],arr[left]
	return left
}

// 6 堆排序


// 7 归并排序


// -------------------------------------------------------


// 8 计数排序


// 9 桶排序


// 10 基数排序


// 用堆排解决top_k问题




































