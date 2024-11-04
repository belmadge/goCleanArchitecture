CREATE TABLE IF NOT EXISTS orders (
                                      id INT AUTO_INCREMENT PRIMARY KEY,
                                      customer_name VARCHAR(255) NOT NULL,
                                      order_date DATE NOT NULL,
                                      status VARCHAR(50) NOT NULL
);
