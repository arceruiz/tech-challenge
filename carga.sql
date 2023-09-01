-- Inserir clientes fictícios
INSERT INTO "Customer" (Name, Email, Document, Password, CreatedAt)
VALUES
  ('John Doe', 'john@example.com', '123456789', 'password123', NOW())
  RETURNING ID;

-- Inserir produtos fictícios
INSERT INTO "Product" (Name, Description, Price, Category, Status, Imagepath)
VALUES
  ('Smartphone', 'High-end smartphone', 923.1, 'Electronics', 1, 'phone.jpg')
  RETURNING ID;

-- Inserir tipos de pagamento fictícios
INSERT INTO "Payment" (PaymentType, CreatedAt)
VALUES
  (1, NOW())
  RETURNING ID;

-- Inserir pedidos fictícios
INSERT INTO "Order" (CustomerID, PaymentID, Status, CreatedAt, UpdatedAt, Total)
VALUES
  (1, 3, 1, NOW(), NOW(), 799.99)
  RETURNING ID;

-- Inserir itens de pedido fictícios
INSERT INTO "Order_Items" (OrderID, ProductID, Quantity)
VALUES
  (2, 2, 1),
  (3, 3, 2);

-- Inserir itens de pedido fictícios
INSERT INTO "Order_Items" (OrderID, ProductID, Quantity)
VALUES(1,1,1);

INSERT INTO "Order_Items" (OrderID, ProductID, Quantity)
VALUES
  ((SELECT ID FROM "Order" WHERE ID = 2), (SELECT ID FROM "Product" WHERE ID = 2), 1);

INSERT INTO "Order_Items" (OrderID, ProductID, Quantity)
VALUES
  ((SELECT ID FROM "Order" WHERE ID = 3), (SELECT ID FROM "Product" WHERE ID = 3), 2);
