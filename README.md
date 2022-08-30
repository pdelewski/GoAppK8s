# GoAppK8s
Go app in k8s

1. After creating Dockerfile run

```
docker build -f Dockerfile -t logwriter-app:v.01 .
```

This will create docker image and tag it according to parameter passed as -t switch

2. Add tag to localhost 

```
docker tag logwriter-app:v.01 localhost:5000/logwriter-app:v.01
```

3. Push into local repo

```
docker push localhost:5000/logwriter-app:v.01
```

4. Set port forwarding on registry pod to 5000 (localhost:5000) by using k9s shift-f or kubectl command

5. Create deployment

```
kubectl create -f logwriter_deployment.yaml
```

6. Check docker images by 
```
docker images
```

There should be something like below 
```
REPOSITORY                      TAG                        IMAGE ID       CREATED          SIZE
logwriter-app                   v.01                       f99ce98c7ba6   27 minutes ago   2.13MB
localhost:5000/logwriter-app    v.01                       f99ce98c7ba6   27 minutes ago   2.13MB
```

After application update you have to remove images by

```
docker rmi [image hash]
```

run delete deployment 
```
kubectl delete deploy logwriter-app-deployment
```

and repeat process
