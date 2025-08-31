# Ping Pong App

First navigate to the `log-output` directory:

```shell
cd ping-pong
```

To deploy the application, run:

```shell
kubectl apply -k ../kubernetes
```

or use the `deploy.sh` script:

To access the application, run:

```shell
kubectl get ksvc ping-pong-app -o wide
``` 

and then access the URL provided.