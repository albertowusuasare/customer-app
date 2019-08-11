#!/bin/sh
source tag.sh
source artifact.sh
source manifest.sh
## Perform sanity checks 
echo ">>> Performing sanity checks ..."
#make lint
#make unit-test
#make int-test

## Update semantic version
echo ">>> Updating semantic version ..."
HOME=$(pwd)
IncrementMinor
UploadArtifact $HOME/../functions
cd $HOME
UpdateManifest $HOME


## Upload artifact to cloud storage
git tag --delete v0.8.0
echo ">>> Uploading artifact to cloud storage ..."
