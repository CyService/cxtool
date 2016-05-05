# ___cxtool___: CX File Utility

[![Build Status](https://travis-ci.org/cytoscape-ci/cxtool.svg?branch=master)](https://travis-ci.org/cytoscape-ci/cxtool)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LISENSE)

## Introduction
This is a file conversion utility for [CX](https://docs.google.com/document/d/1kAUzVj6X86YCWHnTyZtybh1lt4zO-M6anCMJBD_PyG0/edit?usp=sharing).  Main function of this tool is data format conversion from/to SIF, Cytoscape.js JSON and many others.

## Status
* 5/5/2016: Repository moved to cyService.
* 11/20/2015: Pre alpha.  Simply converts CX to basic Cytoscpae.js JSON.

## Build 

In addition to build executables, you need to generate 

* https://github.com/jteeuwen/go-bindata

Run:

```
go-bindata -o converter/ data/
```

to generate new bundle.  And copy it to  

### Supported Functions

* ___From CX___
    * Cytoscape.js (both Style and Graph)
    
* ___To CX___
    * SIF (Simple Interaction Format)
    * Cytoscape.js (ONGOING)
    * MITAB (TODO)
    * GraphML (TODO)

## How to Use
This is a small collection of tools to support round-trip for CX and 
related documents.
 
### Basic Usage
1. Show Help

```
cxtool -h
```

Or simply

```
cxtool
```


1. Convert CX file into Cytoscape.js

```bash
cxtool input_file
```

We recommend to install [jq](https://stedolan.github.io/jq/) for generating human-friendly output.

The following command creates nicely formatted Cytoscape.js JSON. 

```bash
cxtool input_file | jq
```

Or, use pipe for input

```bash
cat network.cx | cxtool | jq .
curl http://example.com/network.cx | cxtool | jq . > cytoscapejs1.json
```

2. Create CX from SIF

```bash
cxtool -f sif galFiltered.sif | jq | from_sif.cx
```


### Command Line Options
(TBD)

### License
MIT License

### Question?
Please send your question to (kono at ucsd edu).

----

&copy; 2015 The Cytoscape Consortium
