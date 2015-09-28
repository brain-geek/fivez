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
	conn.Write([]byte("To play - use commands 'up', 'down', 'left', 'right' to move.\r\n"))
	conn.Write([]byte("To quit - type 'quit' and press Enter.\r\n"))
	conn.Write(DrawField(game))

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
			} else {
				conn.Write([]byte("Move failed.\r\n"))
				conn.Write(DrawField(game))
			}
		case "down":
			if game.Move(core.DOWN) == nil {
				game.SpawnVertical()
			} else {
				conn.Write([]byte("Move failed.\r\n"))
				conn.Write(DrawField(game))
			}
		case "left":
			if game.Move(core.LEFT) == nil {
				game.SpawnVertical()
			} else {
				conn.Write([]byte("Move failed.\r\n"))
				conn.Write(DrawField(game))
			}
		case "right":
			if game.Move(core.RIGHT) == nil {
				game.SpawnVertical()
			} else {
				conn.Write([]byte("Move failed.\r\n"))
				conn.Write(DrawField(game))
			}
		case "quit":
			conn.Write([]byte(fmt.Sprintf("Your score is %d, re-connect to play again!\n", game.GetScore())))
			conn.Close()
			return
		default:
			conn.Write([]byte("Unknown command. Please use 'up', 'down', 'left' or 'right' to make move or 'quit' to stop game.\n\r"))
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
