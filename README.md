<header>
<div align="center">

<a href="https://github.com/Diaszano">
    <img src=".github/assets/logo.svg" alt="logo" height="150">
</a>

<h1>strrand</h1>

<p>A simple and efficient library written in Go (Golang) for generating customizable random strings.</p>

</div>
</header>

# Introduction

The **strrand** library is designed to simplify the generation of random strings quickly, efficiently, and hassle-free. Moreover, it has no external dependencies, ensuring lightweight and easy integration into your Go projects.

## Installation

To install the library, simply run the command below:

```shell
go get github.com/Diaszano/strrand
```

## Usage

Using **strrand** is very intuitive. Check out some practical examples below to get started:

### Generate Binary Strings
```go
package main

import (
	"fmt"

	"github.com/Diaszano/strrand"
)

func main() {
	for range 3 {
		str := strrand.Binary(10)
		fmt.Println(str)
	}
}

// Example output:
// 0111110101
// 0101111110
// 0110000001
```

### Generate Hexadecimal Strings
```go
package main

import (
	"fmt"

	"github.com/Diaszano/strrand"
)

func main() {
	for range 3 {
		str := strrand.Hexadecimal(10)
		fmt.Println(str)
	}
}

// Example output:
// 8a72bc19a7
// 652bac2a41
// 75ff495c77
```

### Generate Base62 Strings
```go
package main

import (
	"fmt"

	"github.com/Diaszano/strrand"
)

func main() {
	for range 3 {
		str := strrand.Base62(10)
		fmt.Println(str)
	}
}

// Example output:
// Nwc5Q0ARc4
// 6D1RLHU4eL
// 8djCoMKfh8
```

### Generate Strings with Random Characters (Default)
```go
package main

import (
	"fmt"

	"github.com/Diaszano/strrand"
)

func main() {
	for range 3 {
		str := strrand.String(10)
		fmt.Println(str)
	}
}

// Example output:
// %`'s72,6;3
// gt)Dn5rNZi
// @AH!=qk^R8
```

### Generate Custom Strings
You can specify a custom character set for string generation:

```go
package main

import (
	"fmt"

	"github.com/Diaszano/strrand"
)

func main() {
	charset := "abcd12345"
	for range 3 {
		str := strrand.String(10, charset)
		fmt.Println(str)
	}
}

// Example output:
// 4b42233cc3
// ba4c41c2c2
// 345d4acd51
```

## Contribution

Contributions are welcome! If you find any issues, have suggestions, or want to collaborate with improvements, feel free to open an [issue](https://github.com/Diaszano/strrand/issues) or submit a pull request.

---

**License:** This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
