#!/bin/bash

# 获取 Bash 脚本所在目录绝对路径
SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
readonly SCRIPT_DIR

# 获取 Bash 脚本文件完整绝对路径
#SCRIPT_FILE=$SCRIPT_DIR/$(basename "${BASH_SOURCE[0]}")
#readonly SCRIPT_FILE

readonly BASE_DIR=$SCRIPT_DIR

gomobile bind -v -target ios -o "$BASE_DIR/../build/common_utils.xcframework" "$BASE_DIR/../pkg"/*

#gomobile bind -v -target android -androidapi 16 -o "$BASE_DIR/../build/common_utils.aar" "$BASE_DIR/../pkg"/*



