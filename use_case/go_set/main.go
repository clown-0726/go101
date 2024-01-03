package main

import mapset "github.com/deckarep/golang-set/v2"

func main() {
	// Use package: github.com/deckarep/golang-set/v2
	required := mapset.NewSet[string]()
	required.Add("cooking")
	required.Add("english")
	required.Add("math")
	required.Add("biology")
}
