package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWordToken = "<*>"

var transformations = []string{
	otherWordToken,
	otherWordToken + "app",
	otherWordToken + "site",
	otherWordToken + "world",
	otherWordToken + "time",
	"go" + otherWordToken,
	"get" + otherWordToken,
	"lets" + otherWordToken,
}

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		transformation := transformations[rand.Intn(len(transformations))]
		fmt.Println(strings.Replace(transformation, otherWordToken, scanner.Text(), -1))
	}
}
