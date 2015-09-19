package main

// Directory sweeper
import (
	"encoding/json"
	"flag"
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

type Options struct {
	Pattern        string
	MaxDaysOld     int
	MinDaysOld     int
	ListFilesFound bool
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
	options := processArgs()
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

	go filterFilesByPattern(foundChan, options.Pattern, filterChan)

	historyChan := make(chan MatchedFile)

	go filterByHistory(filterChan, history, historyChan)

	var tailChan chan MatchedFile

	if options.MaxDaysOld > 0 || options.MinDaysOld > 0 {
		dateChan := make(chan MatchedFile)
		go filterFilesByDate(historyChan, options.MinDaysOld, options.MaxDaysOld, dateChan)
		tailChan = dateChan
	} else {
		tailChan = historyChan
	}

	// collect results
	selection := []string{}

	for file := <-tailChan; file != endOfStream; file = <-tailChan {
		selection = append(selection, file.Name)

		if options.ListFilesFound {
			fmt.Println(file.Name)
		}
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

func processArgs() Options {
	options := Options{}
	maxDaysOld := flag.Int("maxd", 0, "Maximum days old")
	minDaysOld := flag.Int("mind", 0, "Minimum days old")
	printFound := flag.Bool("l", false, "List files found")

	flag.Parse()

	if flag.NArg() < 1 {
		options.Pattern = "*"
	} else {
		pattern := flag.Arg(0)

		if !strings.Contains(pattern, "*") {
			pattern = "*" + pattern + "*"
		}
		options.Pattern = pattern
	}

	options.MaxDaysOld = *maxDaysOld
	options.MinDaysOld = *minDaysOld
	options.ListFilesFound = *printFound

	return options
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

func filterFilesByDate(in chan MatchedFile, minDaysOld int, maxDaysOld int, out chan MatchedFile) {
	minFileTime := calculateTimeDaysAgo(minDaysOld)
	maxFileTime := calculateTimeDaysAgo(maxDaysOld)

	for file := <-in; file != endOfStream; file = <-in {
		if minDaysOld > 0 && file.FileInfo.ModTime().After(minFileTime) {
			continue
		}
		if maxDaysOld > 0 && file.FileInfo.ModTime().Before(maxFileTime) {
			continue
		}
		out <- file
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

func calculateTimeDaysAgo(days int) time.Time {
	return time.Now().AddDate(0, 0, -days)
}
