package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/nsf/termbox-go"
)

var frame_index = 0
var colors_index = 0

// Frames to play
var frames []string

func reverse(lines []string) []string {
	newLines := make([]string, len(lines))
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		newLines[i], newLines[j] = lines[j], lines[i]
	}
	return newLines
}

func getAnimation(dir string) []string {
	files, err := ioutil.ReadDir("./frames/" + dir)
	if err != nil {
		log.Fatal(err)
	}

	frames := []string{}
	for _, file := range files {
		if !file.IsDir() {
			body, _ := ioutil.ReadFile(fmt.Sprintf("./frames/%s/%s", dir, file.Name()))
			frames = append(frames, string(body))
		}
	}

	return frames
}

func draw(orientation string, animation string, color string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	frames = getAnimation(animation)
	colors := getColor(color)
	lines := strings.Split(frames[frame_index], "\n")

	if orientation == "aussie" {
		lines = reverse(lines)
	}

	for x, line := range lines {
		for y, cell := range line {
			termbox.SetCell(y, x, cell, colors[colors_index], termbox.ColorDefault)
		}
	}

	termbox.Flush()
	frame_index++
	if frame_index == len(frames) {
		frame_index = 0
	}

	colors_index++
	if colors_index == len(colors) {
		colors_index = 0
	}
}
