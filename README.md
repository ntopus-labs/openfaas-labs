# openfaas-labs
#### Obs: You have to install kubernetes. Tip: Use k3d :)
#### Obs2: You have to install arkade. Follow this instructions: https://docs.openfaas.com/deployment/kubernetes/
```
k3d cluster create testopen

arkade install openfaas --function-pull-policy IfNotPresent

arkade info openfaas

faas-cli build -f csv-to-xlsx.yml

k3d image import csv-to-xlsx:latest --cluster testopen

kubectl port-forward -n openfaas svc/gateway 8080:8080

faas-cli login --password $PASSWORD

faas-cli deploy -f csv-to-xlsx.yml

kubectl --context k3d-testopen get pods -n openfaas-fn

cat "./example.csv" | faas-cli invoke csv-to-xlsx > converted-example.xlsx
```

More information: https://www.openfaas.com/
