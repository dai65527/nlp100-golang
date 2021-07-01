#!/bin/bash
for n in 1 3 5 7 10; do
    echo "case n = $n"
    go run main.go $n < input.txt > mine.log
    head -n $n < input.txt > ref.log
    diff mine.log ref.log
    if [ $? -ne 0 ]; then
        echo 'KO :('
        rm -f mine.log ref.log
        exit 1
    fi
    rm -f mine.log ref.log
    echo 'OK :)'
done
echo 'ALL CASE OK :)'
