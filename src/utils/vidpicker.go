package main

// Directory sweeper
import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	player          = "C:/Program Files (x86)/GRETECH/GomPlayer/GOM.exe"
	historyFileName = "vidpicker.history"
)

var (
	fileTypes = []string{".mp4", ".avi", ".mov", ".wmv", ".flv", ".mkv", ".mpg", ".m4v"}
)

type History struct {
	Visited []string
}

func main() {
	pattern := processArgs()
	history := History{}

	// do we have a history? If so, read it
	historyFile, err := os.Open(historyFileName)
	if err == nil {
		// got history
		dec := json.NewDecoder(historyFile)
		err = dec.Decode(&history)

		// close history file now cos we're gonna write to it later
		historyFile.Close()

		if err != nil {
			fmt.Println("error reading history file:", err)
			// stop now before we lose the history
			return
		}
	}

	// go find files
	foundChan := make(chan string)

	go walkFiles(".", foundChan)

	filterChan := make(chan string)

	go filterFilesByPattern(foundChan, pattern, filterChan)

	historyChan := make(chan string)

	go filterByHistory(filterChan, history, historyChan)

	finalChan := historyChan

	// collect results
	selection := []string{}

	for file := <-finalChan; file != ""; file = <-finalChan {
		selection = append(selection, file)
		//TODO: Add a cmdline flag to print the found files
		//fmt.Println(file)
	}

	numFiles := len(selection)
	fmt.Println("\n", numFiles, "files found.")

	if numFiles == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	chosenIndex := rand.Intn(numFiles)
	chosenFile := selection[chosenIndex]
	fmt.Println("\nChose: ", chosenFile)

	cmd := exec.Command(player, chosenFile)
	err = cmd.Start()
	if err != nil {
		fmt.Println("error starting vidplayer:", err)
		return
	}

	// add file to history
	history.Visited = append(history.Visited, chosenFile)

	historyFile.Close()

	if historyFile, err = os.Create(historyFileName); err != nil {
		fmt.Println("error (re)creating history file:", err)
		return
	}
	defer historyFile.Close()
	enc := json.NewEncoder(historyFile)
	enc.Encode(&history)
	return
}

func Contains(slice []string, value string) bool {
	for _, s := range slice {
		if value == s {
			return true
		}
	}

	return false
}

func processArgs() string {
	if len(os.Args) < 2 {
		return "*"
	}

	pattern := os.Args[1]

	if !strings.Contains(pattern, "*") {
		pattern = "*" + pattern + "*"
	}

	return pattern
}

func filterByHistory(in chan string, history History, out chan string) {
	for file := <-in; file != ""; file = <-in {
		if !Contains(history.Visited, file) {
			out <- file
		}
	}

	out <- ""
}

func filterFilesByPattern(in chan string, pattern string, out chan string) {
	for file := <-in; file != ""; file = <-in {
		name := filepath.Base(file)
		if matched, _ := filepath.Match(pattern, name); matched {
			out <- file
		}
	}

	out <- ""
}

func walkFiles(dir string, out chan string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, innerErr error) error {
		if innerErr != nil {
			return nil
		}

		if Contains(fileTypes, filepath.Ext(path)) {
			out <- path
		}

		return nil
	})

	// send end of stream signal
	out <- ""
}
