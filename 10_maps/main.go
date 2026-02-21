package main

import "fmt"

func main() {
	websites := map[string]string{
		"google":              "https://google.com",
		"amazon web services": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["amazon web services"])

	// Adding a new key-value pair
	websites["larian studios"] = "https://larian.com"
	fmt.Println(websites)

	// Deleting a key-value pair
	delete(websites, "google")
	fmt.Println(websites)
}
