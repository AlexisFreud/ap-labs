package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
	//"Node.go"
)

// Map struct
type Map struct {
	nodes         [][]*Node
	height, width int
	exitDoors     int
	path          string
}

func newMapa(exitDoors int, path string, height int, width int) *Map {
	return &Map{
		height:    height,
		width:     width,
		exitDoors: exitDoors,
		path:      path,
	}
}

func (mapa *Map) initializeMap() {

	fmt.Println("--------------------------------------------\n" +
				"---------------- Map run -------------------\n" +
				"--------------------------------------------\n\n")

	//mapa := Map{nil, 16, 12}
	matrix := readFile(mapa.path, mapa.width, mapa.height)
	mapa.nodes = convertToNodes(matrix, mapa.exitDoors, mapa.width, mapa.height)

	/*
	for i := 0; i < len(mapa.nodes); i++ {
		for j := 0; j < len(mapa.nodes[i]); j++ {
			if mapa.nodes[i][j] == nil {
				fmt.Print(0)
			} else if mapa.nodes[i][j].isExit {
				fmt.Print(2)
			} else {
				fmt.Print(1)
			}
		}
		fmt.Println()
	}
	*/

	/*
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if mapa.nodes[i][j] != nil && mapa.nodes[i][j].isExit {
				fmt.Println(i, j)
			}
		}
	}
	*/
	for i := 0; i < len(mapa.nodes); i++ {
		for j := 0; j < len(mapa.nodes[i]); j++ {
			if mapa.nodes[i][j] == nil {
				fmt.Print("--", 0, "-")
			} else if mapa.nodes[i][j].isExit {
				fmt.Print("--S-")
			} else {
				if mapa.nodes[i][j].stepsToExit < 10 {
					fmt.Print("--", mapa.nodes[i][j].stepsToExit, "-")
				} else {
					fmt.Print("-", mapa.nodes[i][j].stepsToExit, "-")
				}
			}
		}
		fmt.Println()
	}

	fmt.Print("\n\n")
	//fmt.Println("-----------------------------------------\n" +
	//	"--------------End run ------------------\n" +
	//	"-----------------------------------------")
}

func readFile(path string, width int, height int) [][]int {
	// create the arrays needed
	matrix := make([][]int, height)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}

	// create the file and verify errors
	mapFile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Transform the file to an string array
	buffer := csv.NewReader(bufio.NewReader(mapFile))

	// Reads a line per iteration and get the int values
	for i := 0; i < height; i++ {
		line, err := buffer.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for j := 0; j < width; j++ {
			matrix[i][j], err = strconv.Atoi(line[j])
		}
	}
	return matrix
}

func convertToNodes(matrix [][]int, numSalidas int, width int, height int) [][]*Node {
	//Initialize the nodes array
	nodos := make([][]*Node, len(matrix))
	for i := range matrix {
		nodos[i] = make([]*Node, len(matrix[i]))
	}

	// First all Nodes are created without reference to other nodes
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				nodos[i][j] = nil
				continue
			}
			if i == 0 || j == 0 || i == len(matrix)-1 || j == len(matrix[i])-1 {
				nodos[i][j] = NewNode(false, true, 1000)
			} else {
				nodos[i][j] = NewNode(false, false, 1000)
			}
		}
	}

	// Now connects nodes between them
	/*
	 i j --- >
	 |
	 v
	*/
	for i := range nodos {
		for j := range nodos[i] {
			if nodos[i][j] == nil {
				continue
			}
			if i-1 >= 0 {
				nodos[i][j].front = nodos[i-1][j]
			} else {
				nodos[i][j].front = nil
			}
			if i+1 < len(nodos) {
				nodos[i][j].back = nodos[i+1][j]
			} else {
				nodos[i][j].back = nil
			}
			if j-1 >= 0 {
				nodos[i][j].left = nodos[i][j-1]
			} else {
				nodos[i][j].left = nil
			}
			if j+1 < len(nodos[i]) {
				nodos[i][j].right = nodos[i][j+1]
			} else {
				nodos[i][j].right = nil
			}
		}
	}
	generarSalidas(numSalidas, width, height, nodos)
	return nodos
}

func generarSalidas(numSalidas int, width int, height int, nodes [][]*Node) {
	if numSalidas > width*height-height {
		fmt.Println("El número de salidas es muy grande")
		os.Exit(1)
	}

	var r int = numSalidas
	for i := 0; i < numSalidas; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		r = rand.Int()
		if r%4 == 0 {
			// Estaran sobre el eje x hasta arriba
			r = rand.Intn(len(nodes[i]))
			if nodes[0][r] != nil {
				nodes[0][r].isExit = true
				distanciasDeSalida(nodes[0][r], nodes)
			} else {
				i--
				continue
			}
		} else if r%4 == 1 {
			r = rand.Intn(len(nodes[i]))
			if nodes[len(nodes)-1][r] != nil {
				nodes[len(nodes)-1][r].isExit = true
				distanciasDeSalida(nodes[len(nodes)-1][r], nodes)
			} else {
				i--
				continue
			}
		} else if r%4 == 2 {
			r = rand.Intn(len(nodes))
			if nodes[r][0] != nil {
				nodes[r][0].isExit = true
				distanciasDeSalida(nodes[r][0], nodes)
			} else {
				i--
				continue
			}
		} else {
			r = rand.Intn(len(nodes))
			if nodes[r][len(nodes[i])-1] != nil {
				nodes[r][len(nodes[i])-1].isExit = true
				distanciasDeSalida(nodes[r][len(nodes[i])-1], nodes)
			} else {
				i--
				continue
			}
		}
	}
}

func distanciasDeSalida(salida *Node, nodos [][]*Node) {
	cola := NewQueue(10000, 0, 0)
	currentDist := 0
	salida.stepsToExit = currentDist
	cola.add(salida)
	for cola.current < cola.lastPos {
		currentNode := cola.pop()
		// Incrementar currentDist
		if currentDist < currentNode.stepsToExit {
			currentDist++
		}
		nextNodes := currentNode.getNext()
		for i := 0; i < 4; i++ {
			if nextNodes[i] == nil {
				continue
			}
			if currentDist < nextNodes[i].stepsToExit {
				nextNodes[i].stepsToExit = currentDist + 1
				nextNodes[i].nextHop = currentNode
				cola.add(nextNodes[i])
			}
		}
	}
}

func (mapa *Map) getRealNodes() []*Node {
	virtualNodes := make([]*Node, mapa.height*mapa.width)
	numNodes := 0
	for i := 0; i < mapa.height; i++ {
		for j := 0; j < mapa.width; j++ {
			if mapa.nodes[i][j] == nil{
				continue
			}else if mapa.nodes[i][j].isExit == true{
				continue
			} else{
				virtualNodes[numNodes] = mapa.nodes[i][j]
				numNodes++
			}
		}
	}
	realNodes := make([]*Node, numNodes-1)
	for i := 0; i < len(realNodes);i++{
		realNodes[i] = virtualNodes[i]
	}
	return realNodes
}
