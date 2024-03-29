package tune123

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

type Player struct {
	PlayerName string      //имя плеера (просто имя)
	Binary     string      //имя исполняемого файла
	Path       string      //путь к исполняемым файлам плеера
	Command    chan string //канал для передачи команд процессу плеера
}

func (p *Player) Exec() {
	subProcess := exec.Command(p.Path+"\\"+p.Binary, "-slave", "-playlist", "playlist01.m3u")

	stdin, err := subProcess.StdinPipe()
	if err != nil {
		fmt.Println(err) //replace with logger, or anything you want
	}
	defer stdin.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line

	stdout, err := subProcess.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	defer stdout.Close()
	scanner := bufio.NewScanner(stdout)

	fmt.Println("START")                      //for debug
	if err = subProcess.Start(); err != nil { //Use start, not run
		fmt.Println("An error occured: ", err) //replace with logger, or anything you want
	}

	go func() {
		for {
			fmt.Println("msg")
			msg := <-p.Command
			fmt.Println(msg)
			io.WriteString(stdin, msg+"\n")
			if strings.Compare(msg, "Quit") == 0 {
				break
			}
		}
		fmt.Println("WORKER CLOSED!")
	}()

	go func() {
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
	}()

	subProcess.Wait()
	fmt.Println("END") //for debug
}
