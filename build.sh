#!/bin/bash
#
# Author by Bai Neng
# Write on Dec 16,2021
####### Build Docker Images ###################################################
docker build -t icn.ocir.io/oraclepartnersas/baineng-oke-registry/demo-redis:v2 .

##### Push daocker images to registry #####################################
docker push icn.ocir.io/oraclepartnersas/baineng-oke-registry/demo-redis:v2

#### Deploy demo to oks ###################################################
kubectl delete -f ./doc/micro-redis.yml 

sleep 5

kubectl apply -f ./doc/micro-redis.yml 



