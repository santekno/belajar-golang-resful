package mysql

const queryGetAll = `SELECT id, title, content, create_at, update_at FROM articles`

const queryGetById = `SELECT id, title, content, create_at, update_at FROM articles WHERE id=?`

const queryUpdate = `UPDATE articles SET title=?, content=?, update_at=? WHERE id=?`

const queryDelete = `DELETE articles WHERE id=?`

const queryInsert = `INSERT INTO articles(title, content, create_at, update_at) VALUES(?,?,?,?)`
