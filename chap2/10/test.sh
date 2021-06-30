#!/bin/bash
go run main.go > mine.log
wc -l ../popular-names.txt | awk '{print $1}' > ref.log
diff mine.log ref.log
if [ $? -ne 0 ]; then
    echo 'KO :('
    rm -f mine.log ref.log
    exit 1
fi
rm -f mine.log ref.log
echo 'OK :)'
