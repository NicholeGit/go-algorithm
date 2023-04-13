package spiral_matrix_ii

func generateMatrix(n int) [][]int {
	startx, starty := 0, 0
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	count := 1
	offset := 1
	var loop = n / 2
	for loop > 0 {
		x, y := startx, starty

		//行数不变 列数在变
		for ; y < n-offset; y++ {
			res[x][y] = count
			count++
		}
		//列数不变是y 行数变
		for ; x < n-offset; x++ {
			res[x][y] = count
			count++
		}
		for ; y > starty; y-- {
			res[x][y] = count
			count++
		}
		for ; x > startx; x-- {
			res[x][y] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}
	if n%2 == 1 {
		var center = n / 2
		res[center][center] = n * n
	}
	return res
}
