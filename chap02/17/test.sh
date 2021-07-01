#!/bin/bash
go run main.go < ../popular-names.txt | sort > mine.log
cut -f 1 ../popular-names.txt | sort | uniq > ref.log
diff ref.log mine.log
if [ $? -ne 0 ]; then
    echo 'KO :('
    rm -f ref.log mine.log
    exit 1
fi
rm -f ref.log mine.log
echo 'OK :)'
