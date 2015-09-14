package main

import (
	"fmt"
	"github.com/brain-geek/fivez/core"
	"net"
	"os"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "3333"
	TYPE = "tcp"
)

func main() {
	l, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening on " + HOST + ":" + PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	game := core.NewGame()

	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		command := strings.TrimSpace(string(buf))
		command = strings.Trim(command, "\x00")
		command = strings.Trim(command, "\n")
		command = strings.Trim(command, "\r")

		switch command {
		case "up":
			if game.Move(core.UP) == nil {
				game.SpawnVertical()
			}
		case "down":
			if game.Move(core.DOWN) == nil {
				game.SpawnVertical()
			}
		case "left":
			if game.Move(core.LEFT) == nil {
				game.SpawnVertical()
			}
		case "right":
			if game.Move(core.RIGHT) == nil {
				game.SpawnVertical()
			}
		case "quit":
			conn.Write([]byte(fmt.Sprintf("Your score is %d, re-connect to play again!\n", game.GetScore())))
			conn.Close()
			return
		default:
			conn.Write([]byte("Unknown command. Please use 'up', 'down', 'left' or 'right' to make move or 'quit' to stop game.\n\r"))
			// conn.Write([]byte(command))
			// conn.Write([]byte(fmt.Sprintf("%d", len(command))))
		}

		conn.Write(DrawField(game))

		loss, _ := game.HaveLost()

		if loss {
			conn.Write([]byte(fmt.Sprintf("Game has ended. Your score is %d, re-connect to play again!\n", game.GetScore())))
			conn.Close()
		}
	}
}

func DrawField(game *core.Game) []byte {
	var buf []byte

	for _, row := range game.Positions() {
		str := fmt.Sprintf(" %4d | %4d | %4d | %4d \r\n", row[0], row[1], row[2], row[3])
		buf = append(buf, []byte(str)...)
	}

	return buf
}
