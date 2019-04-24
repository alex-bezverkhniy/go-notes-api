GRANT ALL ON gonotes.* TO 'gonotes'@'%';

USE gonotes;
CREATE TABLE notes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    body TEXT
);
INSERT INTO notes (title, body) VALUES
    ('Very first note', 'Hello this is very first note'),
    ('Second note', 'Simple second note')
;
COMMIT;