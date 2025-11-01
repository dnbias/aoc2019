#!/bin/bash

cookie_file="$1"
cookie=$(cat "$cookie_file" 2> /dev/null)
echo cookie: $cookie
part_num=""
part=""

if [ "$cookie" == "" ]; then
  read -p "Enter your session cookie: " cookie
fi
read -p "Enter the day number: " day
read -p "Enter the day part (leave blank for first part): " part_num

mkdir "day$day" 2> /dev/null

if [ "$part_num" != "" ] ; then
    part="#part$part_num"
fi

url="https://adventofcode.com/2019/day/$day$part/input"
echo "contacting $url..."

curl -s "$url"  --cookie "session=$cookie" >> day$day/input$part_num

echo "written response to day$day/input$part_num"

bat "day$day/input$part_num"
