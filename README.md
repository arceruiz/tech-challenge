# DELIVERABLES

## 1 - Kubernetes config files

1. Application deployment containing 2+ PODs
1. Application service load balancer with NLB or ALB type
1. AWS access credentials using secrets configuration

## 2 - Refactor application mathing clean architecture

### Create/Modify APIs

1. Order checkout, should receive ordered products and return order identification;
1. Check order payment status. It should inform payment aproval status;
1. Create a webhook to receive payment approval status;
1. Order list should show retrieve description and ordered by creation time and status following this priority: READY > PREPARING > RECEIVED. Status DELIVERED shouldn't appear on the list.
1. Update order status
1. EXTRA CHALLENGE: implement Mercado Pago integration generating QR Code and add webhook to recieve payments.

### POC validation infra limitation

1. 1 DB instance
1. 1 aplication instance