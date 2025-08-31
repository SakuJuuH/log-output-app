# log-output-app

To deploy the app, run the deploy script:

```bash
scripts/deploy.sh
```

Alternatively, you can deploy it manually:

```bash
kubectl apply -k kubernetes/
```

> [!NOTE]
> Currently, the ping pong app is serverless and has not been tested with the broader log-output application, so it might not work.
