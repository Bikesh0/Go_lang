#!/bin/bash

# Count regular files and directories including the current directory
count=$(find . -type f -o -type d | wc -l)

# Print only the number
echo "$count"