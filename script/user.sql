CREATE
DATABASE IF NOT EXISTS bookstore;

ALTER
DATABASE bookstore
  DEFAULT CHARACTER SET utf8
  DEFAULT COLLATE utf8_general_ci;

GRANT ALL PRIVILEGES ON bookstore.* TO
'bookstore@%' IDENTIFIED BY 'bookstore';
