/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    array1 := make([]int, 0)
    array2 := make([]int, 0)
    for  ; l1 != nil; l1 = l1.Next {
        array1 = append(array1, l1.Val)
        
    }
    for ; l2 != nil; l2 = l2.Next {
        array2 = append(array2, l2.Val)
    }
    // 获取最大的长度
    length1 := len(array1)
    length2 := len(array2)
    var add int
    var addMore int
    if (length1 < length2) {
        array2SubArray1Value := length2 - length1
        for  i := length2 -1; i >= 0; i -- {
            array2[i] = array2[i] + add  
            if (i - array2SubArray1Value) >= 0 {
              array2[i] += array1[i - array2SubArray1Value]
            }
            add = 0
            if (array2[i] >= 10) {
                array2[i] = array2[i] - 10
                add = 1
                if (i == 0) {
                   addMore = 1 
                }
            }   
        }
        result := new(ListNode)
        resultHead := result
        if (addMore == 1) {
            result.Val = 1
            result.Next = new(ListNode)
            result = result.Next
        }
        for  i := 0 ; i < length2 ; i++ {
            result.Val = array2[i]
            if (i != length2 - 1) {
               result.Next = new(ListNode)
                result = result.Next 
            } 
        }
        return resultHead
    } else {
        array1SubArray2Value := length1 - length2
        for i := length1 - 1; i >= 0; i -- {
            array1[i] = array1[i] + add   
            if (i - array1SubArray2Value) >= 0 {
                array1[i] += array2[i - array1SubArray2Value]
            }
            add = 0
            if (array1[i] >= 10) {
                array1[i] = array1[i] - 10
                add = 1
                if (i == 0) {
                    addMore = 1
                }
            }
        }
        result := new(ListNode)
        resultHead := result
        if (addMore == 1) {
            result.Val = 1
            result.Next = new(ListNode)
            result = result.Next
        }
        for  i := 0 ; i < length1 ; i++ {
            result.Val = array1[i]
            if (i != length1 - 1) {
               result.Next = new(ListNode)
                result = result.Next 
            } 
        }
        return resultHead
    }  
}