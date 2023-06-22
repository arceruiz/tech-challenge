CREATE DATABASE FIAP_TECH_CHALLENGE;

-- Utilização do banco de dados
\c FIAP_TECH_CHALLENGE;

-- Criação das tabelas
CREATE TABLE Customer (
  ID VARCHAR(255) PRIMARY KEY,
  Name VARCHAR(255),
  Email VARCHAR(255),
  Document VARCHAR(255),
  Password VARCHAR(255),
  CreatedAt VARCHAR(255)
);

CREATE TABLE Product (
  ID VARCHAR(255) PRIMARY KEY,
  Name VARCHAR(255),
  Description VARCHAR(255),
  Price DECIMAL,
  Category VARCHAR(255),
  Imagepath VARCHAR(255)
);

CREATE TABLE "Order_Items" (
  ID VARCHAR(255) PRIMARY KEY,
  OrderID VARCHAR(255) REFERENCES "Order"(ID),
  ProductID VARCHAR(255) REFERENCES Product(ID),
  Quantity VARCHAR(255)
);

CREATE TABLE "Order" (
  ID VARCHAR(255) PRIMARY KEY,
  CustomerID VARCHAR(255) REFERENCES Customer(ID),
  Status VARCHAR(255)
);

CREATE TABLE Payment (
  ID VARCHAR(255) PRIMARY KEY,
  PaymentType INT,
  CreatedAt VARCHAR(255)
);

ALTER TABLE "Order_Items"
ADD CONSTRAINT order_items_order_fk FOREIGN KEY (OrderID) REFERENCES "Order" (ID),
ADD CONSTRAINT order_items_product_fk FOREIGN KEY (ProductID) REFERENCES Product (ID);

-- Adicionando constraints para a tabela Order
ALTER TABLE "Order"
ADD CONSTRAINT order_customer_fk FOREIGN KEY (CustomerID) REFERENCES Customer (ID);

-- Adicionando constraints para a tabela Payment
ALTER TABLE Payment
ADD CONSTRAINT payment_pk PRIMARY KEY (ID);