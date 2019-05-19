package main

import (
	"github.com/pkg/errors"
	"log"
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

	rover :=  Rover{3,5,"E", Zone{8,7}}

	move := "MLRMRR"

	rover.MoveRover(move)

	log.Printf("X Postition: %v\nY Position: %v\nCurrent Direction: %v", rover.CurrentX, rover.CurrentY, rover.CurrentDirection)
}

func (z *Zone)CreateZone(){
	var ar  []string

	log.Println("--------------------------")
	log.Println("Your zone:")
	log.Println("---------------------------")
	for q := z.Y; q > 0 ; q-- {
		for p := z.X; p > 0; p-- {

			ar = append(ar, "X ")
		}
		log.Println(ar)
		log.Println()
		ar = []string{}
	}
}

func (r *Rover)MoveRover(movement string) {

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
	log.Println("| Rover movement sequence complete |")
	log.Println("------------------------------------")
}