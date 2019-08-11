#!/bin/sh

function UpdateManifest(){
    set -e
    MANIFEST_HOME=manifest
    SOURCE_VERSION=$1
    mkdir $MANIFEST_HOME
    cd $MANIFEST_HOME
    git clone git@github.com:albertowusuasare/onua-infra.git
    cd onua-infra/customer-app/cloud-func/
    git checkout staging
    sed -i "4c\SOURCE_VERSION=$SOURCE_VERSION" deploy.sh
    git add deploy.sh
    git commit -m "Source version update to $SOURCE_VERSION"
    git push origin head
    cd ..
    rm -rf $MANIFEST_HOME
}

