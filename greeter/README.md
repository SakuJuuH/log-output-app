# Greeter service

Traffic-splitting-based service that serves 75% of the traffic to greeter service v1 and 25% to the greeter v2
service.

To deploy first navigate to the log-output-app parent folder and run the deploy script:

```bash
scripts/deploy.sh
```

Alternatively you can deploy the service manually:

```
kubectl apply -k kubernetes/
```