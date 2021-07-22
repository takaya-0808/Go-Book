DROP TABLE IF EXISTS books;

CREATE TABLE books (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(32) NOT NULL,
    content VARCHAR(32) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO books (id,title,content) VALUES (1, 'TOM','xxxx@mail.co.jp');
INSERT INTO books (id,title,content) VALUES (2, 'Thoge','hogehoge');
INSERT INTO books (id,title,content) VALUES (3, 'fuga','fugafuga');