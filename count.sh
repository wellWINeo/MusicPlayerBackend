#!/usr/bin/env sh

count=0

src=$(find . -name "*.go")

for f in $src; do
    num=$(cat $f | wc -l)
    echo "$f = $num"
    echo "---"
    count=$(( $count + $num ))
done

echo "Code lines: $count"
