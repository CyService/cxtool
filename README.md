# ___cxtool___: CX to Cytoscape.js File Format Converter

## Introduction
This is a command line tool to convert [CX](https://docs.google.com/document/d/1kAUzVj6X86YCWHnTyZtybh1lt4zO-M6anCMJBD_PyG0/edit?usp=sharing) 
format into Cytoscape.js compatible JSON.

## How to Use
This is a command-line utility to convert
 
```bash
cxtool input_file
```

We recommend to install [jq](https://stedolan.github.io/jq/) for generating human-friendly output.
 
The following command creates nicely formatted Cytoscape.js JSON. 

```bash
cxtool input_file | jq .

(or use pipe for input)

cat network.cx | cxtool | jq .
```

### Options
(TBD)

### License
MIT License

### Question?
Please send your question to (kono at ucsd edu).
