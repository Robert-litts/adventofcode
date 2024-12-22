#!/bin/bash
#This script generates a new Advent of Code folder and pulls the input for the day
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

#Set the day variable from the command line argument
day=$1
dir="day$day"

# Check if the directory exists & create folder
if [ -d "$dir" ]; then
    echo "Directory $dir already exists."
    exit 1
else
    echo "Creating new dir for day $day"
    mkdir -p "$dir"
fi

#Navigate to the new folder
cd "day$day" || exit

#Create the new Go file
touch "main.go"

#Pull the input for the day
cd .. && ./input/input --day 17

#Create blank example file
touch ./input/day$day.example

echo "Created folder and input for day $day"