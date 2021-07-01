#!/bin/bash
n=7
go run main.go $n < ../popular-names.txt
l=`wc -l ../popular-names.txt | awk '{ print $1 }'`
if [ $(($l % $n)) -eq 0 ]; then
    sp=$(($l / $n))
else
    sp=$(($l / $n + 1))
fi
split -l $sp ../popular-names.txt ref

for suf in aa ab ac ad ae af ag; do
    echo conparing mine$suf and ref$suf
    diff mine$suf ref$suf > diff.log
    if [ $? -ne 0 ]; then
        echo 'KO :('
        rm -f mine* ref*
        exit 1
    fi
done
rm -f mine* ref*
echo 'OK :)'
