package main

import "fmt"

type Points struct{ X, Y int }

func (p *Points) showLoc() string {
	return fmt.Sprintf("Location: X:%d Y:%d", p.X, p.Y)
}

func main() {
	p := Points{1, 2}
	fmt.Println(p.showLoc())

}
