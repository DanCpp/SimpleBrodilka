package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
)

type _map struct {
	field [][]byte
	rows  int
	cols  int
}

func (Map _map) printMap() {
	for _, line := range Map.field {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Print("\n")
	}
}
func (Map *_map) setPlayerOnMap(Player *_player) {
	Player.standing_on = Map.field[Player.y_coordinate][Player.x_coordinate]
	Map.field[Player.y_coordinate][Player.x_coordinate] = Player.view_char
}
func (Map *_map) removePlayerOnMap(Player *_player) {
	Map.field[Player.y_coordinate][Player.x_coordinate] = Player.standing_on
}

type _player struct {
	view_char    byte
	x_coordinate int
	y_coordinate int
	standing_on  byte
}

func (Player *_player) move(Map *_map, key keyboard.Key) {
	var x, y *int = &Player.x_coordinate, &Player.y_coordinate

	Map.removePlayerOnMap(Player)

	switch key {
	case keyboard.KeyArrowUp:
		if Map.field[*x][*y-1] != byte('#') {
			*y -= 1
		}
	case keyboard.KeyArrowDown:
		if Map.field[*x][*y+1] != byte('#') {
			*y += 1
		}
	case keyboard.KeyArrowLeft:
		if Map.field[*x-1][*y] != byte('#') {
			*x -= 1
		}
	case keyboard.KeyArrowRight:
		if Map.field[*x+1][*y] != byte('#') {
			*x += 1
		}
	default:
		fmt.Println("Undefined")
	}

	Map.setPlayerOnMap(Player)

}

func main() {
	var Map _map = ReadMap("map.txt")
	var Player _player = InitPlayer(byte('@'), 1, 1)
	Map.setPlayerOnMap(&Player)

	if err := keyboard.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer keyboard.Close()

	//The main loop
	for {
		exec.Command("cls")
		Map.printMap()
		_, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if key == keyboard.KeyEsc {
			fmt.Println("Exit")
			break
		}

		Player.move(&Map, key)

		time.Sleep(200 * time.Millisecond)
	}
}

func ReadMap(path string) (Map _map) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var reader bufio.Reader = *bufio.NewReader(file)

	var index_line = 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		Map.field = append(Map.field, []byte(line))
		index_line += 1
		Map.cols = max(Map.cols, len(line))
	}
	Map.rows = index_line
	return Map
}

func InitPlayer(symbol byte, x_coord, y_coord int) (Player _player) {
	Player.view_char = symbol
	Player.x_coordinate = x_coord
	Player.y_coordinate = y_coord
	Player.standing_on = '.'

	return Player
}
