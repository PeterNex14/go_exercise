package main

import "sync"

func countWordsConcurrent(texts []string) int {
	if len(texts) == 0 {
		return 0
	}
	
 	channel := make(chan int)
	var wg sync.WaitGroup

	for _, text := range texts {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			channel <- wordCount(t)
		}(text)
	}
	go func() {
		wg.Wait()
		close(channel)
	}()
	

	result := 0
 	for value := range channel {
		result += value
 	}
 
 	return result
}

func wordCount(text string) int {
 if text == "" {
  return 0
 }

 count := 0
 inWord := false

 for _, r := range text {
  if r == ' ' || r == '\t' || r == '\n' {
   if inWord {
    count++
    inWord = false
   }
  } else {
   inWord = true
  }
 }

 if inWord {
  count++
 }

 return count
}
