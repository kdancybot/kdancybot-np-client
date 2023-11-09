#!/usr/bin/env bash

target=kdancybot-np-client
	
platforms=("windows/amd64" "windows/386" "linux/amd64" "linux/386")

mkdir -p build

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$target'-'$GOOS'-'$GOARCH
	flags=""
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
		flags+="-ldflags -H=windowsgui"
	fi	
	
	env GOOS=$GOOS GOARCH=$GOARCH go build $flags -o build/$output_name
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done