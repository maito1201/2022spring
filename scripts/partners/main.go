package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v2"
)

type Partner struct {
	Category    string `csv:"category"`
	Key         string `csv:"key"`
	Title       string `csv:"title"`
	Order       int    `csv:"order"`
	Description string `csv:"description"`
	Logo        string
}

func createPartner(ps []*Partner, category string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirPath := filepath.Join(wd, "content/partners", category)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	for _, p := range ps {
		p.Logo = fmt.Sprintf("/images/partners/%s.png", p.Key)
		f, err := os.OpenFile(filepath.Join(dirPath, fmt.Sprintf("%s.md", p.Key)), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		out, err := yaml.Marshal(p)
		if err != nil {
			panic(err)
		}
		body := fmt.Sprintf(`---
%s---
%s`, string(out), p.Description)
		rd := strings.NewReader(body)

		_, err = io.Copy(f, rd)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	f, err := os.Open("./raw_data/partners.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var partners, platinums, golds, silvers, bronzes []*Partner
	if err := gocsv.UnmarshalFile(f, &partners); err != nil { // Load clients from file
		panic(err)
	}
	for _, p := range partners {
		switch p.Category {
		case "platinum":
			platinums = append(platinums, p)
		case "gold":
			golds = append(golds, p)
		case "silver":
			silvers = append(silvers, p)
		case "bronze":
			bronzes = append(bronzes, p)
		}
	}
	createPartner(platinums, "platinum")
	createPartner(golds, "gold")
	createPartner(silvers, "silver")
	createPartner(bronzes, "bronze")
}
