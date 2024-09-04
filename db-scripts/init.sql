CREATE TABLE IF NOT EXISTS greets (
  id INT PRIMARY KEY,
  message VARCHAR(255) NOT NULL
);

INSERT INTO greets (id, message)
VALUES (1,'hello from mysql ;)')
ON DUPLICATE KEY UPDATE message = VALUES(message);
