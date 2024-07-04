package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	story_json_file_name := parseArgs()
	story := parseJsonStory(story_json_file_name)
	arc := getIntro(story)

	for len(arc.Options) > 0 {
		showArc(arc)
		showOptions(arc.Options)
		choice := getUserChoice(arc.Options)
		chosen_arc := arc.Options[choice].Arc
		arc = getChosenArc(story, chosen_arc)
	}

	fmt.Println("The end")
}

type ArcOption struct {
	Text string
	Arc  string
}

type Arc struct {
	Title   string
	Story   []string
	Options []ArcOption
}

func parseArgs() string {
	story_json_file_name := flag.String("story", "gopher.json", "json file with the story")
	flag.Parse()
	return *story_json_file_name
}

func parseJsonStory(story_json_file_name string) map[string]Arc {
	story_json, err := os.ReadFile(story_json_file_name)
	if err != nil {
		log.Fatalf("Failed to open file %s: %e", story_json_file_name, err)
	}

	var story map[string]Arc
	err = json.Unmarshal(story_json, &story)
	if err != nil {
		log.Fatalf("Failed to unmarshal file %s: %e", story_json_file_name, err)
	}
	return story
}

func getIntro(story map[string]Arc) Arc {
	arc, found := story["intro"]
	if !found {
		log.Fatalf("failed to find intro in story. Make sure your story has 'intro' section")
	}
	return arc
}

func showArc(arc Arc) {
	fmt.Println(arc.Title)
	for _, piece := range arc.Story {
		fmt.Println(piece)
	}
}

func showOptions(options []ArcOption) {
	for i, arcOption := range options {
		fmt.Printf("%d) %s\n", i, arcOption.Text)
	}
}

func getUserChoice(options []ArcOption) int {
	var choice int
	for {
		var rawChoice string
		fmt.Printf("Your choice: ")
		fmt.Scanln(&rawChoice)
		choice, err := strconv.Atoi(rawChoice)
		if err == nil && choice < len(options) {
			fmt.Printf("You choosed %s\n", options[choice].Arc)
			break
		}
		fmt.Printf("invalid input. Try again.")
	}
	return choice
}

func getChosenArc(story map[string]Arc, next_arc string) Arc {
	arc, found := story[next_arc]
	if !found {
		log.Fatalf("Failed to find arc: %s", next_arc)
	}
	return arc
}
