package main

import (
	"flag"
	"fmt"
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

func tableCommits(SliceCommit []Commmit) {
	now := time.Now()
	now = now.AddDate(0, 0, 1)
	var ant time.Time
	aux := 0

	switch now.Format("Monday") {
	case "Sunday":
		ant = now.AddDate(0, 0, -8*7)
		aux = 0
	case "Monday":
		ant = now.AddDate(0, 0, -1+(-8*7))
		aux = 1
	case "Tuesday":
		ant = now.AddDate(0, 0, -2+(-8*7))
		aux = 2
	case "Wednesday":
		ant = now.AddDate(0, 0, -3+(-8*7))
		aux = 3
	case "Thursday":
		ant = now.AddDate(0, 0, -4+(-8*7))
		aux = 4
	case "Friday":
		ant = now.AddDate(0, 0, -5+(-8*7))
		aux = 5
	case "Saturday":
		ant = now.AddDate(0, 0, -6+(-8*7))
		aux = 6
	}
	dateCommit := []string{}

	flag := false
	fmt.Println(now)
	for ant.Before(now) {
		for _, commit := range SliceCommit {
			if commit.Key.Format("02/01/2006") == ant.Format("02/01/2006") {
				dateCommit = append(dateCommit, "( "+ant.Format("02-01")+" - "+"["+strconv.Itoa(commit.Amount)+"]"+" )")
				flag = true
			}
		}
		if flag == false {
			dateCommit = append(dateCommit, "( "+ant.Format("02")+" )")
		}
		flag = false
		ant = ant.AddDate(0, 0, 1)
	}
	createTable(dateCommit, aux)
}

func createTable(dateCommit []string, aux int) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"  Sunday  ", "  Monday  ", "  Tuesday  ", "  Wednesday  ", "  Thursday  ", "  Friday  ", "  Saturday  "})

	var tempRow []interface{}
	//fmt.Println(dateCommit)
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
	flag.StringVar(&folder, "add", "", "novo diretório para análise")
	flag.Parse()
	if verifierDir(folder) {
		SliceCommit := getCommits(folder)
		tableCommits(SliceCommit)
	}
}
