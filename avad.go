package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

func printBanner() {
	fmt.Println("\033[91m" + `
	    (      
		(               (     )\ )   
		)\     (   (    )\   (()/(   
	 ((((_)(   )\  )\((((_)(  /(_))  
	  )\ _ )\ ((_)((_))\ _ )\(_))_   
	  (_)_\(_)\ \ / / (_)_\(_)|   \  
	   / _ \   \ V /   / _ \  | |) | 
	  /_/ \_\   \_/   /_/ \_\ |___/  
	` + "\033[0m")
	fmt.Println("developed by avyaysec")
}

func convertToPostRequest(url string, data string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return client.Do(req)
}

func main() {
	printBanner()

	var (
		url      string
		urlList  string
		dir      string
		dirList  string
		conc     int
		client   = &http.Client{}
		wg       sync.WaitGroup
		urlCh    = make(chan string)
		dirCh    = make(chan string)
		resultCh = make(chan string)
	)

	flag.StringVar(&url, "u", "", "single URL to scan, e.g., http://example.com")
	flag.StringVar(&urlList, "U", "", "path to list of URLs, e.g., urllist.txt")
	flag.StringVar(&dir, "d", "/", "single directory to scan, e.g., /admin")
	flag.StringVar(&dirList, "D", "", "path to list of directories, e.g., dirlist.txt")
	flag.IntVar(&conc, "c", 10, "number of concurrent requests")

	flag.Parse()

	if url == "" && urlList == "" {
		fmt.Println("Please provide a single URL or a list either! (-u or -U)")
		os.Exit(1)
	}

	go func() {
		defer close(urlCh)
		if url != "" {
			urlCh <- url
		} else {
			file, err := os.Open(urlList)
			if err != nil {
				fmt.Printf("Error opening URL list file: %s\n", err)
				os.Exit(1)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				urlCh <- strings.TrimSpace(scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading URL list: %s\n", err)
				os.Exit(1)
			}
		}
	}()

	go func() {
		defer close(dirCh)
		if dirList != "" {
			file, err := os.Open(dirList)
			if err != nil {
				fmt.Printf("Error opening directory list file: %s\n", err)
				os.Exit(1)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				dirCh <- strings.TrimSpace(scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading directory list: %s\n", err)
				os.Exit(1)
			}
		} else {
			dirCh <- dir
		}
	}()

	wg.Add(conc)
	for i := 0; i < conc; i++ {
		go func() {
			defer wg.Done()
			for u := range urlCh {
				for d := range dirCh {
					// Send GET request
					reqURL := u + d
					resp, err := client.Get(reqURL)
					if err != nil {
						fmt.Printf("Error sending GET request to %s: %s\n", reqURL, err)
						continue
					}
					defer resp.Body.Close()

					
					result := fmt.Sprintf("URL: %s, Directory: %s, Status: %s, Size: %d", u, d, resp.Status, resp.ContentLength)
					fmt.Println(result)
					resultCh <- result

					
					postResp, postErr := convertToPostRequest(u, "data=example")
					if postErr != nil {
						fmt.Printf("Error converting GET to POST request: %s\n", postErr)
						continue
					}
					defer postResp.Body.Close()

					postResult := fmt.Sprintf("Converted POST Request: URL: %s, Directory: %s, Status: %s, Size: %d", u, d, postResp.Status, postResp.ContentLength)
					fmt.Println(postResult)
					resultCh <- postResult
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		fmt.Println(res)
	}
}
