package arraysandmatrices

//顺时针打印矩阵
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	step := 0
	size := len(matrix) * len(matrix[0])
	res := make([]int, 0, size)
	top, bottm, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for step < size {
		//从左到右
		for i := left; i <= right && step < size; i++ {
			res = append(res, matrix[top][i])
			step++
		}
		top++
		//从上到下
		for i := top; i <= bottm && step < size; i++ {
			res = append(res, matrix[i][right])
			step++
		}
		right--
		//从右到左
		for i := right; i >= left && step < size; i-- {
			res = append(res, matrix[bottm][i])
			step++
		}
		bottm--
		//从下到上
		for i := bottm; i >= top && step < size; i-- {
			res = append(res, matrix[i][left])
			step++
		}
		left++

	}

	return res
}
