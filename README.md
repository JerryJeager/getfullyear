# GetFullYear Package

The `getFullYear` package provides a simple interface to retrieve the current year and related details from the [GetFullYear API](https://getfullyear.com).

## Features

- Fetch the current year as an integer.
- Retrieve the sponsor of the year data.
- Get a human-readable string representation of the year.

## Installation

To install the package, use the `go get` command:

```bash
go get github.com/JerryJeager/getFullYear
```

## Usage 

- Import the package 

```bash 
import "github.com/JerryJeager/getFullYear"
``` 
- Fetch the Year Data
You can use the GetFullYear function to retrieve year details:

```bash 
package main

import (
	"fmt"
	"log"

	"github.com/JerryJeager/getFullYear"
)

func main() {
	yearData, err := getFullYear.GetFullYear()
	if err != nil {
		log.Fatalf("Error fetching year data: %v", err)
	}

	fmt.Printf("Year: %d\nSponsored By: %s\nYear String: %s\n",
		yearData.Year, yearData.SponsoredBy, yearData.YearString)
}
```

- Output Example 
```bash 
Year: 2025
Sponsored By: "Coca-cola"
Year String: "2025"
```

