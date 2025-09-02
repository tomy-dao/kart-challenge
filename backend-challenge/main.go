package main

import "local/cmd"

func main() {
	// db, err := bolt.Open("./repository/data/couponbase1.db", 0o644, &bolt.Options{Timeout: 1 * time.Minute})
	// if err != nil { log.Fatal(err) }
	// defer db.Close()
	// exists := initdata.ValidateCode("NXAYTZ6M")
	// fmt.Println(exists)
	// initdata.InitData()
	cmd.Run()
}
