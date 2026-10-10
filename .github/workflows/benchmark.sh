#!/usr/bin/env bash

set -euo pipefail

# Runs the benchmarks of two checkouts and writes the results to
# <out-dir>/old and <out-dir>/new for benchstat.
#
# The checkouts are benchmarked alternately, one pass at a time, so that slow
# periods on a shared CI machine affect both sides equally rather than skewing
# whichever side happened to run during them.

if [[ $# -ne 3 ]]; then
  echo "usage: $0 <old-dir> <new-dir> <out-dir>"
  exit 255
fi

old_dir="$1"
new_dir="$2"
out_dir="$(cd "$3" && pwd)"

# Noise comes from periods when the machine is busy, which longer runs don't
# average out. More samples do, because benchstat compares medians. Every
# benchmark performs at least thousands of iterations within 100ms, so time is
# better spent on more samples than on longer ones.
count="${BENCH_COUNT:-20}"
benchtime="${BENCH_TIME:-100ms}"
# The _Proto benchmarks measure google.golang.org/protobuf, which PRs don't
# change. They are only useful for manually comparing against canoto.
skip="${BENCH_SKIP:-_Proto}"

run_benchmarks() {
  (
    cd "$1"
    go test -run='^$' -bench=. -skip="$skip" -benchmem -benchtime="$benchtime" ./...
  ) >> "$2"
}

: > "$out_dir/old"
: > "$out_dir/new"
for i in $(seq "$count"); do
  echo "pass $i/$count"
  run_benchmarks "$old_dir" "$out_dir/old"
  run_benchmarks "$new_dir" "$out_dir/new"
done
