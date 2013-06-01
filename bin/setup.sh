#!/usr/bin/env bash
DB=db/statistics.db
rm $DB
sqlite3 $DB < config/schema.sql

go get github.com/kuroneko/gosqlite3
