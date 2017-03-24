package main

// Vid Picker
// TODO: Make case insensitive
// TODO: Add history review interface?
// TODO: Add date to history
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
	ShowVersion    bool
	Prune          bool
	Json           bool
}

type SearchResult struct {
	Matches    int    `json:"matches"`
	ChosenFile string `json:"chosenFile"`
}

const (
	player           = "C:/Program Files/Combined Community Codec Pack 64bit/MPC/mpc-hc64.exe"
	historyFileName  = "vidpicker.history"
	vidPickerVersion = "1.0.170324.0"
)

var (
	fileTypes   = []string{".mp4", ".avi", ".mov", ".wmv", ".flv", ".mkv", ".mpg", ".m4v"}
	endOfStream = MatchedFile{}
)

func main() {
	options := processArgs()

	if options.ShowVersion {
		fmt.Println("VidPicker version ", vidPickerVersion)
		return
	}

	sharedHistory, err := readHistory(historyFileName)

	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading history file:", err)
		// stop now before we lose the history
		return
	}

	if options.Prune {
		sharedHistory = pruneHistory(sharedHistory)
	} else {
		pickVideo(options, &sharedHistory)
	}

	writeHistory(sharedHistory, historyFileName)
}

func pickVideo(options Options, historyPtr *History) {

	// build filter chain that will find the files
	outChan := buildFilterChain(options, *historyPtr)

	// collect results
	selection := []string{}

	for file := <-outChan; file != endOfStream; file = <-outChan {
		selection = append(selection, file.Name)

		if options.ListFilesFound {
			fmt.Println(file.Name)
		}
	}

	numFiles := len(selection)
	if !options.Json {
		fmt.Println("\n", numFiles, "files found.")
	}

	if numFiles == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	chosenIndex := rand.Intn(numFiles)
	chosenFile := selection[chosenIndex]
	if !options.Json {
		fmt.Println("\nChose: ", chosenFile)
	}

	cmd := exec.Command(player, chosenFile)
	err := cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error starting vidplayer:", err)
		return
	}

	addFileToHistory(historyPtr, chosenFile)

	if options.Json {
		result := SearchResult{}
		result.ChosenFile = chosenFile
		result.Matches = numFiles

		bytes, err := json.Marshal(result)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error encoding JSon:", err)
			return
		}
		os.Stdout.Write(bytes)
	}
}

func buildFilterChain(options Options, history History) chan MatchedFile {
	foundChan := make(chan MatchedFile)

	go walkFiles(".", foundChan)

	filterChan := make(chan MatchedFile)

	go filterFilesByPattern(foundChan, options.Pattern, filterChan)

	historyChan := make(chan MatchedFile)

	go filterByHistory(filterChan, history, historyChan)

	var outChan chan MatchedFile

	if options.MaxDaysOld > 0 || options.MinDaysOld > 0 {
		dateChan := make(chan MatchedFile)
		go filterFilesByDate(historyChan, options.MinDaysOld, options.MaxDaysOld, dateChan)
		outChan = dateChan
	} else {
		outChan = historyChan
	}

	return outChan
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
	maxDaysOld := flag.Int("o", 0, "Oldest modified file to consider (days)")
	minDaysOld := flag.Int("n", 0, "Newest modified file to consider (days)")
	printFound := flag.Bool("l", false, "List files found")
	version := flag.Bool("v", false, "Show version information")
	prune := flag.Bool("p", false, "Prune history of missing files")
	json := flag.Bool("j", false, "Output JSon format to stdout")

	flag.Parse()

	if flag.NArg() < 1 {
		options.Pattern = "*"
	} else {
		pattern := strings.ToLower(flag.Arg(0))

		if !strings.Contains(pattern, "*") {
			pattern = "*" + pattern + "*"
		}
		options.Pattern = pattern
	}

	options.MaxDaysOld = *maxDaysOld
	options.MinDaysOld = *minDaysOld
	options.ListFilesFound = *printFound
	options.ShowVersion = *version
	options.Prune = *prune
	options.Json = *json

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
		name := strings.ToLower(filepath.Base(file.Name))
		if matched, _ := filepath.Match(pattern, name); matched {
			out <- file
		}
	}

	out <- endOfStream
}

func filterFilesByDate(in chan MatchedFile, minDaysOld int, maxDaysOld int, out chan MatchedFile) {
	minFileTime := time.Now().AddDate(0, 0, -minDaysOld)
	maxFileTime := time.Now().AddDate(0, 0, -maxDaysOld)

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

func readHistory(fileName string) (History, error) {
	history := History{}

	// do we have a history? If so, read it
	historyFile, err := os.Open(fileName)
	if os.IsNotExist(err) {
		// we're ok with file not existing.
		err = nil
	} else if err == nil {
		// got history
		dec := json.NewDecoder(historyFile)
		err = dec.Decode(&history)

		// close history file now cos we're gonna write to it later
		historyFile.Close()
	}

	return history, err
}

func pruneHistory(history History) History {
	newHistory := History{}

	for _, fileName := range history.Visited {
		if _, err := os.Stat(fileName); err == nil {
			addFileToHistory(&newHistory, fileName)
		}
	}

	return newHistory
}

func addFileToHistory(historyPtr *History, fileName string) {
	historyPtr.Visited = append(historyPtr.Visited, fileName)
}

func writeHistory(history History, fileName string) {
	var historyFile *os.File
	var err error

	if historyFile, err = os.Create(fileName); err != nil {
		fmt.Println("error (re)creating history file:", err)
		return
	}
	defer historyFile.Close()

	enc, err := json.MarshalIndent(&history, "", "\t")
	if err != nil {
		fmt.Println("error formatting history:", err)
		return
	}

	historyFile.Write(enc)
}

// walkFiles walks all the files in the current directory and subdirectories,
// outputting files of the correct type to the out channel.
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
