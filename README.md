# go-html2json
Golang html convert to json.  Golang implementation html2json lib.

#EXAMPLE

main.go

```Go
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/v-grabko1999/go-html2json"
)

func main() {
	d, err := html2json.New(strings.NewReader(`
		<html>
			<head>
				<title>Hello World</title>
			</head>
			<body>
			Hello World!
			<p>P</p>
			</body>
		</html>
	`))
	if err != nil {
		log.Fatal(err)
	}
	json, err := d.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}

```

```go run main.go```

Output:
```json
[
    {
        "elements": [
            {
                "name": "html",
                "elements": [
                    {
                        "name": "head",
                        "elements": [
                            {
                                "name": "title",
                                "text": "Hello World"
                            }
                        ]
                    },
                    {
                        "name": "body",
                        "text": "Hello World!",
                        "elements": [
                            {
                                "name": "p",
                                "text": "P"
                            }
                        ]
                    }
                ]
            }
        ]
    }
]
```