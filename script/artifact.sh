#!/bin/sh
set -e
ARTIFACT_HOME="artifacts"

# Create Artifact
HOME=$(pwd)
FUNCTIONS_HOME="functions"
CURRENT_VERSION=$(git describe --abbrev=0)
ARTIFACT_NAME="customer_svs_function_$CURRENT_VERSION.zip"

echo ">>> Creating function zip file ..."
cd $FUNCTIONS_HOME
$HOME/script/gomod.sh $MOD_PKG_NAME
mkdir $ARTIFACT_HOME
zip -rj $ARTIFACT_HOME/$ARTIFACT_NAME $HOME/$FUNCTIONS_HOME

# Upload Artifact
echo ">>> Uploading function zip file to cloud storage ..."
CLOUD_STORAGE_PATH="gs://onua-service-artifacts/customer-svs/"
gcloud auth activate-service-account --key-file $GOOGLE_APPLICATION_CREDENTIALS
gcloud config set account onua-cicd@onua-246719.iam.gserviceaccount.com
gsutil cp $ARTIFACT_HOME/$ARTIFACT_NAME $CLOUD_STORAGE_PATH
#Clean up
rm go.mod 
rm go.sum
rm -rf $ARTIFACT_HOME
echo ">>> Artifact build completed :)"