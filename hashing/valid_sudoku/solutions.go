package validsudoku

func helper_gen_large_map(size int) []map[byte]struct{} {

	sets := make([]map[byte]struct{}, size)

	for i := 0; i < size; i++ {
		sets[i] = make(map[byte]struct{}, size)
	}

	return sets
}

func check_data_struct(
	index int,
	val byte,
	data []map[byte]struct{},
) bool {
	sub_data := data[index]
	if val != '.' {
		if _, exist := sub_data[val]; exist {
			return true
		}
		sub_data[val] = struct{}{}
	}
	return false
}

func isValidSudoku(board [][]byte) bool {

	var box_idx int

	column_data_struct := helper_gen_large_map(9)
	grid_data_struct := helper_gen_large_map(9)

	for idx, row := range board {
		for jdx, val := range row {
			row_map := make(map[byte]struct{})
			if val != '.' {
				if _, exist := row_map[val]; exist {
					return false
				}
				row_map[val] = struct{}{}
			}

			box_idx = (idx % 3) + (jdx % 3)

			if check_data_struct(idx, val, column_data_struct) {
				return false
			}
			if check_data_struct(box_idx, val, grid_data_struct) {
				return false
			}
		}
	}

	return true
}
