package main

import (
	"fmt"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/planar"
)

func main() {
	midlands := orb.Polygon{
		{
			{-3.4807606735073593, 53.528185328046426},
			{0.5622080764926407, 53.39737237773016},
			{1.2213877639926407, 51.605351024878765},
			{-3.9421864547573593, 51.23539468935261},
			{-3.4807606735073593, 53.528185328046426},
		},
	}

	glasgow := orb.Point{-4.260827064765764,55.86209558485744}
	result := planar.PolygonContains(midlands, glasgow)
	fmt.Printf("Is Glasgow in the Midlands? %t\n", result)

	birmingham := orb.Point{-1.8547841110073593, 52.48374053852205} 
	result = planar.PolygonContains(midlands, birmingham)
	fmt.Printf("Is Birmingham in the Midlands? %t\n", result)
}