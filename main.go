package main

import (
    "log"
    "os"
)

var template = `<html lang="ja">
    <head>
        <meta charset="utf-8">
    </head>
    <body>
        <h1>Hello!</h1>
    </body>
</html>
`

func writingTemplate(f *os.File, template string) (n int, err error) {
    n, err = f.WriteString(template)
    return n, err
}

func main() {
    f, err := os.Create("index.html")
    if err != nil {
        log.Println(err)
        return
    }

    _, err = writingTemplate(f, template)
    if err != nil {
        log.Println(err)
    }
}


