#!/bin/sh
source tag.sh
source artifact.sh
source manifest.sh

## Setup
cd ..
SRC_ROOT=$(pwd)
HOME=$SRC_ROOT/build_$LATEST_COMMIT
LATEST_COMMIT=$(git describe --abbrev=0)
HOME=$SRC_ROOT/build_$LATEST_COMMIT

## Perfom tests
echo ">>> Performing sanity checks ..."
cd $SRC_ROOT
make lint
make unit-test
make int-test

## Increment semantic version
echo ">>> Updating semantic version ..."
IncrementMinor

## Upload artifact
echo ">>> Uploading artifact to cloud storage ..."
UploadArtifact $SRC_ROOT/functions

## Update manifest
cd $HOME
echo ">>> Updating Manifest with new version"
UpdateManifest $(git describe --abbrev=0)

## Clean up build 
echo ">>> Cleaning up build location"
cd $HOME
cd ..
rm -rf $HOME
