# README

## Overview

This Go package provides functionality for generating sorted Base64-encoded IDs and performance testing the ID generation process. The package includes methods for encoding and decoding strings using a K-sorted Base64 encoding scheme.


## Generating IDs
To generate a unique ID, call the GenerateId function:

```bash
id := KSortedUID.GenerateId()
```

## Ids Examples
Ids looks like this after generation:
```bash
.NvE9rA62_5B2C1eqjM
.NvEA5nD5AfCQVthfMp
.NvEB1VA59_Rd.nRTAB
```

## Performance Testing
The package includes a performance testing function PerfTest to measure the time taken to generate a specified number of IDs:

```bash
KSortedUID.PerfTest(length)
```

## Credits
https://firebase.blog/posts/2015/02/the-2120-ways-to-ensure-unique_68

https://github.com/twitter-archive/snowflake?tab=readme-ov-file