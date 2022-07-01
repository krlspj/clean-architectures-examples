modelID=test_`date +%Y%m%d%H%M%S`
modelName=test1
versionName=v1

######################
# pull GCS file
###################### 

go run main.go gcs pull_dataset mnist.npz

######################
# store to GCS 
######################

# make test model dir
mkdir ./export/${modelID}
echo ${modelID} >> ./export/${modelID}/saved_model.pb
echo ${modelID} >> ./export/${modelID}/vec.txt
mkdir ./export/${modelID}/variables
echo ${modelID} >> ./export/${modelID}/variables/variables.data-00000-of-00001
echo ${modelID} >> ./export/${modelID}/variables/ariables.index

go run main.go gcs store_model ${modelName} ${versionName} ${modelID}

######################
# model create
###################### 

go run main.go model create ${modelName} 

######################
# version create
###################### 

go run main.go version create ${modelName} ${versionName} ${modelID}

######################
# show config
######################

go run main.go config ${modelName} ${versionName} ${modelID}