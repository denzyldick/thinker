curl -X PUT 'http://134.209.84.9:8080/services/ilsvrc_googlenet' -d '{
 "description": "image classification service",
 "model": {
  "repository": "/opt/models/ilsvrc_googlenet",
  "init": "https://deepdetect.com/models/init/desktop/images/classification/ilsvrc_googlenet.tar.gz",
  "create_repository": true
 },
 "mllib": "caffe",
 "type": "supervised",
 "parameters": {
  "input": {
   "connector": "image"
  }
 }
}'

curl -X PUT http://134.209.84.9:8080/services/faces -d '{
 "description": "face detection service",
 "model": {
  "repository": "/opt/models/faces",
  "create_repository": true,
  "init":"https://deepdetect.com/models/init/desktop/images/detection/faces_512.tar.gz"
 },
 "mllib": "caffe",
 "type": "supervised",
 "parameters": {
  "input": {
   "connector": "image"
  }
 }
}'


curl -X PUT http://134.209.84.9:8080/faces_emo -d '{
 "description": "face emotion detection service",
 "model": {
  "repository": "/opt/models/faces_emo",
  "create_repository": true,
  "init":"https://deepdetect.com/models/init/desktop/images/detection/faces_emo.tar.gz"
 },
 "mllib": "caffe",
 "type": "supervised",
 "parameters": {
  "input": {
   "connector": "image"
  }
 }'