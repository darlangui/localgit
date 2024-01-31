package main

import (
	"flag"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"io"
	"log"
	"os"
	"time"
)

type Commmit struct {
	Key    time.Time
	Amount int
}

func verifierDir(folder string) bool {
	f, err := os.Open(folder)

	if err != nil {
		panic(err)
	}

	for {
		files, err := f.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}

		if files[0].Name() == ".git" {
			return true
		}
	}

	return false
}

func getCommits(folder string) []Commmit {
	repo, err := git.PlainOpen(folder)

	if err != nil {
		log.Fatal(err)
	}

	ref, err := repo.Head()

	if err != nil {
		log.Fatal(err)
	}

	cIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})

	if err != nil {
		log.Fatal(err)
	}

	SliceCommit := []Commmit{}
	var temp, tempFormat time.Time
	err = cIter.ForEach(func(commit *object.Commit) error {
		if len(SliceCommit) == 0 {
			temp = commit.Author.When
			SliceCommit = append(SliceCommit, Commmit{commit.Author.When, 1})
		} else {
			if tempFormat.Format("01/02/2006") != commit.Author.When.Format("01/02/2006") {
				SliceCommit = append(SliceCommit, Commmit{commit.Author.When, 1})
				temp = commit.Author.When
			} else {
				for i := range SliceCommit {
					if SliceCommit[i].Key == temp {
						SliceCommit[i].Amount += 1
						break
					}
				}
			}
		}
		tempFormat = commit.Author.When
		return nil
	})

	return SliceCommit
}

func main() {
	var folder string
	flag.StringVar(&folder, "add", "", "novo diretório para análise")
	flag.Parse()
	if verifierDir(folder) {
		SliceCommit := getCommits(folder)

		for i := range SliceCommit {
			i++
			//fmt.Printf("%q ----- %d\n", SliceCommit[i].Key, SliceCommit[i].Amount)
		}

	}
}
