# HDIUtiler

HDIUtiler is a simple Go application that verifies the integrity of DMG files, mounts them, checks for applications, and verifies code signing.

## Overview

This tool performs the following operations:

1. Verifies the integrity of a DMG file using `hdiutil`.
2. Mounts the DMG file.
3. Locates the application within the mounted DMG.
4. Checks the code signing of the application to ensure it is properly signed and trusted.
5. Unmounts the DMG after processing.

## Prerequisites

- Go (version 1.16 or later)
- `hdiutil` (macOS command-line tool for handling disk images)
- Access to a DMG file
- `make` (for building the application)

## Installation

1. Clone the repository:

2. Change to the project directory:

    ```bash
    cd hdiutiler
    ```

3. Build the application using `make`:

    ```bash
    make build
    ```

   This command will:
   - Create the `bin` directory if it doesn't already exist.
   - Build the application for macOS and Linux platforms.
   - Output the binary to the `bin` directory.

4. Optionally, install the application:

    ```bash
    make install
    ```

   This command will copy the macOS binary to `/usr/local/bin` for global access.

## Usage

Run the application with the path to the DMG file as an argument:

```bash
./bin/hdiutiler_macos <file-path>
