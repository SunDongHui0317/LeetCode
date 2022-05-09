package swordoffer

type p struct {
	x, y int
}

var directions = []p{
	{
		-1, 0,
	}, {
		1, 0,
	}, {
		0, -1,
	}, {
		0, 1,
	},
}

func exist(board [][]byte, word string) bool {
	for i, list := range board {
		for j := range list {
			if search(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, word string, i, j, k int) bool {
	if board[i][j] != word[k] {
		return false
	}
	if k == len(word) -1 { // 单词存在于网格中
		return true
	}
	temp := board[i][j]
	board[i][j] = ' '
	defer func() {
		board[i][j] = temp
	}()
	for _, dir := range directions {
		if newI, newJ := i + dir.x, j + dir.y; newI >=0 && newJ >= 0 && newI < len(board) && newJ < len(board[0]) && board[newI][newJ] != ' ' {
			if search(board, word, newI, newJ, k+1) {
				return true
			}
		}
	}
	return false
}
