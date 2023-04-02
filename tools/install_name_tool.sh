#!/bin/bash -eux
cd $(dirname $0)

export INSTALL_NAME_TOOL=$PLATFORM-install_name_tool
export CODESIGN_ALLOCATE=$PLATFORM-codesign_allocate

$INSTALL_NAME_TOOL -add_rpath /opt/homebrew/opt/libusb/lib/ chinese-mini-keyboard
$INSTALL_NAME_TOOL -change /opt/local/lib/libusb-1.0.0.dylib @rpath/libusb-1.0.0.dylib chinese-mini-keyboard
codesign --force -s - chinese-mini-keyboard
