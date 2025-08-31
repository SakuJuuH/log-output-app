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
> Currently, the ping pong kustomization file has been set to Serverless mode