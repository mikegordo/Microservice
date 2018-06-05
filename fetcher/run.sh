#!/bin/sh

while :
do
    /go/bin/fetcher
    sleep $FETCHER_FREQ
done