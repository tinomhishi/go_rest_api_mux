package main

import "os"

func main() {
	a: = App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME")
		os.Getenv("APP_DB_NAME")
		os.Getenv("APP_DB_PASSWORD")
	)

	a.Run(":8080")
}