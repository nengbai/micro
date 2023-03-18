1. build docker images
docker build -t icn.ocir.io/oraclepartnersas/baineng-oke-registry:demo-app.v6 .

2. Push images to oracle registry 
docker push icn.ocir.io/oraclepartnersas/baineng-oke-registry:demo-app.v6

3. Deploy to k8s
kubectl apply -f micro-redis.yml

4. Check run status
kubectl -n demo get pod,svc,ing

5. Web access and verify



git commit -a -m "test"
## 打tag
git tag -a "v1.0.0" -m "test ci"
git push origin v1.0.0