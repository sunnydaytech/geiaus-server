TAG=20160626-01
docker build -t gcr.io/nich01as-com/geiaus-server:$TAG .
gcloud docker push gcr.io/nich01as-com/geiaus-server:$TAG
# update deplouments
cd ..
kubectl apply -f deploy/deployment.yaml

