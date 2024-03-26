# README

## Overview

This Go package provides functionality for generating sorted Base64-encoded IDs and performance testing the ID generation process. The package includes methods for encoding and decoding strings using a sorted Base64 encoding scheme, as well as a method for generating unique IDs based on the current time and random numbers.

## Installation

To use this package, ensure you have Go installed on your system. Then, you can install the package using `go get`:

```bash
go get github.com/your-username/your-package-name
```

## Generating IDs
To generate a unique ID, call the GenerateId function:

```bash
id := yourpackagename.GenerateId()
```

## Performance Testing
The package includes a performance testing function PerfTest to measure the time taken to generate a specified number of IDs:

```bash
KSortedUID.PerfTest(length)
```

