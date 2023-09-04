# FIAP - Tech Challenge II

> **Pré-requisitos:**
> 1. make
> 1. docker

# Como buildar:

```shell
make build
```

# Como executar:

```shell
make install
```



### Mauricio:
1. Ajustar implementacao do retorno de order para retornar pagamento
1. Order total field should be calculated in service OK
1. Create secret to resolve db connection and jwt secret
3. Expose endpoint to create orders OK 
1. Order checkout, should receive ordered products and return order identification; OK
2. Event storm's diagram
1. EXTRA CHALLENGE: implement Mercado Pago integration generating QR Code and add webhook to recieve payments.

### Lucas:
1. Update order status
1. Order list should show retrieve description and ordered by creation time and status following this priority: READY > PREPARING > RECEIVED. Status DELIVERED shouldn't appear on the list.
4. Should improove checkout
6. Add build instructions to readme.md
1. Create a webhook to receive payment approval status;
1. Check order payment status. It should inform payment aproval status;
1. Application deployment containing 2+ PODs
1. Application service load balancer with NLB or ALB type




criei /build
criei /deployments
cargas .sql nao deveria estar no /assets? https://github.com/golang-standards/project-layout/tree/master/assets
imagem ainda ta errada no docker compose, precisa usar container repo/registry



implementar repo/updateorder
devolver order no checkout
linkar callback na rest pela negocio 