#!/bin/bash -eux
cd $(dirname $0)

case "$PLATFORM" in
x86_64*)
  ARCH="x86_64"
  GOARCH="amd64"
  ;;
arm64*)
  ARCH="arm64"
  GOARCH="arm64"
  ;;
esac

if [ ! -e osxcross ]; then
  git clone https://github.com/tpoechtrager/osxcross.git
fi
(
  cd osxcross
  git pull
  OSXCROSS=$(pwd)

  sudo ./tools/get_dependencies.sh
  wget -q https://github.com/phracker/MacOSX-SDKs/releases/download/$SDK_VERSION/MacOSX$SDK_VERSION.sdk.tar.xz -O tarballs/MacOSX$SDK_VERSION.sdk.tar.xz

  UNATTENDED=1 ./build.sh

  export PATH="$OSXCROSS/target/bin:$PATH"
  export MACOSX_DEPLOYMENT_TARGET="$SDK_VERSION"

  sed -i -e "s/ARCH=\"x86_64\"/ARCH=\"$ARCH\"/" $OSXCROSS/target/bin/osxcross-macports
  sed -i -e 's/verifyFileIntegrity "$pkgfile"/#verifyFileIntegrity "$pkgfile"/' $OSXCROSS/target/bin/osxcross-macports

  UNATTENDED=1 osxcross-macports install libusb

  export OSXCROSS_MP_INC="1"
  export CC="o64-clang"
  export CXX="o64-clang++"
  export AR="$PLATFORM-ar"
  export LD="$PLATFORM-ld"
  export CGO_ENABLED="1"
  export GOOS="darwin"
  export GOARCH="$GOARCH"

  echo PATH=$OSXCROSS/target/bin:$PATH >>$GITHUB_ENV
  echo OSXCROSS_MP_INC=$OSXCROSS_MP_INC >>$GITHUB_ENV
  echo CC=$CC >>$GITHUB_ENV
  echo CXX=$CXX >>$GITHUB_ENV
  echo AR=$AR >>$GITHUB_ENV
  echo LD=$LD >>$GITHUB_ENV
  echo CGO_ENABLED=$CGO_ENABLED >>$GITHUB_ENV
  echo GOOS=$GOOS >>$GITHUB_ENV
  echo GOARCH=$GOARCH >>$GITHUB_ENV
)
