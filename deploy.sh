curl -X PUT 'http://localhost:8008/services/ilsvrc_googlenet' -d '{
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