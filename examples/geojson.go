package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

func main() {
	resp, err := http.Get("https://docs.mapbox.com/mapbox-gl-js/assets/earthquakes.geojson")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fc, err := geojson.UnmarshalFeatureCollection(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(fc.Features[0].Properties)
	fmt.Println(fc.Features[0].Geometry.(orb.Point))
}
