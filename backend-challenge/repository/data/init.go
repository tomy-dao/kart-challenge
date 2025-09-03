package initdata

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var files = []string{}

func init() {
	// Scan the data directory to find .gz files
	dataDir := "./repository/data"
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		fmt.Printf("Error reading data directory: %v\n", err)
		return
	}
	
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".gz") {
			files = append(files, fmt.Sprintf("%s/%s", dataDir, entry.Name()))
		}
	}
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("Alloc = ", bToMb(m.Alloc), " MiB")
	fmt.Println("TotalAlloc = ", bToMb(m.TotalAlloc), " MiB")
	fmt.Println("Sys = ", bToMb(m.Sys), " MiB")
	fmt.Println("NumGC = ", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func loadData(file string) map[string]bool {
	f, err := os.Open(file)
	if err != nil { return nil }
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil { return nil }
	defer gr.Close()

	codeMap := make(map[string]bool)
	
	sc := bufio.NewScanner(gr)
	buf := make([]byte, 0, 64*1024)
	sc.Buffer(buf, 1024*1024)
	i := 0
	batchSize := 10_000_000
	for sc.Scan() {
		if v := strings.TrimSpace(sc.Text()); v != "" {
			codeMap[v] = true
		}
		// fmt.Print(file, " length ", i, "\n")
		i++
		if i % batchSize == 0 {
			fmt.Println("Loaded ", i, " codes from ", file)
		}
	}
	if err := sc.Err(); err != nil { return nil }
	fmt.Println("Loaded data from ", file)
	return codeMap
}

func handleAllowedCodes(codes []map[string]bool) map[string]bool {
	codeMap1 := codes[0]
	codeMap2 := codes[1]
	codeMap3 := codes[2]

	allowedCodesLock := sync.Mutex{}
	allowedCodes := make(map[string]bool)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		for code, _ := range codeMap1 {
			if codeMap2[code] || codeMap3[code] {
				allowedCodesLock.Lock()
				allowedCodes[code] = true
				allowedCodesLock.Unlock()
			}
		}
		timeEnd := time.Now()
		fmt.Println("handleAllowedCodes 1 time taken: ", timeEnd.Sub(start))
	}()

	
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		for code, _ := range codeMap2 {
			if codeMap1[code] && codeMap3[code] {
				allowedCodesLock.Lock()
				allowedCodes[code] = true
				allowedCodesLock.Unlock()
			}
		}
		timeEnd := time.Now()
		fmt.Println("handleAllowedCodes 2 time taken: ", timeEnd.Sub(start))
	}()

	wg.Wait()
	return allowedCodes
}

var allowedCodes = make(map[string]bool)

func InitData() {
	// Check if allowed_codes.txt already exists and load it
	if _, err := os.Stat("./repository/data/allowed_codes.txt"); err == nil {
		// File exists, load from it
		file, err := os.Open("./repository/data/allowed_codes.txt")
		if err != nil {
			fmt.Printf("Error opening allowed_codes.txt: %v\n", err)
		} else {
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if code := strings.TrimSpace(scanner.Text()); code != "" {
					allowedCodes[code] = true
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading allowed_codes.txt: %v\n", err)
			} else {
				fmt.Printf("Loaded %d allowed codes from file\n", len(allowedCodes))
				return
			}
		}
	}
	timeStart := time.Now()
	wg := sync.WaitGroup{}
	codesLock := sync.Mutex{}
	codes := make([]map[string]bool, 0, len(files))
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			codeMap := loadData(file)
			codesLock.Lock()
			defer codesLock.Unlock()
			codes = append(codes, codeMap)
		}(file)
	}
	wg.Wait()
	fmt.Println("memory usage after loading data")
	printMemUsage()

	timeEnd := time.Now()
	fmt.Println("Load data time: ", timeEnd.Sub(timeStart))
	timeStart = time.Now()
	allowedCodes = handleAllowedCodes(codes)
	timeEnd = time.Now()
	fmt.Println("Time taken: ", timeEnd.Sub(timeStart))
	fmt.Println("Allowed codes: ", len(allowedCodes))
	fmt.Println("memory usage after handling allowed codes")
	printMemUsage()
	// Save allowedCodes to a file
	saveAllowedCodes(allowedCodes)
}

func saveAllowedCodes(allowedCodes map[string]bool) {
	saveFile, err := os.Create("./repository/data/allowed_codes.txt")
	if err != nil {
		fmt.Printf("Error creating allowed codes file: %v\n", err)
	} else {
		defer saveFile.Close()
		writer := bufio.NewWriter(saveFile)
		for code := range allowedCodes {
			writer.WriteString(code + "\n")
		}
		writer.Flush()
	}
}

func CheckCodeExistsInAllowedCodes(code string) bool {
	if code == "" {
		return false
	}
	if _, ok := allowedCodes[code]; ok {
		return true
	}
	return false
}

func CheckCodeExists(file string, code string) (bool, error) {
	f, err := os.Open(file)
	if err != nil { return false, err }
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil { return false, err }
	defer gr.Close()

	sc := bufio.NewScanner(gr)
	buf := make([]byte, 0, 64*1024)
	sc.Buffer(buf, 1024*1024)
	for sc.Scan() {
		if v := strings.TrimSpace(sc.Text()); v != "" {
			if v == code {
				return true, nil
			}
		}
	}
	if err := sc.Err(); err != nil { return false, err }
	return false, nil
}

func ValidateCode(code string) bool {
	mu := sync.Mutex{}
	count := 0
	wg := sync.WaitGroup{}
	timeStart := time.Now()
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			exists, err := CheckCodeExists(file, code)
			if err != nil {
				return
			}
			if exists {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}(file)
	}
	wg.Wait()
	timeEnd := time.Now()
	fmt.Println("Time taken: ", timeEnd.Sub(timeStart))
	return count > 2
}
