# Repositorio: https://github.com/arceruiz/tech-challenge

# FIAP - Tech Challenge II

## Integrantes
- Lucas Arce Ruiz - RM349580
- Mauricio Gonçalves Pires Jr - RM349581

## Como buildar local:

> **Pré-requisitos:**
> 1. make
> 1. docker
> 1. kubernetes

```shell
make build-all
```
### Como executar local:
```shell
make run
```
### Como enviar para o k8s:
```shell
kubectl create namespace fiap && kubectl apply -f deployments/eks-manifests -n fiap
```

# Challange 3 Entregáveis:

Dando continuidade ao desenvolvimento do software para a lanchonete, teremos as seguintes melhorias e alterações:

1. Implementar um API Gateway e um function serverless para autenticar o cliente com base no CPF
1. Integrar o sistema de autenticação para identificar o cliente.
1. Implementar as melhores práticas de CI/CD para a aplicação, segregando os códigos em repositórios, por exemplo:
    1. 1 repositório para o lambda.
    1. 1 repositório para sua infra kubernetes com Terraform.  
    1. 1 repositório para sua infra banco com Terraform.
    1. 1 repositório para sua aplicação que é executada no Kubernetes.
1. Os repositórios devem fazer deploy automatizado na conta AWS utilizando Github Actions, as branchs main/master devem ser protegidas não permitindo commits direto, sempre utilizar pull request.
1. Melhorar a estrutura do banco de dados escolhido, documentar seguindo os padrões de modelagem de dados e justificar a escolha do banco de dados.
1. Utilizar serviços serverles (functions)


# Links

- [Link](https://github.com/mauriciodm1998/tech-challenge-auth) para repositorio do Lambda
- [Link](https://github.com/mauriciodm1998/tech-challenge-eks-gitops) para repositorio EKS com Terraform
- [Link](https://github.com/arceruiz/tech-challange-db-gitops) para repositorio RDS com Terraform
- [Link](https://github.com/arceruiz/tech-challenge) para repositório da aplicação principal


# Porque escolhemos PostgreSQL?

Escolhemos utilizar um banco SQL em vez de NoSQL por conta de várias vantagens significativas que os bancos de dados SQL oferecem. Em primeiro lugar, bancos de dados SQL se destacam em termos de integridade e consistência de dados. Eles fornecem um esquema bem definido e que impõe restrições aos dados, garantindo a estrutura e precisão. Isso é particularmente importante para projetos como o nosso, do fast-food, que qualidade e confiabilidade dos dados são fundamentais. Além disso, DBs SQL são adequados para consultas mais complexas e modelagem de dados, facilitando muito em cenários nos quais é preciso analisar dados de forma abrangente e/ou realizar manipulações avançadas de dados para relatórios gerenciais por exemplo. Escolhemos PostgreSQL e não outro DB SQL por diversos fatores, dentre eles:
 - Grande comunidade, facilitando acesso a exemplos e problemas já resolvidos;
 - Free/Opensource, diminuindo a barreira de entrada, diminuindo custos e promovendo uma evolução simplificada da plataforma;
 - Extensibilidade, permite criação de tipos de dado customizados, bem como operadores e funções;
 - Herança de tabelas, permitindo que uma tabela herde propriedades de outra tabela;
 - Full-Text Search e Text Indexing que apesar de existir em outras linguagens, aqui alinhado com os tipos de dados customizáveis permite também a criação de dicionarios de busca textual customizados;
 
