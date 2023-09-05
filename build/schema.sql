CREATE TABLE "Customer" (
  ID SERIAL PRIMARY KEY,
  Name VARCHAR(255),
  Email VARCHAR(255),
  Document VARCHAR(255),
  Password VARCHAR(255),
  CreatedAt TIMESTAMP WITH TIME ZONE
);

CREATE TABLE "Product" (
  ID SERIAL PRIMARY KEY,
  Name VARCHAR(255),
  Description VARCHAR(500),
  Price DECIMAL(16, 6),
  Category VARCHAR(255),
  Status INT,
  Imagepath VARCHAR(255)
);

CREATE TABLE "Payment" (
  ID SERIAL PRIMARY KEY,
  PaymentType INT,
  CreatedAt TIMESTAMP WITH TIME ZONE,
  Status INT
);

CREATE TABLE "Order" (
  ID SERIAL PRIMARY KEY,
  CustomerID INT REFERENCES "Customer"(ID),
  PaymentID INT REFERENCES "Payment"(ID),
  Status INT,
  CreatedAt TIMESTAMP WITH TIME ZONE,
  UpdatedAt TIMESTAMP WITH TIME ZONE,
  Total DECIMAL(16, 6)
);

CREATE TABLE "Order_Items" (
  OrderID INT REFERENCES "Order"(ID),
  ProductID INT REFERENCES "Product"(ID),
  Quantity INT 
);

ALTER TABLE "Order_Items"
ADD PRIMARY KEY (OrderID, ProductID),
ADD CONSTRAINT order_items_order_fk FOREIGN KEY (OrderID) REFERENCES "Order" (ID),
ADD CONSTRAINT order_items_product_fk FOREIGN KEY (ProductID) REFERENCES "Product" (ID);

ALTER TABLE "Order"
ADD CONSTRAINT order_customer_fk FOREIGN KEY (CustomerID) REFERENCES "Customer" (ID);
