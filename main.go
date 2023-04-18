package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Make GET request to https://www.quantstart.com/articles/
	resp, err := http.Get("http://www.quantstart.com/articles/")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Find all article links in the response body
	articleLinks := regexp.MustCompile(`<a href="/articles/[\w-]+/"`).FindAllSubmatch(body, -1)

	// Create a directory to store HTML files if it doesn't already exist
	if _, err := os.Stat("html_files"); os.IsNotExist(err) {
		os.Mkdir("html_files", os.ModePerm)
	}

	// Loop through article links and save each HTML file
	for _, link := range articleLinks {
		var link = string(link[0])
		link = strings.Trim(link, "<a href=\"")
		fmt.Println(link)
		// Make GET request to article link
		url := "http://www.quantstart.com" + link

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making GET request:", err)
			continue
		}
		defer resp.Body.Close()

		// Read response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		// Create file name for HTML file
		fileName := "html_files/" + strings.Trim(link, "/articles/") + ".html"

		// Write response body to HTML file
		err = ioutil.WriteFile(fileName, body, os.ModePerm)
		if err != nil {
			fmt.Println("Error writing HTML file:", err)
			continue
		}

		fmt.Println("Saved HTML file:", fileName)
	}
}
