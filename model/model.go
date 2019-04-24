package model

// Note - structure for notes data
type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

//NoteDDL - DDL of `notes` table
const NoteDDL = `
CREATE TABLE notes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    body TEXT
);
`
