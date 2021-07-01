#!/bin/bash
go run main.go
cut -f 1 ../popular-names.txt > col1ref.txt
cut -f 2 ../popular-names.txt > col2ref.txt
diff col1.txt col1ref.txt
if [ $? -ne 0 ]; then
    echo 'KO :('
    rm -f col*.txt
    exit 1
fi
diff col2.txt col2ref.txt
if [ $? -ne 0 ]; then
    echo 'KO :('
    rm -f col*.txt
    exit 1
fi
rm -f col*.txt
echo 'OK :)'
