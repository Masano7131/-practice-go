package main

import (
	"fmt"
	"strconv"
)

//x縦y横
//0 == 空,1 == white, 2==black

var cell = make(map[int]map[int]int)

func main() {
	var count int
	fmt.Scan(&count)
	setupBoard()
	for i := 0; i < count; i++ {
		var (
			color string
			x     int
			y     int
		)
		fmt.Scan(&color, &x, &y)
		switch color {
		case "B":
			reverse(x, y, 2)
		case "W":
			reverse(x, y, 1)
		}
	}
	var b = 0
	var w = 0

	for j := 1; j <= 8; j++ {
		for k := 1; k <= 8; k++ {
			if cell[j][k] == 1 {
				w++
			} else if cell[j][k] == 2 {
				b++
			}
		}
	}

	if w > b {
		fmt.Printf("%s-%s The white won!", format(b), format(w))
	} else if w < b {
		fmt.Printf("%s-%s The black won!", format(b), format(w))
	} else {
		fmt.Printf("%s-%s Draw!", format(b), format(w))
	}
}

func format(s int) string {
	if s < 10 {
		return "0" + strconv.Itoa(s)
	} else {
		return strconv.Itoa(s)
	}
}

func setupBoard() {
	for x := 1; x <= 8; x++ {
		cell[x] = make(map[int]int)
		for y := 1; y <= 8; y++ {
			if (x == 4 && y == 4) || (x == 5 && y == 5) {
				cell[x][y] = 1
			} else if (x == 4 && y == 5) || (x == 5 && y == 4) {
				cell[x][y] = 2
			} else {
				cell[x][y] = 0
			}
		}
	}
}

func reverse(nx, ny, color int) {
	cell[nx][ny] = color
	x := nx
	y := ny
	//top
	for x >= 1 {
		if cell[x-1][y] == 0 {
			break
		} else if cell[x-1][y] == color {
			if x == nx {
				break
			} else {
				for x <= nx {
					cell[x][y] = color
					x++
				}
			}
		}
		x--
	}
	//down
	x = nx
	y = ny
	for x <= 8 {
		if cell[x+1][y] == 0 {
			break
		} else if cell[x+1][y] == color {
			if x == nx {
				break
			} else {
				for x >= nx {
					cell[x][y] = color
					x--
				}
			}
		}
		x++
	}
	x = nx
	y = ny
	//right
	for y >= 1 {
		if cell[x][y-1] == 0 {
			break
		} else if cell[x][y-1] == color {
			if y == ny {
				break
			} else {
				for y <= ny {
					cell[x][y] = color
					y++
				}
			}
		}
		y--
	}
	x = nx
	y = ny
	//left
	for y <= 8 {
		if cell[x][y+1] == 0 {
			break
		} else if cell[x][y+1] == color {
			if y == ny {
				break
			} else {
				for y >= ny {
					cell[x][y] = color
					y--
				}
			}
		}
		y++
	}
	x = nx
	y = ny
	for x >= 1 && y >= 1 {
		if cell[x-1][y-1] == 0 {
			break
		} else if cell[x-1][y-1] == color {
			if x == nx && y == ny {
				break
			} else {
				for x <= nx && y <= ny {
					cell[x][y] = color
					x++
					y++
				}
			}
		}
		x--
		y--
	}
	x = nx
	y = ny
	for x >= 1 && y <= 8 {
		if cell[x-1][y+1] == 0 {
			break
		} else if cell[x-1][y+1] == color {
			if x == nx && y == ny {
				break
			} else {
				for x <= nx && y >= ny {
					cell[x][y] = color
					x++
					y--
				}
			}
		}
		x--
		y++
	}
	x = nx
	y = ny
	for x <= 8 && y >= 1 {
		if cell[x+1][y-1] == 0 {
			break
		} else if cell[x+1][y-1] == color {
			if x == nx && y == ny {
				break
			} else {
				for x >= nx && y <= ny {
					cell[x][y] = color
					x--
					y++
				}
			}
		}
		x++
		y--
	}
	x = nx
	y = ny
	for x <= 8 && y <= 8 {
		if cell[x+1][y+1] == 0 {
			break
		} else if cell[x+1][y+1] == color {
			if x == nx && y == ny {
				break
			} else {
				for x >= nx && y >= ny {
					cell[x][y] = color
					x--
					y--
				}
			}
		}
		x++
		y++
	}
}
