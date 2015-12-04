#!/bin/bash

dirs="types logger flags loaders"
for d in $dirs; do
  cd $d > /dev/null 2>&1
  echo "Building $d"
  go build
  go install
  cd ..
done

echo "Building ConfigurationProcessor"
go build
