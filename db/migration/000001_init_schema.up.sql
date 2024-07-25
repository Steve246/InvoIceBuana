
-- created table customer
CREATE TABLE Customer (
  id INT AUTO_INCREMENT PRIMARY KEY,
  customer_id VARCHAR(255) NOT NULL UNIQUE, 
  customer_name VARCHAR(255) NOT NULL,
  customer_address VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- created table items


CREATE TABLE Item (
  id INT AUTO_INCREMENT PRIMARY KEY,
  item_id VARCHAR(255) NOT NULL UNIQUE, 
  item_name VARCHAR(255) NOT NULL,
  item_type VARCHAR(255) NOT NULL,
  item_price DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


-- created table invoice


CREATE TABLE Invoice (
  id INT AUTO_INCREMENT PRIMARY KEY,
  invoice_id INT NOT NULL UNIQUE,
  subject VARCHAR(255) NOT NULL,
  customer_id VARCHAR(255) NOT NULL, 
  issue_date DATETIME NOT NULL,
  due_date DATETIME NOT NULL,
  sub_total DECIMAL(10, 2) NOT NULL,
  tax DECIMAL(10, 2) NOT NULL,
  grand_total DECIMAL(10, 2) NOT NULL,
  status VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (customer_id) REFERENCES Customer(customer_id) 
);

-- created table invoice_item


CREATE TABLE InvoiceItem (
  id INT AUTO_INCREMENT PRIMARY KEY,
  invoice_id INT NOT NULL, 
  item_id VARCHAR(255) NOT NULL, 
  quantity INT NOT NULL,
  total_price DECIMAL(10, 2) NOT NULL,
  FOREIGN KEY (invoice_id) REFERENCES Invoice(invoice_id), 
  FOREIGN KEY (item_id) REFERENCES Item(item_id) 
);
