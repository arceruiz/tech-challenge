# FIAP - Tech Challenge II

> **Pré-requisitos:**
> 1. make
> 1. docker
> 1. kubernetes

# Como buildar local:

```shell
make build-all
```

# Como executar local:

```shell
make run
```

# Como enviar para o k8s:

```shell
kubectl create namespace fiap && kubectl apply -f deployments/manifest.yml -n fiap
```

# DB Efemero

Para executar o DB optamos por seguir com DB efemero. Sendo assim, cada vez que a pod do DB for terminada, os dados serão perdidos. O motivo do DB estar efemero é porque o tipo de volume persistente que a instância de Kubernetes que utilizamos suporta necessita de um caminho local do node em que o volume é montado. 
Por conta disso, seguimos com DB efemero para que os manifestos sejam aplicados e tudo funcione sem nenhum problema/configuração adicional. Para que a persistencia dos dados não seja perdida, basta alterar o arquivo deployments/manifest.yml descomentando as linhas comentadas e ajustar o caminho no PersistentVolume.