#!/bin/bash

show=$1

cwd=`pwd`
SAVEIFS=$IFS
IFS=$(echo -en "\n\b")
files=`ls *.avi`

alias plex="/home/rob/bin/plex"

for f in $files
do
  echo "    Executing plex -f $f"
  new_file=`plex -f $f`
  echo "    Moving $f -> \"$new_file\""
  mv "$f" "$new_file"
  echo ""
done

IFS=$SAVEIFS
echo "Done"
