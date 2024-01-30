package main

import (
	"flag"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"io"
	"log"
	"os"
	"time"
)

type Commmit struct {
	Key   time.Time
	Count int
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

func MapCommits(folder string) []Commmit {
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

	SliceWeek := []Commmit{}
	var aux time.Time
	var auxTwo time.Time
	err = cIter.ForEach(func(commit *object.Commit) error {
		if len(SliceWeek) == 0 {
			aux = commit.Author.When
			SliceWeek = append(SliceWeek, Commmit{commit.Author.When, 1})
		} else {
			if auxTwo.Format("01/02/2006") != commit.Author.When.Format("01/02/2006") {
				SliceWeek = append(SliceWeek, Commmit{commit.Author.When, 1})
				aux = commit.Author.When
			} else {
				for i := range SliceWeek {
					if SliceWeek[i].Key == aux {
						SliceWeek[i].Count += 1
						break
					}
				}
			}
		}
		auxTwo = commit.Author.When
		return nil
	})

	return SliceWeek
}

func main() {
	var folder string
	flag.StringVar(&folder, "add", "", "novo diretório para análise")
	flag.Parse()
	if verifierDir(folder) {
		MapCommits(folder)
		SliceCommit := MapCommits(folder)

		for i := range SliceCommit {
			fmt.Printf("%q ----- %d\n", SliceCommit[i].Key, SliceCommit[i].Count)
		}

	}
}
