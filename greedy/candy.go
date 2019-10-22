//老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
//
//你需要按照以下要求，帮助老师给这些孩子分发糖果：
//
//每个孩子至少分配到 1 个糖果。
//相邻的孩子中，评分高的孩子必须获得更多的糖果。
//那么这样下来，老师至少需要准备多少颗糖果呢？
//
//示例 1:
//
//输入: [1,0,2]
//输出: 5
//解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
//示例 2:
//
//输入: [1,2,2]
//输出: 4
//解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。
//     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。

package greedy

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	} else if len(ratings) == 1 {
		return 1
	}
	var minCandy int
	lenRatings := len(ratings)
	record := make([]int, lenRatings)
	record[0] = 1
	for i := 0; i < lenRatings-1; i++ {
		if ratings[i+1] > ratings[i] {
			record[i+1] = record[i] + 1
		} else {
			record[i+1] = 1
		}
	}
	for i := lenRatings - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && record[i-1] < record[i]+1 {
			record[i-1] = record[i] + 1
		}
	}
	for i := range ratings {
		minCandy += record[i]
	}
	return minCandy
}

//func candy(ratings []int) int {
//    lenRatings := len(ratings)
//    if lenRatings == 0 {
//        return 0
//    }
//    minCandy := 1
//    record := make([]int, lenRatings)
//    record[0] = 1
//
//    for i := 1; i < lenRatings; i ++ {
//        if ratings[i] > ratings[i-1] {s
//            record[i] = record[i-1]+1
//            minCandy += record[i]
//        } else if ratings[i] == ratings[i-1] {
//            record[i] = 1
//            minCandy += 1
//        } else {
//            record[i] = 1
//            minCandy += 1
//            for j := i-1; j >= 0; j -- {
//                if ratings[j] > ratings[j+1] && record[j] <= record[j+1] {
//                    minCandy += record[j+1]+1-record[j]
//                    record[j] = record[j+1]+1
//                } else {
//                    break
//                }
//            }
//        }
//    }
//    //fmt.Printf("record=%v\n", record)
//    return minCandy
//}
