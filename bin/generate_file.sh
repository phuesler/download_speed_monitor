#!/usr/bin/env bash

mkdir -p tmp
dd if=/dev/zero bs=1024 count=1010 of=tmp/file
