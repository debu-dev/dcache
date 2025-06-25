# dcache
<<<<<<< HEAD
Simple cache based on golang
=======

dcache is a simple in-memory caching library written in Go. It provides a straightforward way to store and retrieve data using key-value pairs, making it easy to enhance the performance of applications by reducing the need for repeated data fetching.

## Features

- Simple API for setting and getting cached values.
- Supports any data type for values.
- Utility functions for key validation and error logging.

## Installation

To install the dcache library, use the following command:

```
go get github.com/yourusername/dcache
```

## Usage

Here is a simple example of how to use dcache in your application:

```go
package main

import (
    "fmt"
    "github.com/yourusername/dcache/internal/cache"
)

func main() {
    c := cache.NewCache()

    // Set a value in the cache
    c.Set("foo", "bar")

    // Get the value from the cache
    value, found := c.Get("foo")
    if found {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key not found")
    }
}
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
>>>>>>> 8dd7ded (init commit)
