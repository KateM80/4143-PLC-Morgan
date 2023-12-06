package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string, outputDir string) error {
	for _, url := range urls {
		// Get the filename from the URL
		filename := filepath.Base(url)
		// Append ".jpg" extension to the filename
		filename = filename + ".jpg"

		// Create the file on disk
		filePath := filepath.Join(outputDir, filename)
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %v", filePath, err)
		}
		defer file.Close()

		// Download the image
		response, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("failed to download image from %s: %v", url, err)
		}
		defer response.Body.Close()

		// Check if the request was successful
		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to download image from %s: status code %d", url, response.StatusCode)
		}

		_, err = io.Copy(file, response.Body)
		if err != nil {
			return fmt.Errorf("failed to save image to file %s: %v", filePath, err)
		}

		fmt.Printf("Downloaded %s to %s\n", url, filePath)
	}

	return nil
}

// Concurrent version of the image downloader.
func downloadImage(url string, outputDir string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Get the filename from the URL
	filename := filepath.Base(url)
	// Append ".jpg" extension to the filename
	filename = filename + ".jpg"

	// Create the file
	filePath := filepath.Join(outputDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	// Download the image
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading image from %s: %v\n", url, err)
		return
	}
	defer response.Body.Close()

	// Check if the request was successful
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Failed to download image from %s: status code %d\n", url, response.StatusCode)
		return
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("Error saving image to file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("Downloaded %s to %s\n", url, filePath)
}
//Actual Concurrent version of the image downloader.
//This really just syncs up the images 
func downloadImagesConcurrent(urls []string, outputDir string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go downloadImage(url, outputDir, &wg)
	}

	wg.Wait()
}

func main() {
	urls := []string{
		"https://unsplash.com/photos/a-black-and-white-photo-of-a-lot-of-numbers-yxJL8gmcSAY",
		"https://unsplash.com/photos/a-man-in-a-hoodie-smoking-a-cigarette-etcN-HM_JRU",
		"https://unsplash.com/photos/a-close-up-of-a-cat-looking-at-the-camera-rj4KLIcDwwc",
		"https://unsplash.com/photos/a-green-and-red-bird-sitting-on-top-of-a-white-fence-vB6HQIerLMI",
	}

	outputDir := "./downloaded_images"
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}
	// Sequential download
	start := time.Now()

	err = downloadImagesSequential(urls, outputDir)
	if err != nil {
		fmt.Printf("Error downloading images: %v\n", err)
	}

	fmt.Printf("Sequential download took: %v\n", time.Since(start))
 
  // Concurrent download
  
	outputDir2 := "./downloaded_images_concurrent"

	err2 := os.MkdirAll(outputDir2, os.ModePerm)
	if err2 != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}
	
	start = time.Now()
	downloadImagesConcurrent(urls, outputDir2)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}
