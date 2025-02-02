# GetFullYear Package

The `getFullYear` package provides a simple interface to retrieve the current year and related details from the [GetFullYear API](https://getfullyear.com).

## Features

- Fetch the current year as an integer.
- Retrieve the sponsor of the year data.
- Get a string representation of the year.

## Installation

To install the package, use the `go get` command:

```bash
go get github.com/JerryJeager/getfullyear
```

## Usage 

- Import the package 

```bash 
import "github.com/JerryJeager/getfullyear"
``` 
- Fetch the Year Data
You can use the Fetch function to retrieve year details:

```bash 
package main

import (
	"fmt"
	"log"

	"github.com/JerryJeager/getfullyear"
)

func main() {
	yearData, err := getfullyear.Fetch()
	if err != nil {
		log.Fatalf("Error fetching year data: %v", err)
	}

	fmt.Printf("Year: %d\nSponsored By: %s\nYear String: %s\n",
		yearData.Year, yearData.SponsoredBy, yearData.YearString)
}
```

- Output Example 
```bash 
year: 2025
sponsored_by: "Coca-cola"
year_string: "2025"
```
