package main

import (
	"bufio"
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type Zone struct {
	X int
	Y int
}

type Rover struct {
	CurrentX int
	CurrentY int
	CurrentDirection string
	Zone Zone
}

func main() {

	zone := ReadZoneDimentions()

	rover :=  Rover{3,5,"E", zone}

	move := ReadRoverMoveOrders()

	rover.MoveRover(move)

	rover.Zone.CreateZone()

	log.Println("------------------------------------")
	log.Println("| Rover ending co-ordinates        |")
	log.Println("------------------------------------")
	log.Printf("X Postition: %v\nY Position: %v\nCurrent Direction: %v", rover.CurrentX, rover.CurrentY, rover.CurrentDirection)
}

func (z *Zone)CreateZone(){
	var ar  []string

	log.Println("---------------------------")
	log.Println("| Your zone:              |")
	log.Println("---------------------------")
	for q := z.X; q > 0 ; q-- {
		ar = append(ar, strconv.Itoa(q))
		for p := z.Y; p > 0; p-- {
			ar = append(ar, " ")
		}
		log.Println(ar)
		ar = []string{}
	}
	yAxisAray := []string{"0"}

	for x := 1; z.Y >= x ; x++ {
		yAxisAray = append(yAxisAray, strconv.Itoa(x))
	}
	log.Println(yAxisAray)
}

func (r *Rover)MoveRover(movement string) {

	log.Println("------------------------------------")
	log.Println("| Rover movement log               |")
	log.Println("------------------------------------")

	for _ , m := range movement {

		if strings.EqualFold(r.CurrentDirection, "N") {
			if strings.EqualFold(string(m), "M") {
				if r.CurrentX < r.Zone.X {
				log.Println("Moving North 1 space")
				r.CurrentX++
				continue
				}else {
					log.Println("Rover is leaving current zone. Signal lost!")
					panic(errors.New("Rover has exited current zone. Signal lost!"))
				}
			}else if strings.EqualFold(string(m), "R") {
				log.Println("Changing direction to East")
				r.CurrentDirection = "E"
				continue
			}else if strings.EqualFold(string(m), "L") {
				log.Println("Changing direction to West")
				r.CurrentDirection = "W"
				continue
			}
		}
		if strings.EqualFold(r.CurrentDirection, "E") {
			if strings.EqualFold(string(m), "M") {
				if r.CurrentY < r.Zone.Y {
					log.Println("Moving East 1 space")
					r.CurrentY++
					continue
				}else {
					log.Println("Rover is leaving current zone. Signal lost!")
					panic(errors.New("Rover has exited current zone. Signal lost!"))
				}
			}else if strings.EqualFold(string(m), "R") {
				log.Println("Changing direction to South")
				r.CurrentDirection = "S"
				continue
			}else if strings.EqualFold(string(m), "L") {
				log.Println("Changing direction to North")
				r.CurrentDirection = "N"
				continue
			}
		}
		if strings.EqualFold(r.CurrentDirection, "S") {
			if strings.EqualFold(string(m), "M") {
				if r.CurrentX > 0 {
					log.Println("Moving South 1 space")
					r.CurrentX--
					continue
				}else {
					log.Println("Rover is leaving current zone. Signal lost!")
					panic(errors.New("Rover is leaving current zone. Signal lost!"))
				}
			}else if strings.EqualFold(string(m), "R") {
				log.Println("Changing direction to West")
				r.CurrentDirection = "W"
				continue
			}else if strings.EqualFold(string(m), "L") {
				log.Println("Changing direction to East")
				r.CurrentDirection = "E"
				continue
			}
		}
		if strings.EqualFold(r.CurrentDirection, "W") {
			if strings.EqualFold(string(m), "M") {
				if r.CurrentX > 0 {
					log.Println("Moving West 1 space")
					r.CurrentY--
					continue
				}else {
					log.Println("Rover is leaving current zone. Signal lost!")
					panic(errors.New("Rover is leaving current zone. Signal lost!"))
				}
			}else if strings.EqualFold(string(m), "R") {
				log.Println("Changing direction to North")
				r.CurrentDirection = "N"
				continue
			}else if strings.EqualFold(string(m), "L") {
				log.Println("Changing direction to South")
				r.CurrentDirection = "S"
				continue
			}
		}
	}

	log.Println("------------------------------------")
}

func ReadZoneDimentions() Zone{

	reader := bufio.NewReader(os.Stdin)

	log.Println("Please enter Zone X axis length:")
	xAxis, _ := reader.ReadString('\n')
	x, _ := strconv.Atoi(xAxis)

	log.Println("Please enter Zone Y axis length:")
	yAxis, _ := reader.ReadString('\n')
	y, _ := strconv.Atoi(yAxis)

	zone := Zone{x, y}

	return zone
}

func ReadRoverMoveOrders()string{

	reader := bufio.NewReader(os.Stdin)

		log.Println("Please enter Rover movement orders:")

		move, _ := reader.ReadString('\n')

		return move
}

