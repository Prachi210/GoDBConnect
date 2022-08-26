DROP TABLE IF EXISTS productInfo;

CREATE TABLE productInfo (
  id         SERIAL,
  model      VARCHAR(128) NOT NULL,
  itemType     VARCHAR(255) NOT NULL,
  category     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (id)
);


CREATE TABLE productInfoDuplicate (
  id         SERIAL,
  model      VARCHAR(128) NOT NULL,
  itemType     VARCHAR(255) NOT NULL,
  category     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO productInfo
  (model, itemType, category, price)
VALUES
  ('A1234', 'chair', 'Living room', 156.99),
  ('A4565', 'table', 'Living room', 117.99),
  ('A6788', 'bowl', 'kitchen', 117.99),
  ('A8789', 'bowl', 'kitchen', 401.99);
  