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
)

type typeXMLInfo struct {
	RelativeURL string `xml:"entry>relative-url"`
}

type typeXMLLog struct {
	Entries []struct {
		Revision string   `xml:"revision,attr"`
		Author   string   `xml:"author"`
		Date     string   `xml:"date"`
		Message  string   `xml:"msg"`
		Paths    []string `xml:"paths>path"`
	} `xml:"logentry"`
}

type typeCommitPath struct {
	Revision string
	Author   string
	Date     string
	Message  string
	Path     string
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
	for _, entry := range logs.Entries {
		for _, path := range entry.Paths {
			if !strings.HasPrefix(path, relativeURL) {
				continue
			} else if path[len(relativeURL):] == "" {
				continue
			}
			path = strings.TrimLeft(path[len(relativeURL):], "/")
			if !inSSlice(res, path) {
				res = append(res, typeCommitPath{
					Revision: entry.Revision,
					Author:   entry.Author,
					Date:     strings.Replace(entry.Date[:19], "T", " ", 1),
					Message:  entry.Message,
					Path:     path,
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

func main() {
	var (
		args []string
		path string
	)
	flag.Parse()
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
	logContent, err := exec.Command("svn", "log", "-v", "-l", "10", "--xml", "-l", "100", path).Output()
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, entry := range getXMLLog(relativeURL, logContent) {
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s\t%s", entry.Date, entry.Revision, entry.Author, entry.Path))
	}
	w.Flush()

}
