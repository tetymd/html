package main

import (
    "log"
    "os"
)

var template = "<html lang=\"ja\">\n</html>"

func writingTemplate(f *os.File, template string) (err error) {
    _, err = f.WriteString(template)
    return err
}

func main() {
    f, err := os.Create("index.html")
    if err != nil {
        log.Println(err)
        return
    }

    err = writingTemplate(f, template)
    if err != nil {
        log.Println(err)
    }
}


