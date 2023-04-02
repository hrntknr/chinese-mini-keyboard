#!/bin/bash -eux
cd $(dirname $0)

if [ ! -e sigtool ]; then
  git clone https://github.com/thefloweringash/sigtool.git
fi
(
  cd sigtool
  git pull
  SIGTOOL=$(pwd)
  mkdir build
  (
    cd build
    cmake ..
    make -j
  )

  export PATH="$SIGTOOL/build:$PATH"
  echo PATH=$SIGTOOL/build:$PATH >>$GITHUB_ENV
)
