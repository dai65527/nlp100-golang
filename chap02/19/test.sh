#!/bin/bash
go run main.go < ../popular-names.txt > mine.log
cut -f 1 ../popular-names.txt | sort | uniq -c | sort -sk 1nr | cut -b 6- > ref.log
diff ref.log mine.log
if [ $? -ne 0 ]; then
    echo 'KO :('
    rm -f ref.log mine.log
    exit 1
fi
rm -f ref.log mine.log
echo 'OK :)'
