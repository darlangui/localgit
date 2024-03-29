package main

import (
	"flag"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"io"
	"log"
	"os"
	"strconv"
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

func tableCommits(SliceCommit []Commmit, weeks string) {
	now := time.Now()
	var ant time.Time

	numWeeks, err := strconv.Atoi(weeks)

	if err != nil {
		panic(err)
	}

	switch now.Format("Monday") {
	case "Sunday":
		ant = now.AddDate(0, 0, -8*7)
	case "Monday":
		ant = now.AddDate(0, 0, -1+(-numWeeks*7))
	case "Tuesday":
		ant = now.AddDate(0, 0, -2+(-numWeeks*7))
	case "Wednesday":
		ant = now.AddDate(0, 0, -3+(-numWeeks*7))
	case "Thursday":
		ant = now.AddDate(0, 0, -4+(-numWeeks*7))
	case "Friday":
		ant = now.AddDate(0, 0, -5+(-numWeeks*7))
	case "Saturday":
		ant = now.AddDate(0, 0, -6+(-numWeeks*7))
	}

	dateCommit := []string{}

	aux := false

	now = now.AddDate(0, 0, 1)
	for ant.Before(now) {
		for _, commit := range SliceCommit {
			if commit.Key.Format("02/01/2006") == ant.Format("02/01/2006") {
				dateCommit = append(dateCommit, "( "+ant.Format("02-01")+" - "+"["+strconv.Itoa(commit.Amount)+"]"+" )")
				aux = true
			}
		}
		if aux == false {
			dateCommit = append(dateCommit, "( "+ant.Format("02")+" )")
		}
		aux = false
		ant = ant.AddDate(0, 0, 1)
	}
	createTable(dateCommit)
}

func createTable(dateCommit []string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"  Sunday  ", "  Monday  ", "  Tuesday  ", "  Wednesday  ", "  Thursday  ", "  Friday  ", "  Saturday  "})

	var tempRow []interface{}
	for i, commit := range dateCommit {
		tempRow = append(tempRow, commit)

		if (i+1)%7 == 0 || i+1 == len(dateCommit) {
			t.AppendRow(tempRow)
			tempRow = []interface{}{}
		}
	}

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 2, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 3, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 4, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
	})
	t.SetStyle(table.StyleColoredRedWhiteOnBlack)
	t.Render()
}

func main() {
	var folder string
	var weeks string
	flag.StringVar(&folder, "add", "", "novo diretório para análise")
	flag.StringVar(&weeks, "weeks", "8", "define quantas semanas serão analisadas")
	flag.Parse()
	if verifierDir(folder) {
		SliceCommit := getCommits(folder)
		tableCommits(SliceCommit, weeks)
	}
}
