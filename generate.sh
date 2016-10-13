#!/bin/bash
FILES=vendor/github.com/fortifi/proto/*.proto
for f in $FILES
do
  file=$(basename "$f")
  package="${file%.*}"
  echo "Processing $package $f"
  mkdir -p $package
  protoc -I vendor/github.com/fortifi/proto --go_out=plugins=grpc:../../../ $f
  git add $package/*
done


FILES=vendor/github.com/fortifi/proto/ftypes/*.proto
for f in $FILES
do
  file=$(basename "$f")
  package="${file%.*}"
  echo "Processing $package $f"
  mkdir -p ftypes/$package
  protoc -I vendor/github.com/fortifi/proto --go_out=../../../ $f
done

git add ftypes/*