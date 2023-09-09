# Repositorio: https://github.com/arceruiz/tech-challenge

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


# Challange 3 Tasks

## A Solução:

Para solucionar o problema, a lanchonete irá investir em um sistema de autoatendimento de fast food, que é composto por uma série de dispositivos e interfaces que permitem aos clientes selecionar e fazer pedidos sem precisar interagir com um atendente, com as seguintes funcionalidades:
              
### Pedido:  

Os clientes são apresentados a uma interface de seleção na qual podem optar por se identificarem via CPF, se cadastrarem com nome, e-mail ou não se identificar, podendo montar o combo na seguinte sequência, sendo todas elas opcionais: 
1. Lanche
1. Acompanhamento
1. Bebida

Em cada etapa é exibido o nome, descrição e preço de cada produto.

### Pagamento:  

O sistema deverá possuir uma opção de pagamento integrada para MVP. A forma de pagamento oferecida será via QRCode do Mercado Pago.
              
### Acompanhamento:  

Uma vez que o pedido é confirmado e pago, ele é enviado para a cozinha para ser preparado. Simultaneamente deve aparecer em um monitor para o cliente acompanhar o progresso do seu pedido com as seguintes etapas:
1. Recebido
1. Em preparação
1. Pronto
1. Finalizado
              
### Entrega:  

Quando o pedido estiver pronto, o sistema deverá notificar o cliente que ele está pronto para retirada. Ao ser retirado, o pedido deve ser atualizado para o status finalizado.
Além das etapas do cliente, o estabelecimento precisa de um acesso administrativo:

### Gerenciar clientes:  

Com a identificação dos clientes o estabelecimento pode trabalhar em campanhas promocionais.

### Gerenciar produtos e categorias:  

Os produtos dispostos para escolha do cliente serão gerenciados pelo estabelecimento, definindo nome, categoria, preço, descrição e imagens. Para esse sistema teremos categorias fixas:

1. Lanche
1. Acompanhamento
1. Bebida
1. Sobremesa

### Acompanhamento do pedido: 
Deve ser possível acompanhar os pedidos em andamento e tempo de espera de cada pedido.
              
As informações dispostas no sistema de pedidos precisarão ser gerenciadas pelo estabelecimento através de um painel administrativo.


## Entregáveis Fase 3:

Dando continuidade ao desenvolvimento do software para a lanchonete, teremos as seguintes melhorias e alterações:

1. Implementar um API Gateway e um AWS Lambda para autenticar o cliente com base no CPF:
    1. Integrar ao Amazon Cognito para identificar o cliente.
    1. Usar Lambda Authorizer para esse processo.
1. Implementar as melhores práticas de CI/CD para a aplicação, segregando os códigos em repositórios, por exemplo:
    1. 1 repositório para o lambda.
    1. 1 repositório para sua infra EKS com Terraform.
    1. 1 repositório para sua infra RDS com Terraform.
    1. 1 repositório para sua aplicação que é executada no Kubernetes.
1. Os repositórios devem fazer deploy automatizado na conta AWS utilizando Github Actions, as branchs main/master devem ser protegidas não permitindo commits direto, sempre utilizar pull request.
1. Melhorar a estrutura do banco de dados escolhido, documentar seguindo os padrões de modelagem de dados e justificar a escolha do banco de dados.