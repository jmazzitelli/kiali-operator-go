To create the skeleton project:

```
export GOPATH=/source/operator-demo/kiali-operator-go
export PATH=$GOPATH/bin:$PATH
mkdir -p $GOPATH/src/github.com/kiali
cd $GOPATH/src/github.com/kiali
operator-sdk new kiali-operator
cd kiali-operator
operator-sdk add api --api-version=op.kiali.io/v1alpha1 --kind=Kiali
operator-sdk add controller --api-version=op.kiali.io/v1alpha1 --kind=Kiali
sed -i 's|REPLACE_IMAGE|kiali/kiali-operator|g' deploy/operator.yaml
```

To build the operator image:

```
operator-sdk build kiali/kiali-operator
```

