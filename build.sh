#!/usr/bin/env bash

build_dir="build"
platforms=("windows/amd64" "windows/386" "darwin/amd64" "linux/amd64" "linux/386")

rm -rf build/

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })

    GOOS=${platform_split[0]}
    echo "GOOS: $GOOS"

    GOARCH=${platform_split[1]}
    echo "GOARCH: $GOARCH"

    file_name="ww"

    if [ $GOOS = "windows" ]; then
        file_name+='.exe'
    fi

    out_dir=$build_dir'/'$GOOS'-'$GOARCH
    echo "out_dir: $out_dir"

    out_file_path=$out_dir'/'$file_name
    echo "out_file_path: $out_file_path"

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $out_file_path
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi

    pushd $out_dir
    tar_file_name='ww-'$GOOS'-'$GOARCH'.tar.gz'
    tar -czvf $tar_file_name $file_name
    mv $tar_file_name ../$tar_file_name
    popd

    echo ""
    echo ""
    echo ""
done