#!/bin/bash
FILES=vendor/github.com/fortifi/proto/*.proto
for f in $FILES
do
  file=$(basename "$f")
  package="${file%.*}"
  echo "Processing $package"
  mkdir -p $package
  protoc -I vendor/github.com/fortifi/proto --go_out=plugins=grpc:$package $f
  git add $package/*
done