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

// const (
// 	dbPath     = "./repository/data/codebase%d.db"
// 	bucketName = "codes"
// 	batchSize  = 200_000 // machine dependent
// )

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

// var Bds []*bolt.DB

// func commit(db *bolt.DB, kv [][]byte) error {
// 	return db.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(bucketName))
// 		for _, k := range kv {
// 			if err := b.Put(k, []byte{1}); err != nil { return err }
// 		}
// 		return nil
// 	})
// }

// func initBds() {
// 	for i := 0; i < 3; i++ {
// 		db, err := bolt.Open(fmt.Sprintf(dbPath, i+1), 0o644, &bolt.Options{
// 			Timeout: 1 * time.Minute,
// 			NoFreelistSync:  true,
// 			FreelistType:   bolt.FreelistArrayType,
// 		})
// 		if err != nil { log.Fatal(err) }
// 		defer db.Close()
// 		Bds = append(Bds, db)
// 	}
// }

// func InitData() {
// 	// Check if database files already exist
// 	for i := 0; i < 3; i++ {
// 		dbFile := fmt.Sprintf(dbPath, i+1)
// 		if _, err := os.Stat(dbFile); err == nil {
// 			fmt.Printf("Database file already exists, skipping initialization\n")
// 			return
// 		}
// 	}

// 	initBds()
// 	for _, db := range Bds {
// 		if err := db.Update(func(tx *bolt.Tx) error {
// 			_, e := tx.CreateBucketIfNotExists([]byte(bucketName))
// 			return e
// 		}); err != nil { log.Fatal(err) }
// 	}
// 	start := time.Now()
// 	wg := sync.WaitGroup{}
// 	for i, gz := range files {
// 		wg.Add(1)
// 		go func(db *bolt.DB, gzPath string) {
// 			defer wg.Done()
// 			if err := ingestGZStream(db, gzPath); err != nil {
// 				// Remove the database file on error
// 				if removeErr := os.Remove(fmt.Sprintf(dbPath, i+1)); removeErr != nil {
// 					log.Printf("Failed to remove db file: %v", removeErr)
// 				}
// 				log.Fatalf("ingest %s: %v", gzPath, err)
// 			}
// 		}(Bds[i], gz)
// 	}
// 	wg.Wait()
// 	end := time.Now()
// 	fmt.Println("Time taken: ", end.Sub(start))
// }

// func ingestGZStream(db *bolt.DB, gzPath string) error {
// 	f, err := os.Open(gzPath)
// 	if err != nil { return err }
// 	defer f.Close()

// 	gr, err := gzip.NewReader(f)
// 	if err != nil { return err }
// 	defer gr.Close()

// 	sc := bufio.NewScanner(gr)
// 	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)

// 	tx, err := db.Begin(true)
// 	if err != nil { return err }
// 	b, err := ensureBucket(tx, bucketName)
// 	if err != nil { tx.Rollback(); return err }

// 	lines := 0
// 	commitTx := func() error {
// 			if err := tx.Commit(); err != nil { return err }
// 			tx, err = db.Begin(true)
// 			if err != nil { return err }
// 			b, err = ensureBucket(tx, bucketName)
// 			if err != nil { tx.Rollback(); return err }
// 			lines = 0
// 			return nil
// 	}

// 	for sc.Scan() {
// 			if v := strings.TrimSpace(sc.Text()); v != "" {
// 					k := make([]byte, len(v))
// 					copy(k, v)
// 					if err := b.Put(k, []byte{}); err != nil { tx.Rollback(); return err }
// 					lines++
// 					if lines >= batchSize {
// 							if err := commitTx(); err != nil { return err }
// 					}
// 			}
// 	}
// 	if err := sc.Err(); err != nil { tx.Rollback(); return err }
// 	if lines > 0 { return tx.Commit() }
// 	return tx.Rollback()
// }

// func ensureBucket(tx *bolt.Tx, name string) (*bolt.Bucket, error) {
// 	b := tx.Bucket([]byte(name))
// 	if b != nil { return b, nil }
// 	return tx.CreateBucketIfNotExists([]byte(name))
// }

// func ingestGZ(db *bolt.DB, gzPath string) error {
// 	f, err := os.Open(gzPath)
// 	if err != nil { return err }
// 	defer f.Close()

// 	gr, err := gzip.NewReader(f)
// 	if err != nil { return err }
// 	defer gr.Close()

// 	sc := bufio.NewScanner(gr)
// 	buf := make([]byte, 0, 64*1024)
// 	sc.Buffer(buf, 1024*1024)

// 	var kv [][]byte
// 	commit := func(batch [][]byte) error {
// 		if len(batch) == 0 { return nil }
// 		return db.Update(func(tx *bolt.Tx) error {
// 			b := tx.Bucket([]byte(bucketName))
// 			for _, k := range batch {
// 				if err := b.Put(k, []byte{1}); err != nil {
// 					return err
// 				}
// 			}
// 			return nil
// 		})
// 	}

// 	index := 0
// 	commit_count := 0

// 	kv = make([][]byte, 0, batchSize)

// 	for sc.Scan() {
// 		line := strings.TrimSpace(sc.Text())
// 		// fmt.Println("line", line)
// 		if line == "" { continue }
// 		kv = append(kv, []byte(line))
// 		if index >= batchSize - 1 {
// 			if err := commit(kv); err != nil { return err }
// 			index = 0
// 			fmt.Println(gzPath, " line", line)
// 			fmt.Println(gzPath, " commit_count", commit_count)
// 			kv = kv[:0]
// 			commit_count++
// 			continue
// 		}

// 		index++
// 	}
// 	if err := sc.Err(); err != nil { return err }
// 	if err := commit(kv); err != nil { return err }
// 	return nil
// }

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
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
			if codeMap2[code] && codeMap3[code] {
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
	codes := make([]map[string]bool, 0)
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
