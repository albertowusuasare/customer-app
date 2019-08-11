#!/bin/sh
set -e
source gomod.sh
ARTIFACT_HOME="artifacts"

function createArtifact(){
    FUNCTIONS_HOME=$1
    echo $FUNCTIONS_HOME
    CURRENT_VERSION=$(git describe --abbrev=0)
    ARTIFACT_NAME="customer_svs_function_$CURRENT_VERSION.zip"
    cd $FUNCTIONS_HOME
    CreateModule functions
    mkdir $ARTIFACT_HOME
    zip -rj $ARTIFACT_HOME/$ARTIFACT_NAME $FUNCTIONS_HOME
}

function uploadToCloudStorage(){
    echo "Uploading function zip file to cloud storage ..."
    CLOUD_STORAGE_PATH="gs://onua-service-artifacts/customer-svs/"
    gcloud auth activate-service-account --key-file $GOOGLE_APPLICATION_CREDENTIALS
    gcloud config set account onua-cicd@onua-246719.iam.gserviceaccount.com
    gsutil cp $ARTIFACT_HOME/$ARTIFACT_NAME $CLOUD_STORAGE_PATH
}

function cleanUp(){
    rm go.mod 
    rm go.sum
    rm -rf $ARTIFACT_HOME
}

function UploadArtifact(){
    echo "Beginning artifact upload ..."
    createArtifact $1
    uploadToCloudStorage
    cleanUp
    echo "Artifact build completed :)"
}
