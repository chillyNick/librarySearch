CREATE TABLE author(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(70) NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO author(name)
VALUES
    ('Terry Pratchett'),
    ('Neil Gaiman'),
    ('David Hayward'),
    ('Lisa Lutz'),
    ('Stephen King'),
    ('Peter Straub'),
    ('Douglas Adams')
;

CREATE TABLE book(
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(70) NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO book(title)
VALUES
    ('Good Omens'),
    ('Heads You Lose'),
    ('The Talisman'),
    ('Billy Summers'),
    ('The Wizards of Odd'),
    ('The Colour of Magic')
;

<<<<<<< HEAD
CREATE TABLE author_book(
=======
CREATE TABLE bookAuthor(
>>>>>>> 81317f3 (docker compose with mysql)
      author_id INT NOT NULL,
      book_id  INT NOT NULL,
      FOREIGN KEY (author_id) REFERENCES author(id),
      FOREIGN KEY (book_id) REFERENCES book(id)
);

<<<<<<< HEAD
INSERT INTO author_book(book_id, author_id)
=======
INSERT INTO bookAuthor(book_id, author_id)
>>>>>>> 81317f3 (docker compose with mysql)
VALUES
    (1, 1),
    (1, 2),
    (2, 3),
    (2, 4),
    (3, 5),
    (3, 6),
    (4, 5),
    (5, 1),
    (5, 7),
    (6, 1)
;