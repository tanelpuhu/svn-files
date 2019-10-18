package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

var (
	flags struct {
		limit    int
		revision string
	}
)

type typeXMLInfo struct {
	RelativeURL string `xml:"entry>relative-url"`
}

type typePath struct {
	Path   string `xml:",chardata"`
	Action string `xml:"action,attr"`
}

type typeXMLLog struct {
	Entries []struct {
		Revision string     `xml:"revision,attr"`
		Author   string     `xml:"author"`
		Date     string     `xml:"date"`
		Message  string     `xml:"msg"`
		Paths    []typePath `xml:"paths>path"`
	} `xml:"logentry"`
}

type typeCommitPath struct {
	Revision string
	Author   string
	Date     string
	Message  string
	Path     string
	Action   string
}

func getRelativeURL(content []byte) string {
	res := typeXMLInfo{}
	err := xml.Unmarshal([]byte(content), &res)
	if err != nil {
		return ""
	}
	return strings.TrimLeft(res.RelativeURL, "^")
}

func inSSlice(hay []typeCommitPath, needle string) bool {
	for _, item := range hay {
		if item.Path == needle {
			return true
		}
	}
	return false
}

func getXMLLog(relativeURL string, content []byte) []typeCommitPath {
	logs := typeXMLLog{}
	err := xml.Unmarshal([]byte(content), &logs)
	if err != nil {
		panic(err)
	}
	var (
		res []typeCommitPath
	)
	entries := logs.Entries
	sort.Slice(entries, func(i, j int) bool {
		i, _ = strconv.Atoi(entries[i].Revision)
		j, _ = strconv.Atoi(entries[j].Revision)
		return i > j
	})

	for _, entry := range entries {
		for _, path := range entry.Paths {
			if !strings.HasPrefix(path.Path, relativeURL) {
				continue
			} else if path.Path[len(relativeURL):] == "" {
				continue
			}
			path.Path = strings.TrimLeft(path.Path[len(relativeURL):], "/")
			if !inSSlice(res, path.Path) {
				res = append(res, typeCommitPath{
					Revision: entry.Revision,
					Author:   entry.Author,
					Date:     textToLocalTimeText(entry.Date),
					Message:  entry.Message,
					Path:     path.Path,
					Action:   path.Action,
				})
			}

		}
	}
	sort.Slice(res, func(i, j int) bool {
		i, _ = strconv.Atoi(res[i].Revision)
		j, _ = strconv.Atoi(res[j].Revision)
		return i < j
	})
	return res
}

func textToLocalTimeText(text string) string {
	result, err := time.Parse("2006-01-02T15:04:05.000000Z", text)
	if err != nil {
		log.Fatalf("error parsing date %s: %v", text, err)
	}
	return result.Local().Format("2006-01-02 15:04:05")
}

func main() {
	flag.IntVar(&flags.limit, "l", 100, "How many last commits to check")
	flag.StringVar(&flags.revision, "r", "", "Revision(s) (or range with NUMBER/DATE/HEAD/etc))")
	flag.Parse()

	var (
		args []string
		path string
	)
	args = flag.Args()

	if len(args) == 0 {
		path, _ = os.Getwd()
	} else {
		path = args[0]
	}

	infoContent, err := exec.Command("svn", "info", "--xml", path).CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", infoContent)
		log.Fatal(err)
	}
	relativeURL := getRelativeURL(infoContent)

	cmdargs := []string{"log", "-v", "--xml"}
	if flags.revision != "" {
		cmdargs = append(cmdargs, []string{"-r", flags.revision}...)
	} else {
		cmdargs = append(cmdargs, []string{"-l", fmt.Sprintf("%d", flags.limit)}...)
	}
	cmdargs = append(cmdargs, path)

	logContent, err := exec.Command("svn", cmdargs...).Output()
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
	for _, entry := range getXMLLog(relativeURL, logContent) {
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s\t%s\t%s", entry.Date, entry.Revision, entry.Author, entry.Action, entry.Path))
	}
	w.Flush()
}
