#!/bin/bash -eux
cd $(dirname $0)

git clone https://github.com/thefloweringash/sigtool.git
(
  cd sigtool
  mkdir build
  (
    cd build
    cmake ..
    make -j
    export PATH="$(pwd):$PATH"
  )
  if [ -e env.sh ]; then
    rm env.sh
  fi
  echo "echo 'PATH=$PATH' >>\$GITHUB_ENV" >>env.sh
  chmod +x env.sh
)
