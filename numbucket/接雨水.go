/* 
 这边按照理解难易程度给上三种解法，最后一种暴力破解最容易理解，可以先看。然后基于暴力破解版进行优化，才有了上面的两种解法。
 其中，难度三星的双指针版本思路与望哥的numbucket.go解法思路类似
 力扣原题链接：https://leetcode-cn.com/problems/trapping-rain-water/submissions/
*/

// 难度三星：双指针版，空间与时间的最优解
func trap(height []int) int {
    length := len(height)
    if length < 2 {
        return 0
    }

    sum := 0
    left, right := 0, length - 1
    leftMax, rightMax := height[0], height[length-1]
    for left < right {
        if height[left] > leftMax {
            leftMax = height[left]
        }

        if height[right] > rightMax {
            rightMax = height[right]
        }

        if leftMax < rightMax {
            sum += leftMax - height[left]
            left++
        } else {
            sum += rightMax - height[right]
            right--
        }
    }
    return sum
}

// 难度二星：dp版，时间O(n)，空间O(n)，提交结果：时间击败100%，空间击败33%
func trap_dp(height []int) int {
    length := len(height)
    if length < 2 {
        return 0
    }

    leftMax := make([]int, length)
    leftMax[0] = 0
    for i := 1; i < length; i++ {
        leftMax[i] = height[i-1]
        if leftMax[i] < leftMax[i-1] {
            leftMax[i] = leftMax[i-1]
        }
    }

    rightMax := make([]int, length)
    rightMax[length-1] = 0
    for i := length-2; i >= 0; i-- {
        rightMax[i] = height[i+1]
        if rightMax[i] < rightMax[i+1] {
            rightMax[i] = rightMax[i+1]
        }
    }

    sum := 0
    for i := 1; i < length - 1; i++ {
        minSider := leftMax[i]
        if minSider > rightMax[i] {
            minSider = rightMax[i]
        }
        if minSider - height[i] > 0 {
            sum += minSider - height[i]
        }
    }

    return sum
}


// 难度一星：暴力版，时间O(n^2)，空间O(1)， 一般是超时啦！提交了下居然没超时，不过时间只击败9%，空间击败100%
func trap_bf(height []int) int {
    length := len(height)
    sum := 0
    for i := 0; i < length; i++ {
        leftMax := 0
        for j := 0; j < i; j++ {
            if height[j] > leftMax {
                leftMax = height[j]
            }
        }

        rightMax := 0
        for j := length - 1; j > i; j-- {
            if height[j] > rightMax {
                rightMax = height[j]
            }
        }

        minSider := leftMax
        if leftMax > rightMax {
            minSider = rightMax
        }
        if minSider - height[i] > 0 {
            sum += minSider - height[i]
        }
    }
    return sum
}
