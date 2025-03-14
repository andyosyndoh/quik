# Text Indexer
**TextIndexer** is a Go-based command-line tool designed to efficiently index and retrieve text content. It processes large text files by generating **SimHash fingerprints** for fixed-size chunks, enabling fast and accurate content lookup. Whether you're indexing documents or searching for specific text, TextIndexer delivers performance and simplicity.

## Table of Contents

## Table of Contents

1. [Introduction](#overview)
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
   -  [Fuzzy Search](#fuzzy-search)
9. [Performance Insights](#performance-insights)
   - [Memory Efficiency](#memory-efficiency)
   - [Scalability](#scalability)
10. [Testing](#testing)
    - [ How to run Unit Tests](#unit-tests)
11. [Contributors](#contributors)
12. [License](#license)
