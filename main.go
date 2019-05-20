package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"strings"
)

type Zone struct {
	X int
	Y int
}

type Rover struct {
	CurrentX         int
	CurrentY         int
	CurrentDirection string
	Zone             Zone
}

func main() {

	zone := ReadZoneDimensions()

	xStart, yStart := zone.ReadRoverStartingCoOrdinates()

	rover := Rover{xStart, yStart, "E", zone}

	move := ReadRoverMoveOrders()

	rover.MoveRover(move)

	log.Println("------------------------------------")
	log.Println("| Rover ending co-ordinates        |")
	log.Println("------------------------------------")
	log.Printf("X Axis: %v", rover.CurrentX)
	log.Printf("Y Axis: %v", rover.CurrentY)
	log.Printf("Direction facing: %v", rover.CurrentDirection)
}

func (r *Rover) MoveRover(movement string) {

	log.Println("------------------------------------")
	log.Println("| Rover movement log               |")
	log.Println("------------------------------------")

	for _, m := range movement {

		if strings.EqualFold(r.CurrentDirection, "N") {
			if strings.EqualFold(string(m), "M") {
				if r.CurrentX < r.Zone.X {
					log.Println("Moving North 1 space")
					r.CurrentX++
					continue
				} else {
					log.Println("Rover is leaving current zone. Signal lost!")
					panic(errors.New("Rover has exited current zone. Signal lost!"))
				}
			} else if strings.EqualFold(string(m), "R") {
				log.Println("Changing direction to East")
				r.CurrentDirection = "E"
				continue
			} else if strings.EqualFold(string(m), "L") {
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
				} else {
					log.Println("Rover is leaving current zone. Signal lost!")
					panic(errors.New("Rover has exited current zone. Signal lost!"))
				}
			} else if strings.EqualFold(string(m), "R") {
				log.Println("Changing direction to South")
				r.CurrentDirection = "S"
				continue
			} else if strings.EqualFold(string(m), "L") {
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
				} else {
					log.Println("Rover is leaving current zone. Signal lost!")
					panic(errors.New("Rover is leaving current zone. Signal lost!"))
				}
			} else if strings.EqualFold(string(m), "R") {
				log.Println("Changing direction to West")
				r.CurrentDirection = "W"
				continue
			} else if strings.EqualFold(string(m), "L") {
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
				} else {
					log.Println("Rover is leaving current zone. Signal lost!")
					panic(errors.New("Rover is leaving current zone. Signal lost!"))
				}
			} else if strings.EqualFold(string(m), "R") {
				log.Println("Changing direction to North")
				r.CurrentDirection = "N"
				continue
			} else if strings.EqualFold(string(m), "L") {
				log.Println("Changing direction to South")
				r.CurrentDirection = "S"
				continue
			}
		}
	}

	log.Println("------------------------------------")
}

func ReadZoneDimensions() Zone {

	var x int
	var y int

	log.Println("Please enter Zone X axis length:")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Println(err)
	}

	log.Println("Please enter Zone Y axis length:")
	_, err = fmt.Scanf("%d", &y)
	if err != nil {
		fmt.Println(err)
	}

	zone := Zone{x, y}

	return zone
}

func (z *Zone) ReadRoverStartingCoOrdinates() (int, int) {
	var x int
	var y int

	for {
		log.Println("Please enter Rover starting X axis:")
		_, err := fmt.Scanf("%d", &x)
		if err != nil {
			fmt.Println(err)
		}
		if x > z.X {
			log.Printf("This point is outside of the current zone X axis of: %v", z.X)
			continue
		}
		break
	}

	for {
		log.Println("Please enter Rover starting Y axis:")
		_, err := fmt.Scanf("%d", &y)
		if err != nil {
			fmt.Println(err)
		}
		if y > z.Y {
			log.Printf("This point is outside of the current zone Y axis of: %v", z.Y)
			continue
		}
		break
	}
	return x, y
}

func ReadRoverMoveOrders() string {

	var move string
	reader := bufio.NewReader(os.Stdin)

	for {
		log.Println("Please enter Rover movement orders:")

		move, _ = reader.ReadString('\n')
		move = strings.TrimSuffix(move, "\n")

		if !ConfirmRoverMoveOrdersValidity(move) {
			log.Println("Rover movement order input is not correct. Only 'M', 'R' or 'L' permitted")
			continue
		}
		break
	}
	return move
}

func ConfirmRoverMoveOrdersValidity(orders string) bool {
	const alpha = "mlr"

	for _, char := range orders {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}
