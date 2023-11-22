package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	for i, url := range urls {
		downloadImage(url, fmt.Sprintf("s_img%d.jpg", i+1))
	}
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	// TODO: Implement concurrent download logic
}

func main() {
	urls := []string{
		"https://images.unsplash.com/photo-1682686580186-b55d2a91053c?q=80&w=1975&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1699694927495-4b8da5df6ce5?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1699704674521-54b3252e52e9?q=80&w=2072&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.pexels.com/photos/18820159/pexels-photo-18820159/free-photo-of-winter-mountain.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
		"https://cdn.stocksnap.io/img-thumbs/960w/deer-animal_NERJJRVKBO.jpg",
	}

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	//start = time.Now()
	//downloadImagesConcurrent(urls)
	//fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	// Create an HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return err
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", response.StatusCode)
		return err
	}

	// Create a new file to save the image
	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return err
	}
	defer outputFile.Close()

	// Copy the HTTP response body to the file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Println("Error saving the image:", err)
		return err
	}

	return nil
}
