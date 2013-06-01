#!/usr/bin/env bash

OUTPUT=tmp/file

mkdir -p tmp
dd if=/dev/zero bs=1024 count=1010 of=$OUTPUT
md5 -q $OUTPUT > tmp/checksum.md5
ls -lah tmp/

