# Text Indexer
**TextIndexer** is a Go-based command-line tool designed to efficiently index and retrieve text content. It processes large text files by generating **SimHash fingerprints** for fixed-size chunks, enabling fast and accurate content lookup. Whether you're indexing documents or searching for specific text, TextIndexer delivers performance and simplicity.

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Algorithms and Techniques](#algorithms-and-techniques)
   - [SimHash](#simhash-implementation)
   - [Concurrency](#concurrency)
   - [Index Structure](#index-structure)
4. [Installation and Usage](#installation-and-usage)
   - [Dependencies](#dependencies)
   - [How to Run](#how-to-run)
   - [Indexing](#indexing-a-text-file)
   - [Lookup](#looking-up-content-by-simhash)
5. [Output](#output)
   - [Indexing](#indexing-output)
   - [Lookup](#lookup-output)
6. [Advanced Features](#advanced-features)
   - [Parallel Processing](#parallel-processing)
   - [Fuzzy Search](#fuzzy-search)
7. [Testing](#testing)
8. [Contributors](#contributors)
9. [License](#license)


## Introduction
This program was developed to address the challenge of efficiently indexing and retrieving content from large text files. Traditional methods of text search can be slow and resource-intensive, especially when dealing with massive datasets. 

By leveraging **SimHash**, a technique for generating similarity-preserving fingerprints, TextIndexer provides a fast and scalable solution for content lookup and near-duplicate detection.

This tool is particularly useful for applications like:
- **Document similarity detection**: Identifying near-duplicate content in large text corpora.
- **Content retrieval**: Quickly locating specific phrases or chunks within large files.
- **Text analysis**: Enabling efficient processing of large datasets for NLP (Natural Language Processing) tasks.

This document provides a detailed explanation of the tool, including the algorithms used, the approach taken, performance insights, unique techniques, and clear instructions on how to run the program and its test cases.


## Features

- **Efficient Text Chunking**: Splits large text files into fixed-size chunks (configurable via `-s` flag) for granular processing, enabling efficient handling of massive files.

- **SimHash Generation**: Computes 64-bit SimHash fingerprints for each chunk using the `FNV-1a` hash function, preserving similarity for near-duplicate detection.

- **Parallel Processing**: Leverages Go's goroutines and worker pools to process chunks concurrently, maximizing performance on multi-core systems.

- **Persistent Indexes**: Stores SimHash values and byte offsets in a serialized index using Go's `gob` encoding, ensuring fast loading and retrieval.

- **Fast Lookup**: Retrieves text chunks and their positions in near-constant time using SimHash values, enabling instant content access.

- **Human-Readable Output**: Generates a `simhash.txt` file alongside the binary index `index.idx`, listing SimHash values and byte offsets for easy inspection.

- **Data Integrity Verification**: Ensures retrieved content matches the original text through checksum validation and file protection mechanisms, preventing tampering or corruption.

- **Robust Error Handling**: Validates input files (checks existence, file type, and non-empty content) and provides clear error messages for reliable operation.

- **Memory-Efficient Design**: Processes files in chunks without loading the entire file into memory, supporting arbitrarily large text files.

- **Configurable Parameters**: Customize chunk size (`-s`) and worker count (based on CPU cores) for optimal performance tuning.

## Algorithms and Techniques

### SimHash Implementation

TextIndexer uses the **SimHash algorithm**, a locality-sensitive hashing technique originally proposed by Moses Charikar. SimHash is particularly valuable for text processing because:

- **Similarity Preservation**: Similar text chunks produce similar hash values.
- **Fixed Output Size**: Regardless of input size, the output is always a fixed-length hash (64 bits).
- **Efficient Comparison**: Hamming distance between hashes correlates with text similarity.

Our implementation follows these steps:
1. **Tokenization**: Text is split into words/tokens.
2. **Weight Assignment**: Token frequencies are used as weights.
3. **Feature Hashing**: Each token is hashed using the `FNV-1a` hash function (fast and non-cryptographic).
4. **Vector Aggregation**: A 64-dimensional vector accumulates weighted contributions from each token.
5. **Threshold Application**: The final hash is constructed by applying a threshold to each dimension.

The implementation efficiently reuses hash function instances within goroutines to minimize memory allocations and improve performance.

---

### Concurrency

TextIndexer employs a sophisticated concurrency model to maximize throughput:

- **Producer-Consumer Pattern**: The main goroutine reads file chunks and feeds them to worker goroutines.
- **Worker Pool**: Multiple worker goroutines process chunks in parallel.
- **Fan-Out/Fan-In**: Chunk processing fans out to workers, and results fan back in to a collector.
- **Synchronized Collection**: A dedicated collector goroutine aggregates results safely.

This architecture ensures:
- **Load Balancing**: Work is distributed evenly across available CPU cores.
- **Resource Utilization**: Maximum CPU utilization without oversubscription.
- **Minimal Contention**: Channel-based communication reduces lock contention.
- **Controlled Memory Usage**: Buffered channels prevent unbounded memory growth.

---

### Index Structure

The index data structure is designed for efficient storage and retrieval:

- **Hash-to-Offset Mapping**: Uses a map with SimHash values as keys and byte offset slices as values.
- **Metadata Inclusion**: Stores the original filename and chunk size for self-contained indexes.
- **Multiple References**: Handles cases where the same SimHash appears in multiple locations.

This structure offers:
- **O(1) Lookup**: Constant-time access to byte offsets for any SimHash.
- **Compact Representation**: Only essential data is stored.
- **Serialization Support**: Compatible with Go's `gob` encoder for efficient persistence.


## Installation and Usage

These are detailed instructions for installing and using **TextIndexer**
Follow these steps to build the executable binary and run the tool for indexing, lookup, and fuzzy search.

---

### Dependencies

Before building and running **TextIndexer**, ensure your system meets the following requirements:

- **Go**: Version 1.21 or higher (as specified in `go.mod`).
- **Operating System**: Linux, macOS, or Windows.
- **Memory**: At least 2GB of RAM (recommended for large files).
- **Disk Space**: Sufficient space to store the input text file and the generated index.

---
### How to Run

#### Step 1: Open the Terminal and Navigate to the Extracted Directory

1. Open a terminal.

2. Navigate to the extracted `quik` directory. For example:

```bash
   cd /path/to/extracted/quik
```
   
### Step 2: Build the Executable Binary
```bash
 go build -o textindex main.go 
 ```
This creates an executable binary named textindex in the current directory.

 ### indexing-a-text-file
 To index a  your text file, use the following command strictly:

 ```bash
 ./textindex -c index -i <input_file.txt> -s <chunk_size> -o <index_file.idx>
 ```
 where
```
 -c index: Specifies the indexing command.

 -i <input_file.txt>: Path to the input text file (must be a .txt file).

 -s <chunk_size>: Size of each chunk in bytes (default: 4096).

 -o <index_file.idx>: Path to save the generated index file(which is a binary file).
```

**Example Command**:
```bash
./textindex -c index -i sample.txt -s 4096 -o index.idx
```

 ### looking-up-content-by-simhash
To look up content using a SimHash value, use the lookup command strictly:

 ```bash
./textindex -c lookup -i index.idx -h <simhash_value>
 ```
 where
```
-c lookup: Specifies the lookup command.

-i index.idx: Path to the index file(the binary file from indexing).

-h <simhash_value>: SimHash value to look up (in hexadecimal format) which you can retrieve from the file simhash.txt after the indexing.
```

**Example Command**:
```bash
./textindex -c lookup -i index.idx -h 3e4f1b2c98a6
```

 ## Output
### Indexing Output

When you run the indexing command, **TextIndexer** generates two output files:

1. **Binary Index File** (`index.idx`):
   - Contains the serialized index data, including SimHash values and byte offsets.
   - Used for fast lookups.

2. **Human-Readable File** (`simhash.txt`):
   - Lists all SimHash values and their corresponding byte offsets.

**Example Command**:
```bash
go run . -c index -i gb.txt -s 4096 -o index.idx
```

**Example Output**:
```bash
Original file: gb.txt
Chunk size: 4096 bytes
SimHash values and byte offsets written to simhash.txt
```

### Lookup Output

When you perform a lookup, TextIndexer retrieves the following information:

- Original File: The name of the input file.
- Byte Offset: The position of the chunk in the file.
- Phrase: A snippet of text from the retrieved chunk.

**Example Command**:
```bash
go run . -c lookup -i index.idx -h 6f39d09b418d006
```

**Example Output**:
```bash
Original file: gb.txt
Byte offset: 16384
Phrase: This command finds the position of the chunk with the given SimHash
----------
```

## Advanced Features

### Parallel Processing

**TextIndexer** leverages Go’s powerful concurrency model to maximize performance:

- **Worker Pool**: Multiple goroutines process chunks of the input file in parallel, ensuring efficient utilization of multi-core CPUs.
- **Load Balancing**: Work is distributed evenly across available CPU cores, preventing bottlenecks.
- **Efficient Resource Utilization**: Ensures maximum CPU usage without oversubscription, making indexing faster for large files.

**How It Works**:
1. The input file is divided into fixed-size chunks.
2. Each chunk is processed by a separate goroutine to compute its SimHash.
3. Results are collected and aggregated into the final index.

**Benefits**:
- **Faster Indexing**: Parallel processing significantly reduces the time required to index large files.
- **Scalability**: Handles larger files efficiently by utilizing available CPU cores.

---

### Fuzzy Search

**TextIndexer** supports **fuzzy search** for near-matching SimHash values, enabling approximate matching of text chunks:

- **Hamming Distance**: Compares SimHash values to find chunks with similar content.
- **Approximate Matching**: Retrieves chunks that are close to the provided SimHash value, even if they are not an exact match.
- **Configurable Threshold**: Users can specify a Hamming distance threshold to control the level of similarity.

**How It Works**:
1. The user provides a SimHash value and a Hamming distance threshold.
2. **TextIndexer** searches the index for SimHash values within the specified threshold.
3. Returns all matching chunks along with their positions in the original file.

**Example Command**:
```bash
./textindexer -c fuzzy -i index.idx -h 6f39d09b418d006
```
where
```
-c fuzzy: Specifies the fuzzy search command.

-i index.idx: Path to the index file.

-h 6f39d09b418d006: SimHash value to search for.

-d 5: Hamming distance threshold (maximum allowed difference between SimHash values).
```

**Example Output**:
```bash
Original file: gb.txt
SimHash: 6f39d09b418d006
Byte offset: 16384
Phrase: ... This command finds the position of the chunk w...
----------
Original file: gb.txt
SimHash: 6f39c0bb418d006
Byte offset: 49152
Phrase: gerprints for chunk similarity. ● The index shou...
----------
```

## Use Cases:

- Near-Duplicate Detection: Find text chunks that are almost identical.

- Similarity Search: Retrieve chunks with similar content, even if they are not exact matches.

**Why These Features Matter**

- Performance: Parallel processing ensures fast indexing, even for large files.

- Flexibility: Fuzzy search allows for approximate matching, expanding the tool’s use cases.

- Scalability: Designed to handle large datasets efficiently.

## Testing
- To run the test and coverage for this project use the following command from the root directory.

```bash
go test ./... -cover
```

## Contributors
- Anne Okingo - [GitHub Profile](https://github.com/Anne-Okingo)
- Kennedy Ada - [GitHub Profile](https://github.com/adaken4)
- Asman Malika - [GitHub Profile](https://github.com/Malika7188)
- Hannah Apiko - [GitHub Profile](https://github.com/hanapiko)
- Andrew Osindo -[GitHub Profile](https://github.com/andyosyndoh)

## License
- This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
