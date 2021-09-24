package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	v := make([][]uint8, dy)
	for i := range v {
		r := make([]uint8, dx)
		for j := range r {
			var value uint8
			value = uint8(255 - (i^2* j^2))
			r[j] = value
		}
		v[i] = r
	}
	return v

}

func main() {
	pic.Show(Pic) // cat /tmp/out | base64 -d > /tmp/out.png
}
