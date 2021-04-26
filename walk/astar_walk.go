package main

import (
	"fmt"
	"math"
)

var (
	uses int
)

type Point struct {
	x int
	y int
}

type Game struct {
	field       [][]int
	objIndex    Point
	indexPlayer Point
	//Manhattan: |x1 – x2| + |y1 – y2|.
	manhattan int
}

//Struct node to create the tree
type Node struct {
	Value     Game
	Parent    *Node
	Childrens []*Node
}

func BFS(tree *Node, game *Game) *Node {
	uses = 0
	queue := []*Node{}
	queue = append(queue, tree)
	return BFSUtil(queue)
}

func BFSUtil(queue []*Node) (res *Node) {
	uses += 1
	//time.Sleep(time.Second * 1)
	printWorld(queue[0].Value)

	for i := 0; i < len(queue); i++ {
		queue[i].Value.manhattan = int(math.Abs(float64(queue[i].Value.indexPlayer.x-queue[i].Value.objIndex.x))) + int(math.Abs(float64(queue[i].Value.indexPlayer.y-queue[i].Value.objIndex.y)))
		printWorld(queue[i].Value)
		fmt.Println(queue[i].Value.manhattan)
	}

	minManhattan := queue[0].Value.manhattan
	indexMin := 0
	for i := 0; i < len(queue); i++ {
		if queue[i].Value.manhattan < minManhattan {
			minManhattan = queue[i].Value.manhattan
			indexMin = i
		}

	}

	if len(queue) == 0 {
		return
	}

	if isWin(queue[indexMin].Value) {
		fmt.Println("Result Founded in ", uses, "steps")
		res = queue[indexMin]
		return
	}

	for _, move := range validMovimments(queue[indexMin].Value) {
		//fmt.Println(move)
		newNode := Node{move, queue[0], []*Node{}}
		queue[indexMin].Childrens = append(queue[indexMin].Childrens, &newNode)
		queue = append(queue, &newNode)
	}

	aux := []*Node{}
	for i := 0; i < indexMin; i++ {
		aux = append(aux, queue[i])
	}
	for i := indexMin + 1; i < len(queue); i++ {
		aux = append(aux, queue[i])
	}
	queue = aux
	return BFSUtil(queue)
}

func validMovimments(game Game) (result []Game) {
	result = []Game{}

	defer recover()

	//Up
	if game.indexPlayer.x-1 >= 0 {
		if game.field[game.indexPlayer.x-1][game.indexPlayer.y] == 0 || game.field[game.indexPlayer.x-1][game.indexPlayer.y] == 2 {
			aux := copy(game)
			aux.field[aux.indexPlayer.x-1][aux.indexPlayer.y] = 1
			aux.field[aux.indexPlayer.x][aux.indexPlayer.y] = 0
			aux.indexPlayer = Point{aux.indexPlayer.x - 1, aux.indexPlayer.y}
			result = append(result, aux)
			//fmt.Println("Can move up")
			//printWorld(aux)
		}
	}
	//Down
	if game.indexPlayer.x+1 <= 3 {
		if game.field[game.indexPlayer.x+1][game.indexPlayer.y] == 0 || game.field[game.indexPlayer.x+1][game.indexPlayer.y] == 2 {
			aux := copy(game)
			aux.field[aux.indexPlayer.x+1][aux.indexPlayer.y] = 1
			aux.field[aux.indexPlayer.x][aux.indexPlayer.y] = 0
			aux.indexPlayer = Point{aux.indexPlayer.x + 1, aux.indexPlayer.y}
			result = append(result, aux)
			//fmt.Println("Can move down")
			//printWorld(aux)
		}
	}

	//Right
	if game.indexPlayer.y+1 <= 4 {
		if game.field[game.indexPlayer.x][game.indexPlayer.y+1] == 0 || game.field[game.indexPlayer.x][game.indexPlayer.y+1] == 2 {
			aux := copy(game)
			aux.field[aux.indexPlayer.x][aux.indexPlayer.y+1] = 1
			aux.field[aux.indexPlayer.x][aux.indexPlayer.y] = 0
			aux.indexPlayer = Point{aux.indexPlayer.x, aux.indexPlayer.y + 1}
			result = append(result, aux)
			//fmt.Println("Can move right")
			//printWorld(aux)
		}
	}

	//Left
	if game.indexPlayer.y-1 >= 0 {
		if game.field[game.indexPlayer.x][game.indexPlayer.y-1] == 0 || game.field[game.indexPlayer.x][game.indexPlayer.y-1] == 2 {
			aux := copy(game)
			aux.field[aux.indexPlayer.x][aux.indexPlayer.y-1] = 1
			aux.field[aux.indexPlayer.x][aux.indexPlayer.y] = 0
			aux.indexPlayer = Point{aux.indexPlayer.x, aux.indexPlayer.y - 1}
			result = append(result, aux)
			//fmt.Println("Can move left")
			//printWorld(aux)
		}
	}

	//fmt.Println(result)
	return
}

func copy(game Game) (c Game) {
	c = Game{}
	c.objIndex = game.objIndex
	c.indexPlayer = game.indexPlayer

	c.field = [][]int{}
	for i := 0; i < len(game.field); i++ {
		c.field = append(c.field, []int{})
		for j := 0; j < len(game.field[i]); j++ {
			c.field[i] = append(c.field[i], game.field[i][j])
		}
	}

	return
}

func isWin(game Game) bool {
	return game.field[game.objIndex.x][game.objIndex.y] == 1
}

func printWorld(game Game) {
	for i := 0; i < len(game.field); i++ {
		for j := 0; j < len(game.field[i]); j++ {
			switch game.field[i][j] {
			case 0:
				fmt.Print("-", " ")
				break
			case 1:
				fmt.Print("☻", " ")
				break
			case 2:
				fmt.Print("♥", " ")
				break
			case 3:
				fmt.Print("█", " ")
				break
			}
		}
		fmt.Print("\n")
	}
}

func initWorld(game *Game) {
	*game = Game{[][]int{{1, 0, 3, 0, 0}, {0, 0, 3, 0, 0}, {0, 3, 3, 0, 0}, {0, 0, 0, 0, 2}}, Point{3, 4}, Point{0, 0}, 0}
}

func main() {
	game := Game{}
	initWorld(&game)
	printWorld(game)
	tree := Node{game, nil, []*Node{}}

	result := BFS(&tree, &game)
	printWorld(result.Value)
}
