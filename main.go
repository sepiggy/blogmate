package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"log"

	"github.com/AlecAivazis/survey/v2"
)

const (
	CONTENT_PATH = "/home/sepiggy/Languages/go-study/blogmate/content"
)

var (
	cs []string
	qs []*survey.Question
)

func init() {
	fileInfoList, err := ioutil.ReadDir(CONTENT_PATH)
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range fileInfoList {
		cs = append(cs, f.Name())
	}
	cs = append(cs, "create a new category")

	qs = []*survey.Question{
		{
			Name:      "url",
			Prompt:    &survey.Input{Message: "What is the blog url?"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "category",
			Prompt: &survey.Select{
				Message: "Choose a category:",
				Options: []string(cs),
			},
		},
	}
}

func main() {
	// ucas aka "Url and Category answers"
	ucas := struct {
		Url      string
		Category string
	}{}

	// nca aka "New Category answer"
	nca := struct {
		Name string
	}{}

	// c aka "Category"
	c := ""

	err := survey.Ask(qs, &ucas)
	if err != nil {
		log.Fatalln(err)
	}

	if ucas.Category == "create a new category" {
		qs := []*survey.Question{
			{
				Name:      "name",
				Prompt:    &survey.Input{Message: "What is the new category?"},
				Validate:  survey.Required,
				Transform: survey.Title,
			},
		}

		err := survey.Ask(qs, &nca)
		if err != nil {
			log.Fatalln(err)
		}
		c = nca.Name
	} else {
		c = ucas.Category
	}

	demoDir := CONTENT_PATH + "/" + c + "/" + strconv.FormatInt(time.Now().Unix(), 10)

	os.MkdirAll(demoDir, os.ModePerm)

	// Update README.md
	s := fmt.Sprintf("Title: [Blog](%s), [Demo](%s)\n", ucas.Url, demoDir)
	f, err := os.OpenFile("README.md", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	f.Write([]byte(s))
}
