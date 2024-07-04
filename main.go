package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ArcOption struct {
	Text string
	Arc  string
}

type Arc struct {
	Title   string
	Story   []string
	Options []ArcOption
}

func main() {
	story_json_file_name := flag.String("story", "gopher.json", "json file with the story")
	flag.Parse()
	story_json, err := os.ReadFile(*story_json_file_name)
	if err != nil {
		log.Fatalf("Failed to open file %s: %e", *story_json_file_name, err)
	}

	var story map[string]Arc
	err = json.Unmarshal(story_json, &story)
	if err != nil {
		log.Fatalf("Failed to unmarshal file %s: %e", *story_json_file_name, err)
	}

	arc, found := story["intro"]
	if !found {
		log.Fatalf("failed to find intro in story. Make sure your story has 'intro' section")
	}
	for {
		fmt.Println(arc.Title)
		for _, piece := range arc.Story {
			fmt.Println(piece)
		}

		if len(arc.Options) == 0 {
			break
		}
		for i, arcOption := range arc.Options {
			fmt.Printf("%d) %s\n", i, arcOption.Text)
		}

		var choice int
		for {
			var rawChoice string
			fmt.Printf("Your choice: ")
			fmt.Scanln(&rawChoice)
			choice, err = strconv.Atoi(rawChoice)
			if err == nil && choice < len(arc.Options) {
				fmt.Printf("You choosed %s\n", arc.Options[choice].Arc)
				break
			}
			fmt.Printf("invalid input. Try again.")
		}
		next_arc := arc.Options[choice].Arc
		arc, found = story[next_arc]
		if !found {
			log.Fatalf("Failed to find arc: %s", next_arc)
		}
	}

	fmt.Println("The end")
}
