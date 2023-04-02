#!/bin/bash -eux

export INSTALL_NAME_TOOL=$PLATFORM-install_name_tool
export CODESIGN_ALLOCATE=$PLATFORM-codesign_allocate

$INSTALL_NAME_TOOL -add_rpath /opt/homebrew/opt/libusb/lib/ $1
$INSTALL_NAME_TOOL -change /opt/local/lib/libusb-1.0.0.dylib @rpath/libusb-1.0.0.dylib $1
codesign --force -s - $1
