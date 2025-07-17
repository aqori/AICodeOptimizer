// cmd/aicodeoptimizer/main.go
package main

import (
"flag"
"log"
"os"

"aicodeoptimizer/internal/aicodeoptimizer"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := aicodeoptimizer.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
