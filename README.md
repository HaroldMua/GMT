# GMT

GMT is a distributed cluster GPU monitor. It can cooperate with SCHEDULER to achieve fine-grained scheduling task. 

GMT is an operator using [kubebuilder](https://book.kubebuilder.io/) tool. Therefore, generating code, creating CRD YAMLs, and deploying code are all done in a standard way.



## GPU metrics that GMT can monitor

- Total Memory 
- Free Memory 
- GPU Number
- Core Frequency
- Memory Frequency
- Model
- Bandwidth
- Power
- GPU Utilization
- Temperature

### CRD Example

```yaml
apiVersion: core.gpumon.com/v1
kind: Gmt
metadata:
  creationTimestamp: "2021-06-22T01:40:45Z"
  generation: 259
  managedFields:
  - apiVersion: core.gpumon.com/v1
    fieldsType: FieldsV1
    fieldsV1:
      f:spec:
        .: {}
        f:updateInterval: {}
      f:status:
        .: {}
        f:cardList: {}
        f:cardNumber: {}
        f:freeMemorySum: {}
        f:totalMemorySum: {}
        f:updateTime: {}
    manager: GMT
    operation: Update
    time: "2021-06-22T01:55:09Z"
  name: gpu03-poweredge-t420
  resourceVersion: "32504258"
  selfLink: /apis/core.gpumon.com/v1/gmts/gpu03-poweredge-t420
  uid: 0ee0cb26-3712-480a-8370-54160c4c981d
spec:
  updateInterval: 1000
status:
  cardList:
  - bandwidth: 8000
    clock: 5705
    core: 1911
    freeMemory: 11408
    gpuUtil: 5
    health: Healthy
    id: 0
    model: TITAN Xp
    power: 250
    temperature: 64
    totalMemory: 12196
  cardNumber: 1
  freeMemorySum: 11408
  totalMemorySum: 12196
  updateTime: "2021-06-22T01:55:09Z"
```

## Deploy

Deploy GMT into kubernetes cluster:

```sh
kubectl apply -f config/crd/bases/core.gpumon.com_gmts.yaml
kubectl apply -f deploy/deploy.yaml
```

## Test

First, check if the CRD was successfully installed by:

```sh
kubectl get CRD
```

Second, check the controller & CR by:

```sh
kubectl get all -n gmt
kubectl get gmt
```

## Undeploy

Undeploy GMT from kubernetes cluster:

```sh
kubectl delete -f deploy/deploy.yaml
kubectl delete -f config/crd/bases/core.gpumon.com_gmts.yaml
```

 