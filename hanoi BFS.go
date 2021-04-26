package main

import (
	"fmt"
)

const (
	Empty  int = 0
	Lower  int = 1
	Midium int = 2
	Bigger int = 3
)

//Struct node to create the tree
type Node struct {
	Value     Hanoi
	Parent    *Node
	Childrens []*Node
}

type Hanoi struct {
	rodA  []int
	rodB  []int
	rodC  []int
	plays int
}

func BFS(tree *Node, game *Hanoi) *Node {
	queue := []*Node{}
	queue = append(queue, tree)
	return BFSUtil(queue)
}

func BFSUtil(queue []*Node) (res *Node) {
	//fmt.Println("Level ", queue[0].Value.plays)
	//fmt.Println("Prox ", queue[0].Value)
	//fmt.Println("Len ", len(queue))

	//time.Sleep(time.Second * 1)

	if len(queue) == 0 {
		//fmt.Println("End tree")
		return
	}
	//res = append(res, queue[0].Value)
	if isWin(queue[0].Value) {
		fmt.Println("Result Founded")
		res = queue[0]
		return
	}
	//fmt.Println("MOVES")
	for _, move := range validMovimments(queue[0].Value) {
		//fmt.Println(move)
		newNode := Node{move, queue[0], []*Node{}}
		queue[0].Childrens = append(queue[0].Childrens, &newNode)
		queue = append(queue, &newNode)
	}

	return BFSUtil(queue[1:])
}

func initGame(game *Hanoi) (ret Hanoi) {
	//*game = Hanoi{}
	*&game.plays = 0
	*&game.rodA = []int{Bigger, Midium, Lower}
	*&game.rodB = []int{Empty, Empty, Empty}
	*&game.rodC = []int{Empty, Empty, Empty}

	return
}

func isWin(game Hanoi) bool {
	return game.rodC[0] == Bigger && game.rodC[1] == Midium && game.rodC[2] == Lower
}

func validMovimments(game Hanoi) (result []Hanoi) {
	result = []Hanoi{}
	lowerRodA, indexLowerA, index0A := lowerOnRod(game.rodA)
	lowerRodB, indexLowerB, index0B := lowerOnRod(game.rodB)
	lowerRodC, indexLowerC, index0C := lowerOnRod(game.rodC)
	if lowerRodA < lowerRodB {
		auxGame := copy(game)
		auxGame.rodA[indexLowerA] = Empty
		auxGame.rodB[index0B] = lowerRodA
		auxGame.plays += 1
		result = append(result, auxGame)
	}
	if lowerRodA < lowerRodC {
		auxGame := copy(game)
		auxGame.rodA[indexLowerA] = Empty
		auxGame.rodC[index0C] = lowerRodA
		auxGame.plays += 1
		result = append(result, auxGame)
	}
	if lowerRodB < lowerRodA {
		auxGame := copy(game)
		auxGame.rodB[indexLowerB] = Empty
		auxGame.rodA[index0A] = lowerRodB
		auxGame.plays += 1
		result = append(result, auxGame)
	}
	if lowerRodB < lowerRodC {
		auxGame := copy(game)
		auxGame.rodB[indexLowerB] = Empty
		auxGame.rodC[index0C] = lowerRodB
		auxGame.plays += 1
		result = append(result, auxGame)
	}
	if lowerRodC < lowerRodA {
		auxGame := copy(game)
		auxGame.rodC[indexLowerC] = Empty
		auxGame.rodA[index0A] = lowerRodC
		auxGame.plays += 1
		result = append(result, auxGame)
	}
	if lowerRodC < lowerRodB {
		auxGame := copy(game)
		auxGame.rodC[indexLowerC] = Empty
		auxGame.rodB[index0B] = lowerRodC
		auxGame.plays += 1
		result = append(result, auxGame)
	}

	return
}

func lowerOnRod(rod []int) (lower int, indexR int, index0 int) {
	lower = 4
	indexR = 0
	index0 = 0
	used_ := false
	for index, elem := range rod {
		if elem != 0 && elem < lower {
			lower = elem
			indexR = index
		}

		if elem == 0 && !used_ {
			index0 = index
			used_ = true
		}
	}
	if lower == 4 {
		lower = 99
		indexR = 0
		index0 = 0
	}
	return
}

func copy(game Hanoi) (c Hanoi) {
	c = Hanoi{}
	for _, elem := range game.rodA {
		c.rodA = append(c.rodA, elem)
	}
	for _, elem := range game.rodB {
		c.rodB = append(c.rodB, elem)
	}
	for _, elem := range game.rodC {
		c.rodC = append(c.rodC, elem)
	}
	c.plays = game.plays
	return
}

func main() {
	game := Hanoi{}
	initGame(&game)

	tree := Node{game, nil, []*Node{}}
	result := BFS(&tree, &game)
	printResult := []Hanoi{}
	for {
		printResult = append(printResult, result.Value)
		if result.Parent == nil {
			break
		} else {
			result = result.Parent
		}
	}
	for index := len(printResult) - 1; index >= 0; index-- {
		fmt.Println(printResult[index])
	}
}
