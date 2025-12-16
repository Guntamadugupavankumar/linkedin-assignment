package stealth

import (
"fmt"
"math/rand"
"time"
)

type Point struct {
X float64
Y float64
}

func MoveMouseHumanLike(from Point, to Point, steps int) {
if steps < 10 {
steps = 10
}

fmt.Println("Simulating human-like mouse movement")

c1 := controlPoint(from, to)
c2 := controlPoint(from, to)

for i := 0; i <= steps; i++ {
    t := float64(i) / float64(steps)
    p := bezier(from, c1, c2, to, t)

    fmt.Printf("Mouse at (%.1f, %.1f)\n", p.X, p.Y)

    time.Sleep(time.Duration(rand.Intn(20)+10) * time.Millisecond)
}


}

func controlPoint(a, b Point) Point {
return Point{
X: (a.X+b.X)/2 + rand.Float64()*100 - 50,
Y: (a.Y+b.Y)/2 + rand.Float64()*100 - 50,
}
}

func bezier(p0, p1, p2, p3 Point, t float64) Point {
u := 1 - t
tt := t * t
uu := u * u
uuu := uu * u
ttt := tt * t

x := uuu*p0.X +
    3*uu*t*p1.X +
    3*u*tt*p2.X +
    ttt*p3.X

y := uuu*p0.Y +
    3*uu*t*p1.Y +
    3*u*tt*p2.Y +
    ttt*p3.Y

return Point{X: x, Y: y}


}