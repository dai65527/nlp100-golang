#!/bin/bash
go run main.go
paste col1.txt col2.txt > mergedref.txt
diff merged.txt mergedref.txt
if [ $? -ne 0 ]; then
    echo 'KO :('
    rm -f mergedref.txt
    exit 1
fi
rm -f mergedref.txt
echo 'OK :)'
