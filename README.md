# Text Indexer
**TextIndexer** is a Go-based command-line tool designed to efficiently index and retrieve text content. It processes large text files by generating **SimHash fingerprints** for fixed-size chunks, enabling fast and accurate content lookup. Whether you're indexing documents or searching for specific text, TextIndexer delivers performance and simplicity.

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Algorithms and Techniques](#algorithms-and-techniques)
   - [SimHash](#simhash)
   - [Concurrency](#concurrency)
   - [Index Structure](#index-structure)
4. [Installation and Usage](#installation)
   - [How to Run](#how-to-run)
   - [Indexing](#indexing-a-text-file)
   - [Lookup](#looking-up-content-by-simhash)
5. [Output](#output)
   - [Indexing](#indexing-output)
   - [Lookup](#lookup-output)
6. [Implementation Details](#implementation-details)
   - [File Processing](#file-processing)
   - [Index Serialization](#index-serialization)
7. [Advanced Features](#advanced-features)
   - [Parallel Processing](#parallel-processing)
   - [Fuzzy Search](#fuzzy-search)
8. [Performance Insights](#performance-insights)
    - [Memory Efficiency](#memory-efficiency)
    - [Scalability](#scalability)
9. [Testing](#testing)
    - [How to Run Unit Tests](#unit-tests)
10. [Contributors](#contributors)
11. [License](#license)


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

### Concurrency Model

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



