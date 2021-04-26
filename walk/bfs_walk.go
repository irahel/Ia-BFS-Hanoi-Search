package main

import (
	"fmt"
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

	if len(queue) == 0 {
		return
	}

	if isWin(queue[0].Value) {
		fmt.Println("Result Founded in ", uses, "steps")
		res = queue[0]
		return
	}

	for _, move := range validMovimments(queue[0].Value) {
		//fmt.Println(move)
		newNode := Node{move, queue[0], []*Node{}}
		queue[0].Childrens = append(queue[0].Childrens, &newNode)
		queue = append(queue, &newNode)
	}

	return BFSUtil(queue[1:])
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
	*game = Game{[][]int{{1, 0, 3, 0, 0}, {0, 0, 3, 0, 0}, {0, 3, 3, 0, 0}, {0, 0, 0, 0, 2}}, Point{3, 4}, Point{0, 0}}
}

func main() {
	game := Game{}
	initWorld(&game)
	printWorld(game)
	tree := Node{game, nil, []*Node{}}

	result := BFS(&tree, &game)
	printWorld(result.Value)

}
