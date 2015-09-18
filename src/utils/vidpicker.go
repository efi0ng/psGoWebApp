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

type History struct {
	Visited []string
}

type MatchedFile struct {
	Name     string
	FileInfo os.FileInfo
}

const (
	player          = "C:/Program Files (x86)/GRETECH/GomPlayer/GOM.exe"
	historyFileName = "vidpicker.history"
)

var (
	fileTypes   = []string{".mp4", ".avi", ".mov", ".wmv", ".flv", ".mkv", ".mpg", ".m4v"}
	endOfStream = MatchedFile{}
)

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
	foundChan := make(chan MatchedFile)

	go walkFiles(".", foundChan)

	filterChan := make(chan MatchedFile)

	go filterFilesByPattern(foundChan, pattern, filterChan)

	historyChan := make(chan MatchedFile)

	go filterByHistory(filterChan, history, historyChan)

	finalChan := historyChan

	// collect results
	selection := []string{}

	for file := <-finalChan; file != endOfStream; file = <-finalChan {
		selection = append(selection, file.Name)
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

func filterByHistory(in chan MatchedFile, history History, out chan MatchedFile) {
	for file := <-in; file != endOfStream; file = <-in {
		if !Contains(history.Visited, file.Name) {
			out <- file
		}
	}

	out <- endOfStream
}

func filterFilesByPattern(in chan MatchedFile, pattern string, out chan MatchedFile) {
	for file := <-in; file != endOfStream; file = <-in {
		name := filepath.Base(file.Name)
		if matched, _ := filepath.Match(pattern, name); matched {
			out <- file
		}
	}

	out <- endOfStream
}

func walkFiles(dir string, out chan MatchedFile) {
	filepath.Walk(dir, func(path string, info os.FileInfo, innerErr error) error {
		if innerErr != nil {
			return nil
		}

		if Contains(fileTypes, filepath.Ext(path)) {
			out <- MatchedFile{path, info}
		}

		return nil
	})

	out <- endOfStream
}
