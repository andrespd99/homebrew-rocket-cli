package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

type fileEntry struct {
	Path                 string `json:"path"`
	Base64EncodedContent string `json:"content"`
}

func main() {
	var (
		src  string
		out  string
		args []string
	)

	args = os.Args[1:]

	if len(args) == 0 {
		showHelp()
		return
	}

	src = os.Args[1]

	if len(args) == 1 {
		out = ".output/"
	} else {
		out = os.Args[2]
	}

	fileName := flag.String("fileName", "entries.json", "name of the output file")
	flag.Parse()

	entries, err := serializeDir(src)
	if err != nil {
		log.Fatalln(err)
	}

	entries = cleanPaths(src, entries)

	data, err := encodeJson(entries)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = os.ReadDir(out)
	if err != nil && os.IsNotExist(err) {
		os.Mkdir(out, 0755)
	}

	file, err := os.Create(path.Join(out, *fileName))
	if err != nil {
		log.Fatalln(err)
	}

	_, err = file.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
}

func serializeDir(src string) ([]fileEntry, error) {
	dir, err := os.ReadDir(src)
	if err != nil {
		return nil, err
	}

	var entries []fileEntry

	for _, d := range dir {
		if !d.IsDir() {
			fPath := path.Join(src, d.Name())
			fBytes, err := os.ReadFile(fPath)
			if err != nil {
				return nil, err
			}

			b64 := encodeBase64(fBytes)
			entries = append(entries, fileEntry{Path: fPath, Base64EncodedContent: b64})

			continue
		}

		subEntries, err := serializeDir(path.Join(src, d.Name()))
		if err != nil {
			return nil, err
		}

		entries = append(entries, subEntries...)
	}

	return entries, nil
}

func encodeBase64(in []byte) string {
	return base64.StdEncoding.EncodeToString(in)
}

func encodeJson(data any) ([]byte, error) {
	return json.Marshal(data)
}

func cleanPaths(src string, entries []fileEntry) []fileEntry {
	var clean []fileEntry
	for _, e := range entries {
		e.Path = strings.TrimPrefix(path.Clean(e.Path), path.Clean(src))
		clean = append(clean, e)
	}
	return clean
}

func showHelp() {
	fmt.Println("Usage: encoder <src> <out-dir> [arguments]")
	fmt.Println("<src>			source directory to be encoded")
	fmt.Println("<out-dir>		output directory (defaults to .output/)")
	fmt.Println("--fileName		name of the output file")
}
