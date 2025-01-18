package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "カサブランカ", Year: 1942, Color: false, Actors: []string{"はんぷ", "てぃー"}},
		{Title: "フレディ", Year: 1950, Color: true, Actors: []string{"よその", "あきこ"}},
		{Title: "ブリジっど", Year: 2010, Color: true, Actors: []string{}},
	}
	// Go構造体 -> JSONはマーシャリング
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	data, err = json.MarshalIndent(movies, "", "\t")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)

	// もちろんすべてアンマーシャルも可能
	var movies2 []Movie
	if err := json.Unmarshal(data, &movies2); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%#v", movies2)
}
