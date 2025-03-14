# Text Indexer
TextIndexer is a Go-based command-line tool that efficiently processes text files, generates SimHash values for chunks of text, and provides lookup capabilities for quick content identification and retrieval.

## Overview

**TextIndexer** works by breaking down text files into configurable chunks, generating a similarity hash (SimHash) for each chunk, and creating an index that maps these hash values to their corresponding positions in the original file. This allows for efficient content lookup and similarity detection within large text files(input file).

This document provides a detailed explanation of the tool, including the algorithms used, the approach taken, performance insights, unique techniques, a few limitations and clear instructions on how to run the program and its test cases.

---