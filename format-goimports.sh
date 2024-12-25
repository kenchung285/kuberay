#!/bin/bash

for file in $(find . -type f -name "*.go"); do
  sed -i '' -e '/import \(/,/\)/ {
      /^\)$/!{
          N
      }
      s/\n[[:space:]]*\n/\n/g
  }' "$file"

  # goimports -local github.com/ray-project/kuberay -w "$file"
done
