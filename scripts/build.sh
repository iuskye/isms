#!/bin/bash
# Copyright 2024 Simon Liu <iuskye@foxmail.com>. All rights reserved.
# Use of this source code is governed by a Apache License Version 2.0 style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/iuskye/isms.


#------------------------------- Environment Variable ------------------------------

tool_name=isms
arch_names=(amd64 arm64)
os_names=(linux windows darwin)


#------------------------------- Function Definetions ------------------------------


#---------------------------------- Shell Main Body --------------------------------
cd ../
if [[ -d _output/ ]]; then
    rm -rf _output/
fi
mkdir -p _output/${tool_name}/${os_name}/${arch_name}
for arch_name in "${arch_names[@]}"; do
    for os_name in "${os_names[@]}"; do
        CGO_ENABLED=0 GOOS=${os_name} GOARCH=${arch_name} go build -o _output/${tool_name}/${os_name}/${arch_name}/${tool_name} main.go
        if [[ $? -eq 0 ]]; then
            echo "${arch_name} ${os_name} ${tool_name} build success"
        else
            echo "${arch_name} ${os_name} ${tool_name} build failed"
        fi
    done
done