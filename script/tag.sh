#!/bin/sh

function IncrementMinor(){
    set -e
    echo "Downloading semver ..."
    npm install -g semver
    LATEST_VERSION=$(git describe --abbrev=0)
    INCREMENTED_VERSION=$(semver $LATEST_VERSION  -i minor)
    UPDATED_TAG=v$INCREMENTED_VERSION
    git tag -a $UPDATED_TAG -m "Bump version"
    git push origin --follow-tags 
    echo "Created tag $UPDATED_TAG"
}
