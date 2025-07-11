// cmd/tokenizedartifactsrepository/main.go
package main

import (
"flag"
"log"
"os"

"tokenizedartifactsrepository/internal/tokenizedartifactsrepository"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := tokenizedartifactsrepository.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
