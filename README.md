# Text Indexer
**TextIndexer** is a Go-based command-line tool designed to efficiently index and retrieve text content. It processes large text files by generating **SimHash fingerprints** for fixed-size chunks, enabling fast and accurate content lookup. Whether you're indexing documents or searching for specific text, TextIndexer delivers performance and simplicity.

## Table of Contents

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Algorithms and Techniques](#algorithms-and-techniques)
   - [SimHash](#simhash)
   - [Concurrency](#concurrency)
   - [Index Structure](#index-structure)
4. [Installation](#installation)
5. [Usage](#usage)
   - [How to Run](#how-to-run)
   - [Indexing](#indexing-a-text-file)
   - [Lookup](#looking-up-content-by-simhash)
6. [Output](#output)
   - [Indexing](#indexing-output)
   - [Lookup](#lookup-output)
7. [Implementation Details](#implementation-details)
   - [File Processing](#file-processing)
   - [Index Serialization](#index-serialization)
8. [Advanced Features](#advanced-features)
   - [Parallel Processing](#parallel-processing)
   - [Fuzzy Search](#fuzzy-search)
9. [Performance Insights](#performance-insights)
    - [Memory Efficiency](#memory-efficiency)
    - [Scalability](#scalability)
10. [Testing](#testing)
    - [How to Run Unit Tests](#unit-tests)
11. [Contributors](#contributors)
12. [License](#license)

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



