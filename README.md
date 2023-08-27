# Super Garbanzo Project

This project is a log parser project designed to process and create reports of Quake 3 Arena Server log files. It provides a command-line interface for various report generations operations.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
    - [Command Line Options](#command-line-options)
    - [Examples](#examples)

## Requirements

- Go 1.21 or later
- Additional dependencies as specified in `go.mod`

## Installation

1. Clone this repository:

   ```sh
   git clone https://github.com/brendontj/super-garbanzo.git
   cd super-garbanzo
   ```

2. Install dependencies:

   ```sh
   go mod download
   ```
   
3. Build the project:

   ```sh
   go build -o ./bin/super-garbanzo
   ```

4. Run the project:

   ```sh
   ./bin/super-garbanzo 
   ```

## Usage

### Command Line Options

- `-file=`, default value `log-samples/qgames.log`: Specify the input file to parse
- `-streamSize=`, default value `1000`: Specify the size of the buffered channel used to stream log entries
- `-reportType=`, default value `game`: Specify the type of report to generate. Valid values are `game`, `score` and `death-causes`.
- `-gameIdentifier=`, default value `-1`: Specify the game identifier to generate a report for. If the value is `-1`, then a report will be generated for all games otherwise a report will be generated for the specified game identifier.

### Examples

```sh
- ./bin/super-garbanzo -reportType=game -gameIdentifier=-1
- ./bin/super-garbanzo -reportType=game -gameIdentifier=10
- ./bin/super-garbanzo -reportType=score -gameIdentifier=-1
- ./bin/super-garbanzo -reportType=score -gameIdentifier=10
- ./bin/super-garbanzo -reportType=death-causes -gameIdentifier=-1
- ./bin/super-garbanzo -reportType=death-causes -gameIdentifier=10
```


