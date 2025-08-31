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

App can be accessed in the browser at:

```shell
http://localhost:8081/pingpong
``` 

To run the application serverless:

