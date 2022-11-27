package main

import (
	"flag"
	id3 "github.com/mikkyang/id3-go"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		panic("You most pass the path of the directory in first argument")
	}
	artist := flag.Arg(1)
	if artist == "" {
		panic("You most pass the artist name of mp3 files in the directory in second argument")
	}
	album := flag.Arg(2)
	if artist == "" {
		panic("You most pass the album name of mp3 files in the directory in third argument")
	}

	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !strings.Contains(file.Name(), ".mp3") {
			continue
		}

		p := filepath.Join(path, file.Name())
		f, err := id3.Open(p)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		f.SetArtist(artist)
		f.SetAlbum(album)

		log.Printf("%18s%16s%20s%s", f.Artist(), f.Album(), f.Title(), file.Name())
	}

}
